CREATE OR REPLACE FUNCTION public.get_books_by_user(p_user_id INTEGER)
RETURNS TABLE (
    id INT,
    title VARCHAR,
    author VARCHAR,
    category VARCHAR,
    status VARCHAR,
    rating INT,
    review VARCHAR,
    started_at TIMESTAMP,
    finished_at TIMESTAMP,
    created_at TIMESTAMP
)
LANGUAGE plpgsql
AS $function$
BEGIN
	RETURN QUERY
	SELECT
		b.id,
        b.title,
        b.author,
        b.category,
        b.status,
        b.rating,
        b.review,
        b.started_at,
        b.finished_at,
        b.created_at
    FROM books b
    WHERE b.user_id = p_user_id;
END;
$function$
;