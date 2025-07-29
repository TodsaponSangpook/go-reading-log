CREATE OR REPLACE FUNCTION public.get_user_profile(p_user_id INTEGER)
RETURNS TABLE (
    id INT,
    email VARCHAR,
    name VARCHAR,
    created_at TIMESTAMP
)
LANGUAGE plpgsql
AS $function$
BEGIN
	RETURN QUERY
	SELECT
		u.id,
        u.email,
        u.name,
        u.created_at
    FROM users u
    WHERE u.id = p_user_id;
END;
$function$
;