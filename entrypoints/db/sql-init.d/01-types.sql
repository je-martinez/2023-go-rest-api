USE main_db;

CREATE TYPE public.sign_in_provider_type AS ENUM (
    'email',
    'google',
    'apple');