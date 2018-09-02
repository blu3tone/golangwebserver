package main

import (
	"github.com/graphql-go/graphql"
)


var BookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if book, ok := p.Source.(*Book); ok == true{
					return book.ID, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if book, ok := p.Source.(*Book); ok == true{
					return book.ID, nil
				}
				return nil, nil
			},
		},
		"body": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if book, ok := p.Source.(*Book); ok == true{
					return book.ID, nil
				}
				return nil, nil
			},
		},
		"author_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if book, ok := p.Source.(*Book); ok == true{
					return book.ID, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	BookType.AddFieldConfig("author", &graphql.Field{
		Type: graphql.NewNonNull(AuthorType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if book, ok := p.Source.(*Book); ok == true {
				return GetAuthorIDBook(book.Author_ID)
			}
			return nil, nil
		},
	})
}