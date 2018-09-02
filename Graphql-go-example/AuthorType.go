package main

import "github.com/graphql-go/graphql"

var AuthorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if author, ok := p.Source.(*Author); ok == true{
					return author.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if author, ok := p.Source.(*Author); ok == true {
					return author.Name, nil
				}
				return nil, nil
			},
		},
		"birthday": &graphql.Field{
			Type: graphql.NewNonNull(graphql.DateTime),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if author, ok := p.Source.(*Author); ok == true {
					return author.Birthday, nil
				}
				return nil, nil
			},
		},
	},
})


