package main

import (
	"Tugas_4MIKTII/config"
	"Tugas_4MIKTII/handlers"
	"Tugas_4MIKTII/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
    // Initialize database
    db := config.InitDB()
    defer db.Close()

    // Create Echo instance
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routes
    e.POST("/login", handlers.Login(db))

    // Todo routes (protected, only for Editor)
    todoGroup := e.Group("/todos")
    todoGroup.Use(middlewares.JWTMiddleware())
    todoGroup.Use(middlewares.RoleMiddleware("Editor"))
    todoGroup.GET("", handlers.GetAllTodos(db))
    todoGroup.POST("", handlers.CreateTodo(db))
    todoGroup.PUT("/:id", handlers.UpdateTodo(db))
    todoGroup.DELETE("/:id", handlers.DeleteTodo(db))

    // User routes (protected, only for Admin)
    userGroup := e.Group("/users")
    userGroup.Use(middlewares.JWTMiddleware())
    userGroup.Use(middlewares.RoleMiddleware("Admin"))
    userGroup.GET("", handlers.GetAllUsers(db))
    userGroup.POST("", handlers.CreateUser(db))
    userGroup.PUT("/:id", handlers.UpdateUser(db))
    userGroup.DELETE("/:id", handlers.DeleteUser(db))

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}
