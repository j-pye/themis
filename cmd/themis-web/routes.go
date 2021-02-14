package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func index(c echo.Context) error {
	content := c.Render(http.StatusOK, "index", map[string]interface{}{
		"data": "Placeholder",
	})
	return content
}
