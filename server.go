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

	e.Use(middleware.CORS())
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "public/View/todo.html")
	e.Static("/nm", "node_modules")

	g := e.Group("static")
	g.Static("/js", "public/js")
	g.Static("/css", "node_modules")

	// Login route
	e.POST("/login", Handlers.Login())

	e.GET("/tasks", Handlers.GetTasks())

	// Restricted group
	r := e.Group("/todo")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", Handlers.Restricted())
	r.PUT("", Handlers.PutTask())
	r.DELETE("/:id", Handlers.DeleteTask())

	// Start as a web server
	e.Logger.Fatal(e.Start(":1323"))
}
