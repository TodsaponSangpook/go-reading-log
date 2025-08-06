CREATE OR REPLACE PROCEDURE public.update_book(
	p_id		INT,
	p_user_id	INT,
 	p_title 	TEXT DEFAULT NULL,
    p_author 	TEXT DEFAULT NULL,
    p_category 	TEXT DEFAULT NULL,
    p_rating 	INT DEFAULT NULL,
    p_review 	TEXT DEFAULT NULL
)
LANGUAGE plpgsql
AS $procedure$
BEGIN
	UPDATE books
	SET
		title = COALESCE(p_title, title),
		author = COALESCE(p_author, author),
        category = COALESCE(p_category, category),
        rating = COALESCE(p_rating, rating),
        review = COALESCE(p_review, review)
	WHERE id = p_id AND user_id = p_user_id;
END;
$procedure$
;
