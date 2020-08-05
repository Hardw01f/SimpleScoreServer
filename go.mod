module guthub.com/Hardw01f/SimpleScoreServer

go 1.13

require (
	github.com/Hardw01f/SimpleScoreServer/pkg/conn v0.0.0-00010101000000-000000000000
	github.com/jinzhu/gorm v1.9.15
	github.com/labstack/echo/v4 v4.1.16
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
)

replace github.com/Hardw01f/SimpleScoreServer/pkg/conn => ./pkg/conn
