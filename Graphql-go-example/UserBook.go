package main

type UserBook struct{
	ID int
	Book_ID int
	User_ID int
}

func InsertUserBook( userbook *UserBook) error{
	var id int
	err := db.QueryRow(`
		INSERT INTO user_books(book_id, user_id)
		VALUES ($1, $2)
		RETURNING id
	`, userbook.Book_ID, userbook.User_ID).Scan(&id)
	if err != nil {
		return err
	}
	userbook.ID = id
	return nil
}