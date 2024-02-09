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


CREATE DOMAIN public.http_client_error_code AS smallint
	CONSTRAINT http_client_error_code_check CHECK (VALUE = 0 OR VALUE >= 400 AND VALUE <= 499);

CREATE DOMAIN public.http_client_error_code_list AS http_client_error_code[];

CREATE TABLE public.test_array_with_go (
	supported_error_codes public.http_client_error_code_list NOT NULL
);


