package main

import (
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//"fmt"
	"strconv"
)

var db *gorm.DB

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("postgres", "user=postgres password=cvbmnb dbname=test1 sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&Person{}, &Address{})
}

type Person struct {
	gorm.Model
	Name    string
	Age     string
	Address []Address `gorm:"foreignkey:Address"`
}

type Address struct {
	PersonID int
	Phone string
	City  string
}

var people []Person
var addresses []Address

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1/person")
	{
		v1.POST("/", createPerson)
		v1.GET("/", fetchAllPerson)
		v1.GET("/:id", fetchSinglePerson)
		v1.DELETE("/:id", deletePerson)
	}

	v2 := router.Group("/api/v2/address")
	{
		v2.POST("/", createAddress)
		v2.GET("/", fetchAllAddress)
		v2.GET("/:id", fetchSingleAddress)
		v2.DELETE("/:id", deleteAddress)
	}
	router.Run(":8089")

}

func createPerson(c *gin.Context)  {
	var person Person

	person = Person{Name: c.PostForm("name"), Age: c.PostForm( "age")}
	db.Save(&person)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Person created successfully!", "resourceId": person.Name})
}

func createAddress(c *gin.Context) {
	var address Address
	s, err:= strconv.Atoi(c.PostForm("personID"))
	if (err != nil) {
		panic("error!")
	}
	address = Address{Phone: c.PostForm("Phone"), City: c.PostForm("City"), PersonID: s }
	db.Save(&address)
	c.JSON(http.StatusCreated, gin.H{"data": address})

}

func fetchAllPerson(c *gin.Context){
	db.Debug().Find(&people)

	c.JSON(http.StatusOK, gin.H{"data": people})

}

func fetchAllAddress(c *gin.Context){
	db.Debug().Find(&addresses)

	c.JSON(http.StatusOK, gin.H{"data": addresses})

}


func fetchSinglePerson (c * gin.Context){
	var person Person
	index := c.Param("id")
	db.Debug().Where("ID = ?", index).Find(&person)

	db.Debug().Where("person_id = ?", person.ID).Find(&addresses)

	person.Address = addresses

	c.JSON(http.StatusOK, gin.H{"data": person})
}


func fetchSingleAddress (c * gin.Context){
	var address Address

	db.Debug().Where("person_id = ?", c.Param("id")).Find(&address)

	c.JSON(http.StatusOK, gin.H{"data": address})
}

func deletePerson(c * gin.Context) {
	var person Person

	//db.Debug().First(&person, c.Param("id"))
	db.Debug().Where("id = ?", c.Param("id")).Delete(&person)
	if person.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No person found!"})
		return
	}
}

func deleteAddress(c * gin.Context){
	var address	 Address

	//db.Debug().Where("person_id = ?", c.Param("id")).First(&address)
	db.Debug().Where("person_id = ?", c.Param("id")).Delete(&address)
	if address.PersonID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No address found!"})
		return
	}


}

