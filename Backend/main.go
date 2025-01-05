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

	app.Post("/api/todos", func(c *fiber.Ctx) error{
		var todos []Todo

		if err := db.Find(&todos).Error; err != nil{
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch todos")
		}

		return c.JSON(todos)
	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error{
		id, err:= c.ParamsInt("id")

		if err != nil{
			return c.Status(401).SendString("Invalid ID")
		}

		for i, t := range todos{
			if t.ID == id{
				todos[i].Done = !todos[i].Done
				break
			}
		}
		return c.JSON(todos)
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error{
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))
}