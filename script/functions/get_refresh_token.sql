CREATE OR REPLACE FUNCTION public.get_refresh_token(p_user_id integer, p_token character varying)
 RETURNS TABLE(expires_at timestamp without time zone)
 LANGUAGE plpgsql
AS $function$
BEGIN
	RETURN QUERY
	SELECT rt.expires_at
	FROM refresh_token rt
	WHERE rt.user_id = p_user_id and rt.token = p_token;

	IF NOT FOUND THEN
		RAISE EXCEPTION 'Refresh token with userID "%" not found', p_user_id;
	END IF;
END;
$function$
;