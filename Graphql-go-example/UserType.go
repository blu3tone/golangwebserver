package main

import (
	"github.com/graphql-go/graphql"
	"strconv"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*User); ok == true{
					return user.ID, nil
				}
				return nil, nil
			},
		},
		"username": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*User); ok == true {
					return user.Username, nil
				}
				return nil, nil
			},
		},
		"password": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*User); ok == true {
					return user.Password, nil
				}
				return nil, nil
			},
		},
	},
})

func init(){
	UserType.AddFieldConfig("book", &graphql.Field{
		Type: BookType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Book ID",
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetBookByIDAndUser(id, user.ID)
			}
			return nil, nil
		},
	})

	UserType.AddFieldConfig("books", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(BookType))),

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				return GetBookByIDAndUse(user.ID)
			}
			return []Book{}, nil
		},
	})


}