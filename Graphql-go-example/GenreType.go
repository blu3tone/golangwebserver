package main

import "github.com/graphql-go/graphql"

var GenreType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Genre",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if genre, ok := p.Source.(*Genre); ok == true{
					return genre.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if genre, ok := p.Source.(*Genre); ok == true {
					return genre.Name, nil
				}
				return nil, nil
			},
		},
		"rate": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if genre, ok := p.Source.(*Genre); ok == true {
					return genre.Rate, nil
				}
				return nil, nil
			},
		},
	},
})

