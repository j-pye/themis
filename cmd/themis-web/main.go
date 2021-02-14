package main

import (
	"html/template"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	appPort := os.Getenv("WEB_PORT")
	if appPort == "" {
		log.Println("web app port not defined")
	}
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	app := echo.New()
	// specify asset directories
	app.Static("/", path+"/cmd/themis-web/public/assets")
	// attach renderer for templates and load all templates
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob(path + "/cmd/themis-web/public/*/*.html")),
	}
	app.Renderer = renderer
	// set basic middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	// set all available routes
	defineRoutes(app)
	// start app on port
	app.Logger.Fatal(app.Start(":" + appPort))
}

func defineRoutes(app *echo.Echo) {
	app.GET("/", index).Name = "/"
	app.GET("/home", index).Name = "/"
	app.GET("/network", index).Name = "/network"
	app.GET("/policy", index).Name = "/policy"
}
