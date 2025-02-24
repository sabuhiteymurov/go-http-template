DROP PROCEDURE IF EXISTS public.create_event;

CREATE OR REPLACE PROCEDURE public.create_event(
    _name VARCHAR(255),
    _description VARCHAR,
    _location VARCHAR(255),
    _created_at TIMESTAMP WITHOUT TIME ZONE,
    _user_id INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO public.events (name, description, location, created_at, user_id)
    VALUES (_name, _description, _location, _created_at, _user_id);
END;
$$;