CREATE OR REPLACE FUNCTION public.get_user_profile(p_user_id integer)
 RETURNS TABLE(id integer, email character varying, name character varying, created_at timestamp without time zone)
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

	IF NOT FOUND THEN
		RAISE EXCEPTION 'User with email "%" not found', p_email;
	END IF;
END;
$function$
;