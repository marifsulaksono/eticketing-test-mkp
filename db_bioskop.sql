--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Ubuntu 16.2-1.pgdg22.04+1)
-- Dumped by pg_dump version 16.2 (Ubuntu 16.2-1.pgdg22.04+1)

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
-- Name: role_users; Type: TYPE; Schema: public; Owner: developer
--

CREATE TYPE public.role_users AS ENUM (
    'admin',
    'customer'
);


ALTER TYPE public.role_users OWNER TO developer;

--
-- Name: seats_status; Type: TYPE; Schema: public; Owner: developer
--

CREATE TYPE public.seats_status AS ENUM (
    'available',
    'reserved',
    'occupied'
);


ALTER TYPE public.seats_status OWNER TO developer;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: branchs; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.branchs (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    address text NOT NULL,
    phonenumber character varying(30) NOT NULL
);


ALTER TABLE public.branchs OWNER TO developer;

--
-- Name: branchs_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.branchs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.branchs_id_seq OWNER TO developer;

--
-- Name: branchs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.branchs_id_seq OWNED BY public.branchs.id;


--
-- Name: genres; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.genres (
    id integer NOT NULL,
    name character varying(50) NOT NULL
);


ALTER TABLE public.genres OWNER TO developer;

--
-- Name: genres_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.genres_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.genres_id_seq OWNER TO developer;

--
-- Name: genres_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.genres_id_seq OWNED BY public.genres.id;


--
-- Name: movie_genres; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.movie_genres (
    movie_id integer NOT NULL,
    genre_id integer NOT NULL
);


ALTER TABLE public.movie_genres OWNER TO developer;

--
-- Name: movies; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.movies (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    duration time without time zone NOT NULL,
    release_date date NOT NULL
);


ALTER TABLE public.movies OWNER TO developer;

--
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.movies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.movies_id_seq OWNER TO developer;

--
-- Name: movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;


--
-- Name: seats; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.seats (
    id integer NOT NULL,
    stage_id integer NOT NULL,
    number integer NOT NULL
);


ALTER TABLE public.seats OWNER TO developer;

--
-- Name: seats_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.seats_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.seats_id_seq OWNER TO developer;

--
-- Name: seats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.seats_id_seq OWNED BY public.seats.id;


--
-- Name: showtimes; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.showtimes (
    id integer NOT NULL,
    branch_id integer NOT NULL,
    stage_id integer NOT NULL,
    movie_id integer NOT NULL,
    start_time timestamp without time zone NOT NULL,
    end_time timestamp without time zone NOT NULL
);


ALTER TABLE public.showtimes OWNER TO developer;

--
-- Name: showtimes_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.showtimes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.showtimes_id_seq OWNER TO developer;

--
-- Name: showtimes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.showtimes_id_seq OWNED BY public.showtimes.id;


--
-- Name: stages; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.stages (
    id integer NOT NULL,
    branch_id integer NOT NULL,
    name character varying(100) NOT NULL,
    capacity integer NOT NULL
);


ALTER TABLE public.stages OWNER TO developer;

--
-- Name: stages_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.stages_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.stages_id_seq OWNER TO developer;

--
-- Name: stages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.stages_id_seq OWNED BY public.stages.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(100) NOT NULL,
    password text NOT NULL,
    email character varying(255) NOT NULL,
    phonenumber character varying(30) NOT NULL,
    firstname character varying(255) NOT NULL,
    lastname character varying(255) NOT NULL,
    role public.role_users
);


ALTER TABLE public.users OWNER TO developer;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO developer;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: branchs id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.branchs ALTER COLUMN id SET DEFAULT nextval('public.branchs_id_seq'::regclass);


--
-- Name: genres id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.genres ALTER COLUMN id SET DEFAULT nextval('public.genres_id_seq'::regclass);


--
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);


--
-- Name: seats id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.seats ALTER COLUMN id SET DEFAULT nextval('public.seats_id_seq'::regclass);


--
-- Name: showtimes id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.showtimes ALTER COLUMN id SET DEFAULT nextval('public.showtimes_id_seq'::regclass);


--
-- Name: stages id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.stages ALTER COLUMN id SET DEFAULT nextval('public.stages_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: branchs; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.branchs (id, name, address, phonenumber) FROM stdin;
1	MKP Wijaya Cinema	Semarang	081234567890
2	Mitra Kasih Cinema	Yogyakarta	081234839450
\.


--
-- Data for Name: genres; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.genres (id, name) FROM stdin;
1	sci-fi
2	romance
3	technology
\.


--
-- Data for Name: movie_genres; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.movie_genres (movie_id, genre_id) FROM stdin;
1	3
2	2
2	3
3	1
3	3
\.


--
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.movies (id, title, duration, release_date) FROM stdin;
1	Sillicon Valey	01:48:12	2022-02-19
2	Cintaku pada seorang Programmer	01:23:48	2023-12-08
3	WALL-E	02:05:12	2021-10-23
\.


--
-- Data for Name: seats; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.seats (id, stage_id, number) FROM stdin;
1	1	1
2	1	2
3	1	3
4	1	4
5	1	5
6	2	1
7	1	6
8	1	7
9	3	1
10	2	2
11	1	8
12	3	2
\.


--
-- Data for Name: showtimes; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.showtimes (id, branch_id, stage_id, movie_id, start_time, end_time) FROM stdin;
1	1	1	2	2024-03-17 09:00:00	2024-03-17 11:00:00
2	1	2	1	2024-03-17 08:15:00	2024-03-17 10:15:00
3	1	1	3	2024-03-17 10:25:00	2024-03-17 12:00:00
\.


--
-- Data for Name: stages; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.stages (id, branch_id, name, capacity) FROM stdin;
1	1	Mitra Room	20
2	1	Kasih Room	15
3	1	Perkasa Room	10
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.users (id, username, password, email, phonenumber, firstname, lastname, role) FROM stdin;
1	admin	$2y$10$6bK9ehBhQ9T2zfvEmklJ8uLGj6inVmMG5k.tlUvyO1LDWn9QY/q9e	marifsulaksono@gmail.com	085331828197	Muhammad Ari	Sulaksono	\N
\.


--
-- Name: branchs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.branchs_id_seq', 2, true);


--
-- Name: genres_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.genres_id_seq', 3, true);


--
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.movies_id_seq', 3, true);


--
-- Name: seats_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.seats_id_seq', 12, true);


--
-- Name: showtimes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.showtimes_id_seq', 3, true);


--
-- Name: stages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.stages_id_seq', 3, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: branchs branchs_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.branchs
    ADD CONSTRAINT branchs_pkey PRIMARY KEY (id);


--
-- Name: genres genres_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.genres
    ADD CONSTRAINT genres_pkey PRIMARY KEY (id);


--
-- Name: movie_genres movie_genres_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.movie_genres
    ADD CONSTRAINT movie_genres_pkey PRIMARY KEY (movie_id, genre_id);


--
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- Name: seats seats_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_pkey PRIMARY KEY (id);


--
-- Name: showtimes showtimes_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_pkey PRIMARY KEY (id);


--
-- Name: stages stages_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.stages
    ADD CONSTRAINT stages_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: stages fk_branchid; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.stages
    ADD CONSTRAINT fk_branchid FOREIGN KEY (branch_id) REFERENCES public.branchs(id);


--
-- Name: seats fk_stageid; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT fk_stageid FOREIGN KEY (stage_id) REFERENCES public.stages(id);


--
-- Name: movie_genres movie_genres_genre_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.movie_genres
    ADD CONSTRAINT movie_genres_genre_id_fkey FOREIGN KEY (genre_id) REFERENCES public.genres(id);


--
-- Name: movie_genres movie_genres_movie_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.movie_genres
    ADD CONSTRAINT movie_genres_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id);


--
-- Name: showtimes showtimes_branch_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_branch_id_fkey FOREIGN KEY (branch_id) REFERENCES public.branchs(id);


--
-- Name: showtimes showtimes_movie_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id);


--
-- Name: showtimes showtimes_stage_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_stage_id_fkey FOREIGN KEY (stage_id) REFERENCES public.stages(id);


--
-- PostgreSQL database dump complete
--

