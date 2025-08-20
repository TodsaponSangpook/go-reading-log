CREATE OR REPLACE PROCEDURE public.update_user_password(p_user_id integer, p_new_password_hash character)
	LANGUAGE plpgsql
AS $procedure$
BEGIN
	UPDATE users
	SET password_hash = p_new_password_hash
	WHERE id = p_user_id;
END;
$procedure$
;