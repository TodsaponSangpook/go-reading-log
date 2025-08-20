CREATE OR REPLACE PROCEDURE public.delete_refresh_token_by_user()
	LANGUAGE plpgsql
AS $procedure$
BEGIN
	DELETE FROM refresh_token
	WHERE user_id = p_user_id;
END;
$procedure$;
;