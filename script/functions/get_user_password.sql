CREATE OR REPLACE FUNCTION public.get_user_password(p_user_id integer)
	RETURNS TABLE(password_hash character varying)
	LANGUAGE plpgsql
AS $function$
BEGIN
	RETURN QUERY
	SELECT u.password_hash
	FROM users u
	WHERE u.id = p_user_id;

	IF NOT FOUND THEN
		RAISE EXCEPTION 'Not found';
	END IF;
END;
$function$
;