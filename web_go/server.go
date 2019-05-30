package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"fmt"
)

func hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World! (Web!!)")
}

func sum(ctx echo.Context) error{
	a := ctx.Param("a")
	b := ctx.Param("b")
	aNum, err := strconv.Atoi(a);
	if err != nil {
		fmt.Println(err)
	}
	bNum, err := strconv.Atoi(b);
	if err != nil {
		fmt.Println(err)
	}
	sum := aNum + bNum
	return ctx.String(http.StatusOK, strconv.Itoa(sum))
}

func main() {
	app := echo.New()
	app.GET("/", hello)
	app.GET("/sum/:a/:b", sum)
	app.Logger.Fatal(app.Start(":1323"))
}