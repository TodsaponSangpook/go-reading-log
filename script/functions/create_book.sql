CREATE OR REPLACE PROCEDURE public.create_book(
	IN p_user_id INT,
    IN p_title VARCHAR,
    IN p_author VARCHAR,
    IN p_category VARCHAR,
    IN p_status VARCHAR,
    IN p_rating INT,
    IN p_review VARCHAR
)
LANGUAGE plpgsql
AS $procedure$
	BEGIN
		INSERT INTO books(
			user_id, title, author, category,
        	status, rating, review, created_at
		)
		VALUES (
			p_user_id, p_title, p_author, p_category,
        	p_status, p_rating, p_review, NOW()
		);
	END;
$procedure$
;