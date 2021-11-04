--
-- PostgreSQL database dump
--

-- Dumped from database version 13.4 (Debian 13.4-4.pgdg110+1)
-- Dumped by pg_dump version 13.4 (Debian 13.4-4.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: postgis; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS postgis WITH SCHEMA public;


--
-- Name: EXTENSION postgis; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION postgis IS 'PostGIS geometry and geography spatial types and functions';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: cities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cities (
    id bigint NOT NULL,
    name character varying,
    geom public.geometry
);


ALTER TABLE public.cities OWNER TO postgres;

--
-- Name: cities_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cities_id_seq OWNER TO postgres;

--
-- Name: cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cities_id_seq OWNED BY public.cities.id;


--
-- Name: cities id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities ALTER COLUMN id SET DEFAULT nextval('public.cities_id_seq'::regclass);


--
-- Data for Name: cities; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cities (id, name, geom) FROM stdin;
1	new delhi	010300000001000000050000000000005A8B4353409CF7506BF8953C4000000092E94353409CF7506BF8953C4000000092E9435340357FE0962F973C400000005A8B435340357FE0962F973C400000005A8B4353409CF7506BF8953C40
2	darjeeling	010300000001000000050000000000009AC6105640DB73D5116B073B40000000A9E4105640DB73D5116B073B40000000A9E410564014591E6FDC073B400000009AC610564014591E6FDC073B400000009AC6105640DB73D5116B073B40
3	jakarta	01030000000100000005000000FFFF7FF9F4B25A4085FDFC2079F418C0000000D016B35A4085FDFC2079F418C0000000D016B35A405AAF6C48E8F018C0FFFF7FF9F4B25A405AAF6C48E8F018C0FFFF7FF9F4B25A4085FDFC2079F418C0
4	ryde	010300000001000000050000000000004876E36240CCF44F0F55E840C000000088AEE36240CCF44F0F55E840C000000088AEE36240A49CDC7BB1E740C00000004876E36240A49CDC7BB1E740C00000004876E36240CCF44F0F55E840C0
\.


--
-- Data for Name: spatial_ref_sys; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.spatial_ref_sys (srid, auth_name, auth_srid, srtext, proj4text) FROM stdin;
\.


--
-- Name: cities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cities_id_seq', 4, true);


--
-- Name: cities cities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

