package main

import (

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"	
)

type Book struct {
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	Rating int `json:"rating"`
}

func GetBooks(c *fiber.Ctx){

	c.Send("All books")
}

func GetBook(c *fiber.Ctx){

	c.Send("One books")
}

func NewBook(c *fiber.Ctx){

	c.Send("Add a new books")
}

func DeleteBook(c *fiber.Ctx){

	c.Send("Delete a new books")

}