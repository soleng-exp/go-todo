package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"./Database"
	"./Handlers"
)

func main() {
	// init db
	Database.InitDB("storage.db")

	// init Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "public/View/todo.html")

	e.Static("/nm", "node_modules")
	g := e.Group("static")
	g.Static("/js", "public/js")
	g.Static("/css", "node_modules")

	e.GET("/tasks", Handlers.GetTasks())
	e.PUT("/tasks", Handlers.PutTask())
	e.DELETE("/tasks/:id", Handlers.DeleteTask())

	// Start as a web server
	e.Logger.Fatal(e.Start(":1323"))
}
