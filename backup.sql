--
-- PostgreSQL database dump
--

-- Dumped from database version 10.3
-- Dumped by pg_dump version 10.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: userbasic; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.userbasic (
    id bigint NOT NULL,
    uuid uuid NOT NULL,
    first_name name NOT NULL,
    middle_name name,
    last_name name NOT NULL,
    date_birth date,
    email character varying(225) NOT NULL,
    primary_contact character varying(13) NOT NULL,
    secondary_contact character varying(10),
    pincode character varying(6),
    adress text,
    state character varying(50),
    password text NOT NULL,
    registerd_as character(6),
    created_at timestamp with time zone NOT NULL,
    is_blocked boolean NOT NULL,
    pending_details boolean NOT NULL,
    alias name,
    google_id character varying(255),
    facebook_id character varying(255),
    otp_expiry timestamp with time zone,
    wrong_pwd_count integer,
    time_till_block timestamp with time zone,
    user_agent character varying(255),
    refresh_token text,
    access_token text,
    is_verified boolean NOT NULL,
    current_otp character varying(6)
);


ALTER TABLE public.userbasic OWNER TO postgres;

--
-- Name: userbasic_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.userbasic ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.userbasic_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: userbasic; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.userbasic (id, uuid, first_name, middle_name, last_name, date_birth, email, primary_contact, secondary_contact, pincode, adress, state, password, registerd_as, created_at, is_blocked, pending_details, alias, google_id, facebook_id, otp_expiry, wrong_pwd_count, time_till_block, user_agent, refresh_token, access_token, is_verified, current_otp) FROM stdin;
44	9bed4f3a-916b-4571-979f-b56eee65609b	sajan		jaiswal	\N	sjnjaiswal@gmail.com	+917064274923					$2a$04$v5MwOTQrjCXRdpEuEqUmvOzPqtWQ0ElYVdX.B.UqztEL9YWthEZ1.	      	2020-05-28 17:58:56+00	f	t				0001-01-01 00:00:00+00	0	0001-01-01 00:00:00+00		eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJzam5qYWlzd2FsQGdtYWlsLmNvbSIsImV4cCI6MTU5MDc5MTEzMiwiaWF0IjoxNTkwNjkxMTMyLCJpc3MiOiJUZXNsYSIsIm5iZiI6MTU5MDY5MTEzMn0.FOilINesav7eNREvaocU5I397Epf1vMRBeCB-hti6Bw	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJzam5qYWlzd2FsQGdtYWlsLmNvbSIsImV4cCI6MTU5MTY5MTEzMiwiaWF0IjoxNTkwNjkxMTMyLCJpc3MiOiJUZXNsYSIsIm5iZiI6MTU5MDY5MTEzMn0.IyCuhiN267GBP0Y3PFu49LldmaRAyTSj-_MxO-WXUuw	t	
\.


--
-- Name: userbasic_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.userbasic_id_seq', 44, true);


--
-- Name: userbasic userbasic_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userbasic
    ADD CONSTRAINT userbasic_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

