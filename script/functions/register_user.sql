CREATE OR REPLACE PROCEDURE public.register_user(IN p_email VARCHAR, IN p_password_hash VARCHAR, IN p_name VARCHAR)
 LANGUAGE plpgsql
AS $procedure$
BEGIN
    IF EXISTS (SELECT 1 FROM users WHERE email = p_email) THEN
        RAISE EXCEPTION 'Email "%" already exists.', p_email;
    END IF;

    INSERT INTO users (email, password_hash, name, created_at)
    VALUES (p_email, p_password_hash, p_name, NOW());
END;
$procedure$
;