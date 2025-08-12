CREATE OR REPLACE PROCEDURE public.insert_refresh_token(IN p_user_id integer, IN p_token character varying, IN p_expires_at timestamp with time zone)
 LANGUAGE plpgsql
AS $procedure$
BEGIN
	INSERT INTO refresh_token(user_id, token, expires_at)
	VALUES (p_user_id, p_token, p_expires_at);
END;
$procedure$
;