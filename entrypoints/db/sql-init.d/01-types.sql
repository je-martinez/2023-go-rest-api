USE main_db;


CREATE TYPE IF NOT EXISTS public.sign_in_provider_type AS ENUM (
    'email',
    'google',
    'apple');

CREATE TYPE IF NOT EXISTS public.reaction_type AS ENUM (
    'love');