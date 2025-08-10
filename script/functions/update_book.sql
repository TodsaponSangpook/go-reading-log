CREATE OR REPLACE PROCEDURE public.update_book(IN p_id integer, IN p_user_id integer, IN p_title text DEFAULT NULL::text, IN p_author text DEFAULT NULL::text, IN p_category text DEFAULT NULL::text, IN p_rating integer DEFAULT NULL::integer, IN p_review text DEFAULT NULL::text)
 LANGUAGE plpgsql
AS $procedure$
BEGIN
  	IF NOT EXISTS (SELECT 1 FROM books WHERE id = p_id AND user_id = p_user_id) THEN
        RAISE EXCEPTION 'Book ID % not found', p_id;
    END IF;

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
