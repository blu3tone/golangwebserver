package bookstore

import (
	"github.com/blu3tone/golangwebserver/graphql-go-example/db"
	"log"

	"github.com/graph-gophers/graphql-go"
)

var Schema = `
	schema{
		query : Query
		mutation : Mutation
	}

	type Query {
		book(id : ID!): Book
		user(id : ID!): User
		author(id : ID!): Author
	}
	
	type Mutation{
		createUser(id: ID!, email: String!, password: String!): User
		deleteUser(id : ID!) : Boolean
		updateUser(id : ID!, password: String!): User
		
		createAuthor (id: ID!, name: String!): Author
		
		createBook(id : ID!, title: String!, body: String!, author_id: ID!): Book
		
		createRelUB(id: ID!, user_id: String!, book_id: String!):User

	}

	input UserInput{
		id : ID!
		email : String!
		password : String!
	}


	type User {
		id : ID!
		email :  String!
		password : String!
		books :[Book]
	}

	type Book{
		id : ID!
		title : String!
		body : String!
		author : Author
	}
	
	type Author{
		id : ID!
		name : String!
		books : [Book]
	}
	

`

type User struct {
	ID       graphql.ID
	Email    string
	Password string
	Books    []graphql.ID
}

type Book struct {
	ID     graphql.ID
	Title  string
	Body   string
	Author graphql.ID
}

type Author struct {
	ID    graphql.ID
	Name  string
	Books []graphql.ID
}

type UserInput struct {
	ID       graphql.ID
	Email    string
	Password string
}

type Resolver struct{}

func (r *Resolver) UpdateUser(args *struct {
	ID       string
	Password string
}) *userResolver {
	user, err := db.UpdateUser(args.ID, args.Password)

	if err != nil {
		return &userResolver{&User{}}
	}

	return &userResolver{&User{ID: user.ID, Email: user.Email, Password: user.Password}}

}

func (r *Resolver) DeleteUser(args *struct{ ID string }) *bool {
	err := db.DeleteUser(args.ID)
	b := true
	if err != nil {
		b = false
	}

	return &b
}

func (r *Resolver) CreateRelUB(args *struct {
	ID      string
	User_id string
	Book_id string
}) *userResolver {

	err := db.InsertRelUB(args.ID, args.User_id, args.Book_id)
	if err != nil {
		log.Printf("Can't insert transaction:", err)
	}

	return &userResolver{&User{ID: graphql.ID(args.User_id)}}

}

func (r *Resolver) CreateUser(args *struct {
	ID       string
	Email    string
	Password string
}) *userResolver {

	err := db.InsertUser(&db.User{ID: graphql.ID(args.ID), Email: args.Email, Password: args.Password})

	if err != nil {
		log.Printf("Can't insert user:", err)
		return &userResolver{&User{}}
	}
	return &userResolver{&User{ID: graphql.ID(args.ID),
		Email:    args.Email,
		Password: args.Password}}
}

func (r *Resolver) User(args struct{ ID graphql.ID }) *userResolver {
	var u *db.User
	u = db.FindUser(args.ID)
	if u == nil {
		log.Println("User does not exist!")
	}

	user := User{ID: u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
	return &userResolver{&user}
}

type userResolver struct {
	u *User
}

func (r *userResolver) ID() graphql.ID {
	return r.u.ID
}

func (r *userResolver) Email() string {
	return r.u.Email
}

func (r *userResolver) Password() string {
	return r.u.Password
}

func (r *userResolver) Books() *[]*bookResolve {
	return resolveUserBooks(r.u.ID)
}

func (r *Resolver) CreateBook(args struct {
	ID       graphql.ID
	Title    string
	Body     string
	AuthorID graphql.ID
}) *bookResolve {
	err := db.InsertBook(&db.Book{ID: graphql.ID(args.ID), Title: args.Title, Body: args.Body, Author_id: graphql.ID(args.AuthorID)})

	if err != nil {
		log.Printf("Can't insert author:", err)
		return &bookResolve{&Book{}}
	}
	return &bookResolve{&Book{ID: graphql.ID(args.ID),
		Title: args.Title,
		Body:  args.Body,
	}}
}

func resolveUserBooks(id graphql.ID) *[]*bookResolve {
	bks := db.FindBooksOfUser(id)
	var bookresolvers []*bookResolve
	for _, bk := range bks {
		var book Book
		book.ID = bk.ID
		book.Title = bk.Title
		book.Body = bk.Body
		bookresolvers = append(bookresolvers, &bookResolve{&book})
	}
	return &bookresolvers
}

type bookResolve struct {
	b *Book
}

func (r *bookResolve) ID() graphql.ID {
	return r.b.ID
}

func (r *bookResolve) Title() string {
	return r.b.Title
}

func (r *bookResolve) Body() string {
	return r.b.Body
}

func (r *bookResolve) Author() *authorResolve {
	return resolveBookAuthor(r.b.ID)
}

func resolveBookAuthor(id graphql.ID) *authorResolve {
	a := db.FindAuthorBook(id)

	author := Author{ID: a.ID,
		Name: a.Name,
	}
	return &authorResolve{&author}
}

func (r *Resolver) Book(args struct{ ID graphql.ID }) *bookResolve {
	b := db.FindBook(args.ID)

	book := Book{ID: b.ID,
		Title: b.Title,
		Body:  b.Body}

	return &bookResolve{&book}
}

type authorResolve struct {
	a *Author
}

func (r *authorResolve) ID() graphql.ID {
	return r.a.ID
}

func (r *authorResolve) Name() string {
	return r.a.Name
}

func (r *Resolver) Author(args struct{ ID graphql.ID }) *authorResolve {
	var u *db.Author
	u = db.FindAuthor(args.ID)
	if u == nil {
		log.Println("Author does not exist!")
	}

	author := Author{ID: u.ID,
		Name: u.Name,
	}
	return &authorResolve{&author}
}

func (r *authorResolve) Books() *[]*bookResolve {
	return resolveAuthorBooks(r.a.ID)
}

func resolveAuthorBooks(id graphql.ID) *[]*bookResolve {
	bks := db.FindBooksOfAuthor(id)
	var bookresolvers []*bookResolve
	for _, bk := range bks {
		var book Book
		book.ID = bk.ID
		book.Title = bk.Title
		book.Body = bk.Body
		bookresolvers = append(bookresolvers, &bookResolve{&book})
	}
	return &bookresolvers
}

func (r *Resolver) CreateAuthor(args struct {
	ID   graphql.ID
	Name string
}) *authorResolve {
	err := db.InsertAuthor(&db.Author{ID: graphql.ID(args.ID), Name: args.Name})

	if err != nil {
		log.Printf("Can't insert author:", err)
		return &authorResolve{&Author{}}
	}
	return &authorResolve{&Author{ID: graphql.ID(args.ID),
		Name: args.Name}}
}
