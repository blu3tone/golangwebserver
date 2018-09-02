package main

type Book struct{
	ID 		int
	Title	string
	Body	string
	Author_ID 	int
}


func GetBookByID(id int) (*Book, error){
	var title,  body string
	var author_id int
	err := db.QueryRow(`SELECT title, body, author_id FROM books WHERE id = $1`,
		id).Scan(&title, &body, &author_id)
	if err != nil {
		return nil, err
	}

	return &Book{
		ID: id,
		Title: title,
		Body:body,
		Author_ID: author_id,
	}, nil
}


func GetBookByIDAndUser(bookid, userid int) (*Book, error){
	var title, body string
	var authorid int
	err := db.QueryRow(`
		SELECT B.title, B.Body, B.author_id
		FROM user_books U JOIN books B 
			ON (B.id = U.book_id)
		WHERE U.user_id=$2
		ANd B.id = $1`,
		bookid, userid).Scan(&title, &body, &authorid)
	if err != nil {
		return nil, err
	}
	return &Book{
		ID:     bookid,
		Author_ID: authorid,
		Title:  title,
		Body:   body,
	}, nil


}

func GetBookByIDAndUse(userid int) ([]*Book, error){
	rows, err := db.Query(`
		SELECT B.id, B.title, B.Body, B.author_id
		FROM user_books U JOIN books B 
			ON (B.id = U.book_id)
		WHERE U.user_id=$1
	`, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		books       = []*Book{}
		bid, authorid        int
		title, body 		string
	)
	for rows.Next() {
		if err = rows.Scan(&bid, &title, &body, &authorid); err != nil {
			return nil, err
		}
		books = append(books, &Book{ID: bid, Title: title, Body: body, Author_ID: authorid})
	}
	return books, nil
}

func InsertBook(book *Book) error{
	var id int
	err := db.QueryRow(`
		INSERT INTO books(title, body, author_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`, book.Title, book.Body, book.Author_ID).Scan(&id)
	if err != nil {
		return err
	}
	book.ID = id
	return nil
}

func RemoveBookByID(id int) error{
	_, err := db.Exec("DELETE FROM books WHERE id=$1", id)
	return err
}
func GetAuthorIDBook(id int) (*Author, error){
	var name string
	err := db.QueryRow("SELECT name FROM authors WHERE id=$1", id).Scan(&name)
	if err != nil {
		return nil, err
	}
	return &Author{
		ID:    id,
		Name: name,

	}, nil
}