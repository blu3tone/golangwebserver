package main

import (
	"log"
	"github.com/go-pg/pg"
	"github.com/matryer/way"
	"net/http"
	"fmt"
	"os"
	"GoHTTP/cmd/db"
	"GoHTTP/internal"
	"strconv"
	"encoding/json"
	"time"
	"github.com/verifier"
	)

type Server struct {
	db     *pg.DB
	router *way.Router
	email  string
}

func (s *Server) handLogin(user *model.User) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		verify := verifier.New()

		username := way.Param(r.Context(), "username")
		password := way.Param(r.Context(), "password")

		user = getUser(1, s)

		log.Printf(strconv.Itoa(user.ID))

		if user == nil{
			fmt.Fprintln(w, "Error")
			os.Exit(100)
			verify.That(user == nil, "Your username or password were wrong!")

		}

		if username == user.Username && password == user.Password && user.RoleID ==1{
			log.Printf("Match!")
			cookie := &http.Cookie{}

			cookie.Name = "sessionID"
			cookie.Value = "something"
			cookie.Expires = time.Now().Add(48 *time.Second)

			http.SetCookie(w, cookie)

			fmt.Fprintln(w, "You are on admin page!!")
		}else{
			fmt.Fprintf(w, "Your username or password were wrong!")
		}


	}
}

func getUser(id int, s *Server) *model.User{
	user := new(model.User)

	err := s.db.Model(user).Where("id = ?", id).Select()
	if err != nil{
		return nil
	}
	return user
}

func (s *Server) handleAPI() http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Wellcome to my API web page!!")
	}
}

func (s *Server) getFirstBook(user model.User)  http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {
		book := new(model.Book)
		fmt.Print(user.ID)
		err := s.db.Model(book).Join("LEFT JOIN user_books ").JoinOn(" user_books.book_id = book.id").Where("user_books.user_id = ?", 1).First()

		if err!=nil{
			fmt.Fprint(w, "No book was found: ", err)
		} else {
			json.NewEncoder(w).Encode(map[string]string{
				"Id":  strconv.Itoa(book.ID),
				"Title": book.Title,
				"AuthorId":  strconv.Itoa(book.AuthorID),
				"Editor": strconv.Itoa(book.Editor),
			})
		}
	}
}

func main()  {
	server := Server{}

	server.db = db.Connect()

	server.router = way.NewRouter()

	var user model.User

	server.router.HandleFunc("GET", "/api/", server.handleAPI())

	server.router.HandleFunc("GET", "/login/:username/:password", server.handLogin(&user))

	server.router.HandleFunc("GET", "/book", server.getFirstBook(user))

	log.Fatal(http.ListenAndServe(":10000", server.router))
}







