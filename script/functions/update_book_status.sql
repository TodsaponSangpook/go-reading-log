CREATE OR REPLACE PROCEDURE public.update_book_status(IN p_id integer, IN p_user_id integer, IN p_status text)
 LANGUAGE plpgsql
AS $procedure$
BEGIN

    IF NOT EXISTS (SELECT 1 FROM books WHERE id = p_id AND user_id = p_user_id) THEN
        RAISE EXCEPTION 'Book ID % not found', p_id;
    END IF;

    IF p_status NOT IN ('Unread', 'Reading', 'Done') THEN
        RAISE EXCEPTION 'Invalid status: %. Allowed: Unread, Reading, Done.', p_status;
    END IF;

	UPDATE books
    SET 
        status = p_status,
        started_at = CASE 
            WHEN p_status = 'Reading' THEN NOW() 
            ELSE started_at 
        END,
        finished_at = CASE 
            WHEN p_status = 'Done' THEN NOW() 
            ELSE finished_at 
        END
    WHERE id = p_id AND user_id = p_user_id;

END;
$procedure$
;