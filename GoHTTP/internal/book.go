package model

type Book struct {
	Base
	Title    string
	AuthorID int
	Author   Author // has one relation
	Editor int
	Genres       []Genre       `json:"genres,omitempty"` // many to many relation
	User 		[]User		`json:"users,omitempty"`
}


type UserBook struct{
	BookId  int `json:"book_id"` // pk tag is used to mark field as primary key
	UserID int `json:"user_id"`
}

type BookGenre struct {
	BookId  int `json:"book_id"` // pk tag is used to mark field as primary key
	GenreId int `json:"genre_id"`
}



type Genre struct {

	Base
	Name   string
	Rating int `json:"-"` // - is used to ignore field

	Books []Book `json:"books, omitempty"` // many to many reblation
}
