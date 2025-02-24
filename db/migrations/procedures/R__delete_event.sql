DROP PROCEDURE IF EXISTS public.delete_event;

CREATE
    OR REPLACE PROCEDURE public.delete_event(
    _event_id INTEGER
)
    LANGUAGE plpgsql
AS
$$
BEGIN
    DELETE
    FROM public.events
    WHERE id = _event_id;
END;
$$;