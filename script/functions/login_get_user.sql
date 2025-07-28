CREATE OR REPLACE FUNCTION public.login_get_user(p_email VARCHAR)
 RETURNS TABLE(id INTEGER, email VARCHAR, password_hash VARCHAR, name VARCHAR)
 LANGUAGE plpgsql
AS $function$
BEGIN
    RETURN QUERY
    SELECT u.id, u.email, u.password_hash, u.name
    FROM users u
    WHERE u.email = p_email;

	IF NOT FOUND THEN
		RAISE EXCEPTION 'User with email "%" not found', p_email;
	END IF;
END;
$function$
;