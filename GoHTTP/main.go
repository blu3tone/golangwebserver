package main

import (
	"log"
	"github.com/go-pg/pg"
	"net/http"
	"fmt"
	"GoHTTP/cmd/db"
	"GoHTTP/internal"
	"strconv"
	"encoding/json"
	"time"
	"context"
	"github.com/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"github.com/gorilla/mux"
	"io/ioutil"
	"strings"
)

type Server struct {
	db     *pg.DB
	router *mux.Router
}
type Login struct {
	Username	string `json:"username"'`
	Password	string `json:"password"`
}
type JwtClaims struct{
	Name	string	`json:"name"`
	jwt.StandardClaims
}

type Exception struct{
	Message string `json:"message"`
}
const(
	tokenName = "AccessToken"
)

func (s *Server) handleAPI() http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Wellcome to my API web page!!")
	}
}

func (s *Server) handLogin() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		var lg Login

		resBody, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(resBody, &lg)
		if err != nil{
			fmt.Fprintf(w, err.Error())
			return
		}

		password :=  lg.Password
		username := lg.Username

		user := getUser(username, s)

		log.Printf("Match!")

		if user == nil{
			fmt.Fprintf(w, "Your username or password were wrong!")
			panic("Your username or password were wrong!")
			os.Exit(100)
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err!=nil{
			fmt.Fprintf(w, "Your username or password were wrong!")

		}

		token, err := createJwtToken(username, strconv.Itoa(user.ID))
		if err!=nil{
			fmt.Fprintf(w, "Can't not create token: " + err.Error())
			return
		}

		expiration := time.Now().Add(365 * 24 * time.Hour)

		cookie := http.Cookie{Name: tokenName, Value : token , Expires: expiration}

		http.SetCookie(w, &cookie)

		fmt.Fprintln(w, "You are on admin page!!")
		json.NewEncoder(w).Encode(map[string]string{"Token": token})
	}
}

func createJwtToken(username, userid string) (string, error){
	claims := JwtClaims{
		username,
		jwt.StandardClaims{
			Id: userid,
			ExpiresAt: time.Now().Add(24*time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err :=rawToken.SignedString([]byte("mySecret"))
	if err != nil{
		return "", err
	}

	return token, err
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

func (s * Server) TestEndpoint(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims");
	if claims == nil{
		fmt.Fprintf(w, "Authentication failed")
	}
	fmt.Fprintf(w, "You are on top secret admin page!")
}

func AddContextID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)
		cookie, _ := r.Cookie("userid")
		if cookie != nil {
			ctx := context.WithValue(r.Context(), "userID", cookie.Value)

			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {

				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}
					return []byte("mySecret"), nil
				})

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					uid := claims["name"]
					log.Printf(uid.(string)+ "!! Done")
					ctx := context.WithValue(req.Context(), "claims", claims)
					next(w, req.WithContext(ctx))
				}

				if err!= nil{
					fmt.Fprintf(w, err.Error())
				}
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	})
}


func main()  {
	server := Server{}

	server.db = db.Connect()

	server.router = mux.NewRouter()

	server.router.HandleFunc(  "/api/", server.handleAPI()).Methods("GET")

	server.router.Handle( "/login", middlewareOne(server.handLogin())).Methods("POST")

	server.router.HandleFunc( "/secret", ValidateMiddleware(server.TestEndpoint)).Methods("GET")

	server.router.HandleFunc( "/book", server.getFirstBook()).Methods("GET")


	newcontext := AddContextID(server.router)


	log.Fatal(http.ListenAndServe(":10000", newcontext))
}

