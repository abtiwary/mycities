package mycitiesserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	mycitiesdb "github.com/abtiwary/mycities/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Router *chi.Mux
	IP     string
	Port   int
	Srv    *http.Server
	CityDB *mycitiesdb.CitiesDB
	Exit   chan struct{}
}

func NewServer(ip string, port int, db *mycitiesdb.CitiesDB) (*Server, error) {
	nMlotdServer := new(Server)
	nMlotdServer.IP = ip
	nMlotdServer.Port = port

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	nMlotdServer.Router = r

	nMlotdServer.CityDB = db
	nMlotdServer.Exit = make(chan struct{})

	return nMlotdServer, nil
}

func (s *Server) initAPI() {
	s.Router.Get("/api/v1/geometry/{city}", s.HandleGetRecommendations)

}

func (s *Server) StartHTTPServer() *http.Server {
	s.initAPI()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", s.IP, s.Port),
		Handler: s.Router,
	}

	go func(exitChan chan struct{}) {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println("ListenAndServe(): %v", err)
		}
	}(s.Exit)

	s.Srv = srv
	return srv
}

func (s *Server) StopHTTPServer() {
	log.Debug("stopping HTTP server")
	defer s.Srv.Close()
	close(s.Exit)
}

func (s *Server) HandleGetRecommendations(w http.ResponseWriter, r *http.Request) {
	cityName := chi.URLParam(r, "city")
	if cityName == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	geom, err := s.CityDB.GetGeometryByCityName(cityName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if len(geom) == 0 || geom == nil {
		emptyMrs := make([]string, 0)
		json.NewEncoder(w).Encode(emptyMrs)
	} else {
		w.Write(geom)
	}
}
