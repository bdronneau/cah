--
-- PostgreSQL database dump
--

-- Dumped from database version 11.7
-- Dumped by pg_dump version 11.7

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

ALTER TABLE IF EXISTS ONLY public.users DROP CONSTRAINT IF EXISTS users_pkey;
ALTER TABLE IF EXISTS ONLY public.rooms DROP CONSTRAINT IF EXISTS rooms_pkey;
ALTER TABLE IF EXISTS ONLY public.rooms DROP CONSTRAINT IF EXISTS rooms_name_key;
ALTER TABLE IF EXISTS ONLY public.cards DROP CONSTRAINT IF EXISTS cards_pkey;
DROP TABLE IF EXISTS public.rooms_users_cards;
DROP TABLE IF EXISTS public.users;
DROP SEQUENCE IF EXISTS public.users_id_seq;
DROP TABLE IF EXISTS public.rooms_users;
DROP TABLE IF EXISTS public.rooms_cards;
DROP TABLE IF EXISTS public.rooms;
DROP SEQUENCE IF EXISTS public.rooms_id_seq;
DROP TABLE IF EXISTS public.cards;
DROP SEQUENCE IF EXISTS public.cards_id_seq;
--
-- Name: cards_id_seq; Type: SEQUENCE; Schema: public; Owner: cah
--

CREATE SEQUENCE public.cards_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


ALTER TABLE public.cards_id_seq OWNER TO cah;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: cards; Type: TABLE; Schema: public; Owner: cah
--

CREATE TABLE public.cards (
    id integer DEFAULT nextval('public.cards_id_seq'::regclass) NOT NULL,
    name text NOT NULL,
    type text NOT NULL
);


ALTER TABLE public.cards OWNER TO cah;

--
-- Name: rooms_id_seq; Type: SEQUENCE; Schema: public; Owner: cah
--

CREATE SEQUENCE public.rooms_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


ALTER TABLE public.rooms_id_seq OWNER TO cah;

--
-- Name: rooms; Type: TABLE; Schema: public; Owner: cah
--

CREATE TABLE public.rooms (
    id integer DEFAULT nextval('public.rooms_id_seq'::regclass) NOT NULL,
    name text,
    description text NOT NULL,
    turn integer DEFAULT 1,
    status text NOT NULL,
    lastupdated timestamp without time zone NOT NULL
);


ALTER TABLE public.rooms OWNER TO cah;

--
-- Name: rooms_cards; Type: TABLE; Schema: public; Owner: cah
--

CREATE TABLE public.rooms_cards (
    room_id integer NOT NULL,
    card_id integer NOT NULL,
    turn integer DEFAULT 1,
    used boolean DEFAULT false
);


ALTER TABLE public.rooms_cards OWNER TO cah;

--
-- Name: rooms_users; Type: TABLE; Schema: public; Owner: cah
--

CREATE TABLE public.rooms_users (
    room_id integer NOT NULL,
    user_id integer NOT NULL,
    enabled boolean DEFAULT true,
    judge boolean DEFAULT false
);


ALTER TABLE public.rooms_users OWNER TO cah;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: cah
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO cah;

--
-- Name: users; Type: TABLE; Schema: public; Owner: cah
--

CREATE TABLE public.users (
    id integer DEFAULT nextval('public.users_id_seq'::regclass) NOT NULL,
    name text NOT NULL,
    lastupdated timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO cah;

--
-- Name: rooms_users_cards; Type: TABLE; Schema: public; Owner: cah
--

CREATE TABLE public.rooms_users_cards (
    room_id integer NOT NULL,
    user_id integer NOT NULL,
    card_id integer NOT NULL,
    used boolean DEFAULT false,
    on_hand boolean DEFAULT false,
    turn integer DEFAULT 0,
    vote integer DEFAULT 0
);


ALTER TABLE public.rooms_users_cards OWNER TO cah;

--
-- Name: cards cards_pkey; Type: CONSTRAINT; Schema: public; Owner: cah
--

ALTER TABLE ONLY public.cards
    ADD CONSTRAINT cards_pkey PRIMARY KEY (id);


--
-- Name: rooms rooms_name_key; Type: CONSTRAINT; Schema: public; Owner: cah
--

ALTER TABLE ONLY public.rooms
    ADD CONSTRAINT rooms_name_key UNIQUE (name);


--
-- Name: rooms rooms_pkey; Type: CONSTRAINT; Schema: public; Owner: cah
--

ALTER TABLE ONLY public.rooms
    ADD CONSTRAINT rooms_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: cah
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

