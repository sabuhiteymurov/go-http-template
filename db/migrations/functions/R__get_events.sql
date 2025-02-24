CREATE OR REPLACE FUNCTION public.get_events()
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
    FROM public.events as event;
    END;
$$;