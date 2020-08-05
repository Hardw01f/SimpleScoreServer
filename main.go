package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Hardw01f/SimpleScoreServer/pkg/conn"
)

func main() {
	e := echo.New()

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/scoreinit", conn.ScoreInit())
	e.POST("/addscore", conn.AddScore())
	e.Logger.Fatal(e.Start(":9090"))
}
