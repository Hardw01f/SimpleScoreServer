package conn

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

type ScoreParameters struct {
	Server string
	Score  int
}

type Score struct {
	gorm.Model
	Server string
	Score  int
}

func ScoreInit() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := gorm.Open("mysql", "root:toor@(mysql)/nanohardening?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			err = xerrors.New("Failed database connection")
			return c.String(http.StatusInternalServerError, err.Error())
		}
		defer db.Close()

		db.Set("gorm:table_options", "ENGINE=InnoDB")
		db.AutoMigrate(&Score{})
		return c.String(http.StatusOK, "Migrate Successful\n")
	}
}

func AddScore() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := gorm.Open("mysql", "root:toor@(mysql)/nanohardening?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			err = xerrors.New("Failed database connection")
			return c.String(http.StatusInternalServerError, err.Error())
		}
		defer db.Close()

		s := new(ScoreParameters)
		if err = c.Bind(s); err != nil {
			err = xerrors.New("Could not parse parameters")
			return c.String(http.StatusInternalServerError, err.Error())
		}

		var score Score
		var count int
		db.Where("server = ?", s.Server).Find(&score).Count(&count)

		fmt.Println(count)

		score.Server = s.Server
		score.Score = s.Score

		if count == 0 {
			db.Create(&score)
		} else {
			var originScore Score
			db.First(&originScore)
			db.Model(&score).Where("server = ?", score.Server).Update("score", originScore.Score+score.Score)
		}

		var returnScore Score
		db.Where("server = ?", s.Server).Find(&returnScore)

		return c.JSON(http.StatusOK, returnScore)

	}
}
