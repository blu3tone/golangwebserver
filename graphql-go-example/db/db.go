package db

import (
	"database/sql"
	"log"

	"github.com/graph-gophers/graphql-go"
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

//const dbURL = "postgres://postgres@localhost:5432/postgres?sslmode=disable"
const dbURL = "postgres://postgres@localhost:5432/test?sslmode=disable"

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"test", driver)
	if err != nil {
		log.Println("Error occurred during migration", err)
	}
	m.Steps(5)

}

type User struct {
	ID       graphql.ID
	Email    string
	Password string
}

type Book struct {
	ID        graphql.ID
	Title     string
	Body      string
	Author_id graphql.ID
}

type Author struct {
	ID   graphql.ID
	Name string
}

func FindBook(id graphql.ID) *Book {
	book := &Book{}

	err := db.QueryRow("SELECT id, title, body FROM books WHERE id = $1", id).Scan(&book.ID, &book.Title, &book.Body)
	if err != nil {
		log.Println("A book does not exists!", err)
	}
	return book
}

func FindUser(id graphql.ID) *User {
	user := &User{}

	err := db.QueryRow("SELECT id, email, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		log.Println("User does not exist!", err)
	}
	return user
}

func FindBooksOfUser(id graphql.ID) []*Book {
	res, err := db.Query("SELECT b.id, title, body FROM user_books a JOIN  books b ON a.book_id = b.id WHERE a.user_id = $1 ", id)
	if err != nil {
		log.Println("The book of user does not exist!", err)
	}
	var books []*Book
	defer res.Close()
	var idx, title, body string
	for res.Next() {
		err := res.Scan(&idx, &title, &body)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, &Book{ID: graphql.ID(idx), Title: title, Body: body})

	}
	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
	return books

}

func InsertUser(u *User) error {
	sqlStatement := `INSERT INTO users VALUES($1, $2, $3)`

	_, err := db.Exec(sqlStatement, &u.ID, &u.Email, &u.Password)
	return err
}

func InsertRelUB(id, user_id, book_id string) error {
	sqlStatement := `INSERT INTO user_books VALUES($1, $2, $3)`
	_, err := db.Exec(sqlStatement, id, user_id, book_id)
	return err
}

func DeleteUser(id string) error {
	sqlStatement := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(sqlStatement, id)
	return err
}

func UpdateUser(id, password string) (*User, error) {
	user := &User{}
	sqlStatement := `UPDATE users SET password = $1 WHERE id = $2`
	_, err := db.Exec(sqlStatement, password, id)
	user = FindUser(graphql.ID(id))
	return user, err

}

func FindAuthorBook(id graphql.ID) *Author {
	author := &Author{}

	err := db.QueryRow("SELECT a.id, name FROM books b INNER JOIN authors a ON b.author_id = a.id WHERE b.id = $1", id).Scan(&author.ID, &author.Name)
	if err != nil {
		log.Println("Author does not exist!", err)
	}
	return author
}

func FindBooksOfAuthor(id graphql.ID) []*Book {
	res, err := db.Query("SELECT id, title, body FROM books WHERE author_id = $1 ", id)
	if err != nil {
		log.Println("The book of author does not exist!", err)
	}
	var books []*Book
	defer res.Close()
	var idx, title, body string
	for res.Next() {
		err := res.Scan(&idx, &title, &body)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, &Book{ID: graphql.ID(idx), Title: title, Body: body})

	}
	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
	return books

}

func FindAuthor(id graphql.ID) *Author {
	author := &Author{}

	err := db.QueryRow("SELECT id, name FROM authors WHERE id = $1", id).Scan(&author.ID, &author.Name)
	if err != nil {
		log.Println("Author does not exist!", err)
	}
	return author
}

func InsertAuthor(a *Author) error {
	sqlStatement := `INSERT INTO authors VALUES($1, $2)`

	_, err := db.Exec(sqlStatement, &a.ID, &a.Name)
	return err
}

func InsertBook(b *Book) error {
	sqlStatement := `INSERT INTO books VALUES($1, $2, $3, $4)`

	_, err := db.Exec(sqlStatement, &b.ID, &b.Title, &b.Body, &b.Author_id)
	return err
}
