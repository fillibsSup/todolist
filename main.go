package main

import (
	"fmt"
	"todolist/database"
	"todolist/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// database config
func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=123456 dbname=goapiDB port=54321"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated DB")
}

// routepath
func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Get("/todos/:id", models.GetTodoById)
	app.Post("/todos", models.CreateTodo)
	app.Put("/todos/:id", models.UpdateTodo)
	app.Delete("/todos/:id", models.DeleteTodo)
}

// main
func main() {
	app := fiber.New()
	app.Use(cors.New())
	initDatabase()
	setupRoutes(app)
	app.Listen(":8000")
}
