package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Hardw01f/SimpleScoreServer/pkg/conn"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}","latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
		Output: os.Stdout,
	}))

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health Active\n")
	})
	e.GET("/scoreinit", conn.ScoreInit())
	e.POST("/addscore", conn.AddScore())
	e.Logger.Fatal(e.Start(":9090"))
}
