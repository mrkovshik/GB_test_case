--
-- PostgreSQL database dump
--

-- Dumped from database version 14.6 (Ubuntu 14.6-1.pgdg22.04+1)
-- Dumped by pg_dump version 15.1 (Ubuntu 15.1-1.pgdg22.04+1)



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
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

-- *not* creating schema, since initdb creates it


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: job_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE IF NOT EXISTS public.job_types (
    id integer NOT NULL,
    job_type text NOT NULL
);


--
-- Name: job_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.job_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: job_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.job_types_id_seq OWNED BY public.job_types.id;


--
-- Name: vacancies; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE IF NOT EXISTS public.vacancies (
    id integer NOT NULL,
    vacancy_name text NOT NULL,
    key_skills text NOT NULL,
    salary integer NOT NULL,
    job_type integer,
    vacancy_desc text NOT NULL
);


--
-- Name: vacancies_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.vacancies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: vacancies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.vacancies_id_seq OWNED BY public.vacancies.id;


--
-- Name: job_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.job_types ALTER COLUMN id SET DEFAULT nextval('public.job_types_id_seq'::regclass);


--
-- Name: vacancies id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.vacancies ALTER COLUMN id SET DEFAULT nextval('public.vacancies_id_seq'::regclass);


--
-- Data for Name: job_types; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.job_types (id, job_type) FROM stdin;
1	в офисе
2	дистанционно
3	гибридный тип
\.


--
-- Data for Name: vacancies; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.vacancies (id, vacancy_name, key_skills, salary, job_type, vacancy_desc) FROM stdin;
1	Уборщик	мыть полы, чистить ковры	300000	1	Человек, который моет полы и чистит ковры
2	Охранник	разгадывать сканворды, следить за порядком	300000	1	Человек, который следит за безопасностью и порядком в компании
3	Медсестра	дуть на ранки, мазать вавки	50000	2	Человек, который спасает жизни с помощью медицинских штук
\.


--
-- Name: job_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.job_types_id_seq', 3, true);


--
-- Name: vacancies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.vacancies_id_seq', 4, true);


--
-- Name: job_types job_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.job_types
    ADD CONSTRAINT job_types_pkey PRIMARY KEY (id);


--
-- Name: vacancies vacancies_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.vacancies
    ADD CONSTRAINT vacancies_pkey PRIMARY KEY (id);


--
-- Name: vacancies vacancies_job_type_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.vacancies
    ADD CONSTRAINT vacancies_job_type_fkey FOREIGN KEY (job_type) REFERENCES public.job_types(id);


--
-- PostgreSQL database dump complete
--

