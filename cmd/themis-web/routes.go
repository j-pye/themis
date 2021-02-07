package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func index(c echo.Context) error {
	content := c.Render(http.StatusOK, "wrapper.html", map[string]interface{}{
		"appContent": "index.html",
		"data":       "Placeholder",
	})
	return content
}
