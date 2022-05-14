package main

import (
	"auth-server/router"

	"github.com/labstack/echo/v4"
)

func main() {
	var (
		e   *echo.Echo = echo.New()
		err error      = nil
	)

	err = router.Init(e) // api handler 함수 설정
	if err != nil {
		e.Logger.Error(err.Error())
	}

	e.Static("/", "static")
	// err = static.Init()
	// if err != nil {
	// 	e.Logger.Error(err.Error())
	// }

	e.Logger.Fatal(e.Start(":1323"))
}
