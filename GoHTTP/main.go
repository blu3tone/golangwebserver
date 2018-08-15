package main

import (
	"log"
	"github.com/go-pg/pg"
	"github.com/matryer/way"
	"net/http"
	"fmt"
		"GoHTTP/cmd/db"
	"GoHTTP/internal"
	"strconv"
	"encoding/json"
	"time"
	"context"

	"golang.org/x/crypto/bcrypt"
	"os"
)

type Server struct {
	db     *pg.DB
	router *way.Router
	email  string
}


func (s *Server) handleAPI() http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Wellcome to my API web page!!")
	}
}

func (s *Server) handLogin() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {


		username := way.Param(r.Context	(), "username")
		password := way.Param(r.Context(), "password")

		user := getUser(username, s)

		if user == nil{
			fmt.Fprintf(w, "Your username or password were wrong!")
			panic("Your username or password were wrong!")
			os.Exit(100)
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err!=nil{
			fmt.Fprintf(w, "Your username or password were wrong!")

		}
		ctx := context.WithValue(r.Context(), "userID", 1)

		r = r.WithContext(ctx)

		log.Printf("Match!")

		expiration := time.Now().Add(365 * 24 * time.Hour)

		cookie := http.Cookie{Name: "userid", Value : strconv.Itoa(user.ID) , Expires: expiration}

		http.SetCookie(w, &cookie)

		fmt.Fprintln(w, "You are on admin page!!")
	}
}

func getUser(username string, s *Server) *model.User{
	user := new(model.User)

	err := s.db.Model(user).Where("username = ? ", username).Select()
	if err != nil{
		return nil
	}
	return user
}

func (s *Server) getFirstBook()  http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {
		book := new(model.Book)

		userid := r.Context().Value("userID")

		if userid ==nil{
			fmt.Fprintf(w, "Please login!!")
			return
		}

		err := s.db.Model(book).Join("LEFT JOIN user_books ").JoinOn(" user_books.book_id = book.id").Where("user_books.user_id = ?", userid).First()
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

/*func AddContextID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)
		cookie, _ := r.Cookie("userid")
		if cookie != nil {
		//Add data to context
			ctx := context.WithValue(r.Context(), "userID", 1)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			ctx := context.WithValue(r.Context(), "userID", 1)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}*/

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

func main()  {
	server := Server{}

	server.db = db.Connect()

	server.router = way.NewRouter()

	server.router.HandleFunc("GET", "/api/", server.handleAPI())

	server.router.Handle("GET", "/login/:username/:password", middlewareOne(server.handLogin()))

	server.router.HandleFunc("		GET", "/book", server.getFirstBook())

	log.Fatal(http.ListenAndServe(":10000", server.router))
}

