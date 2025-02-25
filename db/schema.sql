--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Ubuntu 14.5-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.7 (Ubuntu 14.7-0ubuntu0.22.04.1)

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

--
-- Name: deposit_requests; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.deposit_requests (
    pod_address character varying(42) NOT NULL,
    operator_address character varying(42) NOT NULL,
    transaction_id character varying(64) NOT NULL,
    amount  varchar NOT NULL,
    archived boolean NOT NULL
);

ALTER TABLE public.deposit_requests OWNER TO postgres;

CREATE TABLE public.withdraw_requests (
    PodAddress VARCHAR(42) NOT NULL, -- Ethereum address of the pod
    OperatorAddress VARCHAR(42) NOT NULL, -- Ethereum address of the operator
    TransactionID VARCHAR(64) NOT NULL, -- Transaction ID of the Bitcoin deposit
    WithdrawAddr BYTEA NOT NULL, -- Amount of the deposit
    Archived BOOLEAN NOT NULL -- Status of the deposit request
);

ALTER TABLE public.withdraw_requests OWNER TO postgres;

--
-- Name: multi_sig_address; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.multi_sig_address (
    address character varying(255) NOT NULL,
    script text NOT NULL,
    podaddress character varying(255),
    signed boolean NOT NULL,
    archived boolean NOT NULL
);