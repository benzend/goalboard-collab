-- PostgreSQL database dump
-- Dumped from database version 16.2
-- Dumped by pg_dump version 16.2
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
SET default_tablespace = '';
SET default_table_access_method = heap;
-- Name: activity; Type: TABLE; Schema: public; Owner: postgres
CREATE TABLE public.activity (
    id integer NOT NULL,
    progress double precision,
    goal_id integer
);
ALTER TABLE public.activity OWNER TO postgres;
-- Name: activity_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
CREATE SEQUENCE public.activity_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.activity_id_seq OWNER TO postgres;
-- Name: activity_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
ALTER SEQUENCE public.activity_id_seq OWNED BY public.activity.id;
-- Name: goal; Type: TABLE; Schema: public; Owner: postgres
CREATE TABLE public.goal (
    id integer NOT NULL,
    name character varying(50),
    target_per_day character varying,
    long_term_target character varying,
    user_id integer
);
ALTER TABLE public.goal OWNER TO postgres;
-- Name: goal_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
CREATE SEQUENCE public.goal_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.goal_id_seq OWNER TO postgres;
-- Name: goal_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
ALTER SEQUENCE public.goal_id_seq OWNED BY public.goal.id;
-- Name: migration; Type: TABLE; Schema: public; Owner: postgres
CREATE TABLE public.migration (
    id integer NOT NULL,
    filename character varying(120) NOT NULL
);
ALTER TABLE public.migration OWNER TO postgres;
-- Name: migration_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
CREATE SEQUENCE public.migration_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.migration_id_seq OWNER TO postgres;
-- Name: migration_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
ALTER SEQUENCE public.migration_id_seq OWNED BY public.migration.id;
-- Name: user_; Type: TABLE; Schema: public; Owner: postgres
CREATE TABLE public.user_ (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    password character varying(128) NOT NULL
);
ALTER TABLE public.user_ OWNER TO postgres;
-- Name: user__id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
CREATE SEQUENCE public.user__id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.user__id_seq OWNER TO postgres;
-- Name: user__id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
ALTER SEQUENCE public.user__id_seq OWNED BY public.user_.id;
-- Name: activity id; Type: DEFAULT; Schema: public; Owner: postgres
ALTER TABLE ONLY public.activity ALTER COLUMN id SET DEFAULT nextval('public.activity_id_seq'::regclass);
-- Name: goal id; Type: DEFAULT; Schema: public; Owner: postgres
ALTER TABLE ONLY public.goal ALTER COLUMN id SET DEFAULT nextval('public.goal_id_seq'::regclass);
-- Name: migration id; Type: DEFAULT; Schema: public; Owner: postgres
ALTER TABLE ONLY public.migration ALTER COLUMN id SET DEFAULT nextval('public.migration_id_seq'::regclass);
-- Name: user_ id; Type: DEFAULT; Schema: public; Owner: postgres
ALTER TABLE ONLY public.user_ ALTER COLUMN id SET DEFAULT nextval('public.user__id_seq'::regclass);
-- Name: activity activity_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
ALTER TABLE ONLY public.activity
    ADD CONSTRAINT activity_pkey PRIMARY KEY (id);
-- Name: goal goal_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
ALTER TABLE ONLY public.goal
    ADD CONSTRAINT goal_pkey PRIMARY KEY (id);
-- Name: migration migration_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
ALTER TABLE ONLY public.migration
    ADD CONSTRAINT migration_pkey PRIMARY KEY (id);
-- Name: user_ user__pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
ALTER TABLE ONLY public.user_
    ADD CONSTRAINT user__pkey PRIMARY KEY (id);
-- Name: user_ user__username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
ALTER TABLE ONLY public.user_
    ADD CONSTRAINT user__username_key UNIQUE (username);
-- Name: activity goal_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
ALTER TABLE ONLY public.activity
    ADD CONSTRAINT goal_fk FOREIGN KEY (goal_id) REFERENCES public.goal(id);
-- Name: goal user_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
ALTER TABLE ONLY public.goal
    ADD CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES public.user_(id);
-- PostgreSQL database dump complete
-- PostgreSQL database dump
-- Dumped from database version 16.2
-- Dumped by pg_dump version 16.2
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
-- Data for Name: migration; Type: TABLE DATA; Schema: public; Owner: postgres
COPY public.migration (id, filename) FROM stdin;
1	20240511171546_create_user_table.sql
2	20240511171641_create_goal_table.sql
3	20240511171704_create_activity_table.sql
\.
-- Name: migration_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
SELECT pg_catalog.setval('public.migration_id_seq', 3, true);
-- PostgreSQL database dump complete
