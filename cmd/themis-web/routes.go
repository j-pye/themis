package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func home(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "Jonathan",
	})
}
