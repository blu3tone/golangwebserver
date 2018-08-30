package main

import (
	"github.com/graphql-go/graphql"
	"database/sql"
	"strconv"
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
)




var db *sql.DB


var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "User ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetUserByID(id)
			},
		},
		"book": &graphql.Field{
			Type: BookType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Book ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetBookByID(id)
			},
		},
	},
})


func handler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, string(query))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: string(query),
		})


		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}




func main()  {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: QueryType,
		Mutation: MutationType,
	})
	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", "user=postgres password=cvbmnb dbname=bookstore sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/graphql", handler(schema))
	log.Fatal(http.ListenAndServe("0.0.0.0:8086", nil))

}


