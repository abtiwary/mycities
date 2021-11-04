package mycitiesdb

import (
	"fmt"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Cities struct {
	Geometry string `db:"geom"`
}

type CitiesDB struct {
	DB *sqlx.DB
}

func NewCitiesDatabase(
	host string,
	port int,
	dbname string,
	dbuser string,
	dbpass string) (*CitiesDB, error) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port,
		dbuser, dbpass, dbname)
	db, err := sqlx.Connect("postgres", dbinfo)
	if err != nil {
		return nil, errors.Errorf("error creating database with info: %s", dbinfo)
	}

	return &CitiesDB{
		DB: db,
	}, nil
}

func (db *CitiesDB) GetGeometryByCityName(cityName string) ([]byte, error) {
	query := (fmt.Sprintf(
		"SELECT ST_AsGeoJSON(geom) as geom from cities WHERE name ~ %s LIMIT 1",
		pq.QuoteLiteral(cityName),
	))

	var cities Cities
	err := db.DB.Get(&cities, query)
	if err != nil {
		log.WithError(err).Debugf("error error scanning row")
		return nil, err
	}

	return []byte(cities.Geometry), nil
}
