package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct{
	gorm.Model
	Title string `json:"title"`
	Done bool `json:"done"`
	Body string `json:"body"`
}

func main(){
	fmt.Println("Hello, World!")

	db, err:= gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil{
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Database connected")

	err = db.AutoMigrate(&Todo{})
	if err != nil{
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Database migrated")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))


	app.Get("/healtcheck", func(c *fiber.Ctx) error{
		return c.SendString("OK")
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error{
		var todos []Todo
		if err := db.Find(&todos).Error; err != nil{
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}	
		
		return c.JSON(todos)
	})
		

	app.Post("/api/todos", func(c *fiber.Ctx) error{
		todo := new(Todo)
		if err := c.BodyParser(todo); err != nil{
			return fiber.NewError(fiber.StatusBadRequest,err.Error())
		}

		if todo.Body == ""{
			return fiber.NewError(fiber.StatusBadRequest,"Body is required")
		}

		if err := db.Create(&todo).Error; err != nil{
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(todo)
	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error{
		
	})

	log.Fatal(app.Listen(":4000"))
}