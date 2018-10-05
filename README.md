# golangwebserver

### Getting Started

This is the example project use Graphql + relay + react + postgres


### Setup

```
1. Download the project (git clone, go get...)
2. cd ${GOPATH}/src/github.com/howtographql/graphql-go
3. Replace connection string to Postgres (db.go) and environment (docker-compose.yml) 
4. docker-compose up -d
5. go run server.go

```

### Graphql example

Execute queries in GraphiQL by visiting

```
 http://localhost:4000/
```

#### Query:

User, book, author

```
{ user(id:"user1")
	{email,
	 password, 
	 books{
	 	title, 
	 	body, 
	 	author{name}
	 }
}
```
```
{author(id:"a1")
	{id, name, books{title, body}
}
```

#### Mutation

*** Create user, book, books of user, author
*** Update, delete user
```
mutation {
 createUser (id :"u1", email: "abc@mail.com", password: "123456") {
    email, password, books{id, title}
 }
}
```

```
mutation{
  createAuthor(id: "a1", name:"Khai ve"){id, name}
}
```
```
mutation{
  createBook(id:"b1", title:"Du se quen", body:"body3", author_id:"a1"){id, title}
}
```
```
mutation {
 createRelUB(id :"1", user_id:"u1", book_id:"b1"){
   id, email
 }
}
```

### Reference 
[How to graphql](https://github.com/howtographql/graphql-go)
