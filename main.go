package main

import (
	"os"
	"strings"

	"github.com/leanrgo/scrapper"

	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html") //.String(http.StatusOK, "Hello, world!")
}
func handleScape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))

	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
	//return c.File("home.html") //.String(http.StatusOK, "Hello, world!")
}
func main() {
	//https://echo.labstack.com/guide/installation
	//echo 서버 사용
	//설치
	//go get -u github.com/labstack/echo/...
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScape)
	e.Logger.Fatal(e.Start(":1323"))
}
