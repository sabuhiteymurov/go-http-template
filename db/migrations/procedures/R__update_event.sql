DROP PROCEDURE IF EXISTS public.update_event;

CREATE OR REPLACE PROCEDURE public.update_event(
    _event_id INTEGER,
    _name VARCHAR(255) default null,
    _description VARCHAR default null,
    _location VARCHAR(255) default null
)
    LANGUAGE plpgsql AS
$$
BEGIN
    UPDATE public.events
    SET name        = COALESCE(_name, name),
        location    = COALESCE(_location, location),
        description = COALESCE(_description, description)
    WHERE id = _event_id;
END;
$$;