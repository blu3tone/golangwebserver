package main

import (
	"github.com/graphql-go/graphql"
		"time"
	"strconv"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Description: "User's name",
					Type:  graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Description: "User's password",
					Type:  graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				username := p.Args["username"].(string)
				password := p.Args["password"].(string)

				user := &User{
					Username: username,
					Password: password,
				}
				err := InsertUser(user)
				return user, err
			},
		},
		"createAuthor": &graphql.Field{
			Type: AuthorType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Description: "Author's name",
					Type:  graphql.NewNonNull(graphql.String),
				},
				"birthday": &graphql.ArgumentConfig{
					Description: "Author's birthday",
					Type:  graphql.NewNonNull(graphql.DateTime),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				birthday := p.Args["birthday"].(string)
				t, err := time.Parse("2006-01-02", birthday)
				if err != nil{
					t, _ = time.Parse("2006-01-02", "0001-01-01")
				}

				author := &Author{
					Name: name,
					Birthday: t,
				}
				err = InsertAuthor(author)
				return author, err
			},
		},

		"createGenre": &graphql.Field{
			Type: GenreType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Description: "genre's name",
					Type:  graphql.NewNonNull(graphql.String),
				},
				"rate": &graphql.ArgumentConfig{
					Description: "rate",
					Type:  graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				rate := p.Args["rate"].(int)
				//rateint, _ := strconv.Atoi(rate)
				genre := &Genre{
					Name: name,
					Rate: rate,
				}
				err := InsertGenre(genre)
				return genre, err
			},
		},


		"createBook": &graphql.Field{
			Type: BookType,
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Description: "book's title",
					Type:  graphql.NewNonNull(graphql.String),
				},
				"body": &graphql.ArgumentConfig{
					Description: "book's body",
					Type:  graphql.NewNonNull(graphql.String),
				},
				"author_id": &graphql.ArgumentConfig{
					Description: "authorid",
					Type:  graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				title := p.Args["title"].(string)
				body := p.Args["body"].(string)
				author_id := p.Args["author_id"].(int)
				book := &Book{
					Title: title,
					Body:body,
					Author_ID: author_id,
				}
				err := InsertBook(book)
				return book, err
			},
		},

		"createUserBook": &graphql.Field{
			Type: UserBookType,
			Args: graphql.FieldConfigArgument{

				"book_id": &graphql.ArgumentConfig{
					Description: "bookid",
					Type:  graphql.NewNonNull(graphql.Int),
				},
				"user_id": &graphql.ArgumentConfig{
					Description: "userid",
					Type:  graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				bookid := p.Args["book_id"].(int)
				userid := p.Args["user_id"].(int)
				userbook := &UserBook{
					Book_ID: bookid,
					User_ID: userid,
				}
				err := InsertUserBook(userbook)
				return userbook, err
			},
		},

		"removeUser": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "User ID",
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					i := p.Args["id"].(string)
					id, err := strconv.Atoi(i)
					if err != nil {
						return nil, err
					}
					err =  RemoveUserByID(id)
					return nil, nil
			},
		},

		"removeBook": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Book ID",
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err =  RemoveBookByID(id)
				return nil, nil
			},
		},

		"removeAuthor": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Author ID",
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err =  RemoveAuthorByID(id)
				return nil, nil
			},
		},
	},
})
