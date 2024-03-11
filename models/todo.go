package models

import (
	"todolist/database"
	"github.com/gofiber/fiber/v2"

)

type Todo struct {
	ID 		  uint 	  `gorm:"primarykey" json:"id"`
	Title 	  string  `json:"title"`
	Completed bool 	  `json:"completed"`
}

//get todo
func GetTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []Todo
	db.Find(&todos)
	return c.JSON(&todos)
}

//get todo by id
func GetTodoById(c *fiber.Ctx) error{
	id := c.Params("id")
	db := database.DBConn
	var todo Todo
	err := db.Find(&todo, id).Error 
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "massage": "Could not find todo", "data": err})
	}

	return c.JSON(&todo)
}

//create todo
func CreateTodo(c *fiber.Ctx) error {
	db := database.DBConn
	todo := new(Todo)
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "massage": "Check your input", "data": err})
	}
	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "massage": "Could not create todo", "data": err})
	}
	return c.JSON(&todo)
}

//update todo by id
func UpdateTodo(c *fiber.Ctx) error {
	type UpdateTodo struct {
		Title 	  string  `json:"title"`
		Completed bool 	  `json:"completed"`
	}
	id := c.Params("id")
	db := database.DBConn
	var todo Todo
	err := db.Find(&todo, id).Error 
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "massage": "Could not find todo", "data": err})
	}

	var updatedTodo UpdateTodo
	err = c.BodyParser(&updatedTodo)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "massage": "Review your input", "data": err})
	}

	todo.Title = updatedTodo.Title
	todo.Completed = updatedTodo.Completed
	db.Save(&todo)
	return c.JSON(&todo)
}

//delete todo by id
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var todo Todo
	err := db.Find(&todo, id).Error 
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "massage": "Could not find todo", "data": err})
	}

	db.Delete(&todo)

	return c.SendStatus(200)
}