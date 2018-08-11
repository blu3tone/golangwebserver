package main

import (
	"github.com/go-pg/pg"
	"strings"
	"os"
	"github.com/go-pg/pg/orm"
	"log"
	"GoHTTP/internal"
)

func main() {
	dbInsert := `INSERT INTO authors VALUES (1, now(), now(), NULL, 'Nguyen Nhat Anh', '0932414134', 'District 9');
	INSERT INTO authors VALUES (2, now(), now(), NULL, 'Mai Phuong', '0425254221142', 'District 7');
	
	INSERT INTO genres VALUES(1, now(), now(), NULL, 'Tieu thuyet');
	INSERT INTO genres VALUES(2, now(), now(), NULL, 'Lich su');
	INSERT INTO genres VALUES(3, now(), now(), NULL, 'Khoa hoc');

	INSERT INTO books VALUES(1, now(), now(), NULL, 'Dao mong mo', 1, 1);
	INSERT INTO books VALUES (2, now(), now(), NULL, 'Chuc mot ngay tot lanh', 1, 1);
	INSERT INTO books VALUES(3, now(), now(), NULL, 'Hoa vang tren co xanh', 1, 1);
	INSERT INTO books VALUES(4, now(), now(), NULL, 'Nam cham', 2, 2);
	INSERT INTO books VALUES(5, now(), now(), NULL, 'Hon noi', 2, 1);

	INSERT INTO book_genres VALUES(1,1);
	INSERT INTO book_genres VALUES(2,1);
	INSERT INTO book_genres VALUES(3,1);
	INSERT INTO book_genres VALUES(4,2);
	INSERT INTO book_genres VALUES(5,3);

	INSERT INTO roles VALUES (1, 1, 'SUPER_ADMIN');
	INSERT INTO roles VALUES (2, 2, 'ADMIN');
	INSERT INTO roles VALUES (3, 3, 'COMPANY_ADMIN');
	INSERT INTO roles VALUES (4, 4, 'LOCATION_ADMIN');
	INSERT INTO roles VALUES (5, 5, 'USER');

	INSERT INTO users VALUES (1, now(),now(), NULL, 'Admin', 'Admin', 'admin', 'admin', 'uyen@gmail.com', NULL, NULL, NULL, NULL, true, 1, 1 );
	
	INSERT INTO user_books VALUES(1,1);
	INSERT INTO user_books VALUES(2,1);`
	queries := strings.Split(dbInsert, ";")


	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "cvbmnb",
		Database: "library",
	})


	_, err := db.Exec("SELECT 1")
	if err!= nil{
		panic("Database is not exist!")
		os.Exit(100)
	}
	createSchema(db, &model.Author{}, &model.Book{}, &model.Role{}, &model.User{}, &model.Genre{}, &model.BookGenre{}, &model.UserBook{})

	for _, v := range queries[0 : len(queries)-1] {
		_, err := db.Exec(v)
		if err!= nil{
			panic("Can not insert data: " + err.Error())
		}
	}
	/*userInsert := `INSERT INTO users VALUES (1, now(),now(), NULL, 'Admin', 'Admin', 'admin', '%s', 'uyen@gmail.com', NULL, NULL, NULL, NULL, true, 1, 1 );`
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	_, err = db.Exec(fmt.Sprintf(userInsert, hashedPW))

	if err != nil{
		panic("Can not create User " + err.Error())
	}
	dbInsert = `INSERT INTO user_books VALUES(1,1);
	INSERT INTO user_books VALUES(2,1);`
	queries = strings.Split(dbInsert, ";")

	for _, v := range queries[0 : len(queries)-1] {
		_, err := db.Exec(v)
		if err!= nil{
			panic("Can not insert data: " + v)
		}
	}*/

	log.Printf("Done!")


}

func createSchema(db *pg.DB, models ...interface{}) {
	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
		})
		if err != nil{
			panic("Can not create table: " +err.Error())
		}
	}
}