
package main

import "github.com/graphql-go/graphql"


var UserBookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserBook",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if userbook, ok := p.Source.(*UserBook); ok == true{
					return userbook.ID, nil
				}
				return nil, nil
			},
		},
		"book_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if userbook, ok := p.Source.(*UserBook); ok == true{
					return userbook.Book_ID, nil
				}
				return nil, nil
			},
		},
		"user_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if userbook, ok := p.Source.(*UserBook); ok == true{
					return userbook.User_ID, nil
				}
				return nil, nil
				},
		},

	},
})
