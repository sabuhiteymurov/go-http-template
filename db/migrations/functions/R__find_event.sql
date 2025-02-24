CREATE OR REPLACE FUNCTION public.find_event(_event_id INTEGER)
    RETURNS TABLE (
        id INTEGER, name VARCHAR,
        description VARCHAR,
        location VARCHAR,
        created_at TIMESTAMP WITHOUT TIME ZONE,
        user_id INTEGER
    )
LANGUAGE plpgsql AS
$$
    BEGIN
    RETURN QUERY SELECT
        event.id,
        event.name,
        event.description,
        event.location,
        event.created_at,
        event.user_id
    FROM public.events as event WHERE event.id = _event_id LIMIT 1;
    END;
$$;