package main

import (

	"fmt"
	"github.com/gofiber/fiber"	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

		
)

var (

	DBConn *gorm.DB

)

type Book struct {
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	Rating int `json:"rating"`
}


func GetBooks(c *fiber.Ctx){
	db:= DBConn
	var books []Book
	db.Find(&books)

	c.JSON(books)
}

func GetBook(c *fiber.Ctx){
	id := c.Params("id")
	db: = DBConn
	var book Book
	db.Find(&books, id)

	c.JSON(book)
}

func NewBook(c *fiber.Ctx){
	db: = DBConn
	/*var book Book

	book.Title = "1984"
	book.Author = "GO"
	book.Rating = 5*/

	book: = new(Book)
	if err:= c.BodyParser(book); err !=nil {

		c.Status(503).Send(err)

		return
	}

	db.Create(&book)

	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx){
	id := c.Params("id")
	db: = DBConn
	var book Book

	db.First(&book, id)
	if book.Title == "" {

		c.Status(500).Send("No book")
		return

	}
	
	db.Delete(&book)
	c.Send("Deleted")
}

func helloWorld(c *fiber.Ctx){

	c.Send("Hello, World")
}

func setupRoutes(app *fiber.App){

	app.Get("/api/v1/book", GetBooks)
	app.Get("/api/v1/book/:id", GetBook)
	app.Post("/api/v1/book", NewBook)
	app.Delete("/api/v1/book/:id", DeleteBook)

}

func initDatabase(){

	var err error
	DBConn, err = gorm.Open("sqlite", "books.db")
	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("DB connection successfully openned")

	//Auto igration
	DBConn.AutoMigrate(&Book{})
	fmt.Println("Database Migrated")




}

func main (){

	app := fiber.New()

	initDatabase()
	defer DBConn.Close()

	setupRoutes(app)

	//app.Get("/", helloWorld)

	app.Listen(3000)
}