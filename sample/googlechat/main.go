package main

import (
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"
)

func main() {
	var (
		sListenPort string = ":9999"
	)
	e := echo.New()
	e.GET("/test", SendSampleMessage)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Goodle Chat Message <h1>sample ::: Hello, World</h1>")
	})

	e.Logger.Fatal(e.Start(sListenPort))
}

// api handler functions...
func SendSampleMessage(c echo.Context) error {
	var (
		// sError     error  = nil
		sTimeStr   string = time.Now().Format("2006-01-02 15:04:05")
		sResultMsg string = sTimeStr + " ::: "
	)

	sResultMsg += "Nothing to do..."

	return c.String(http.StatusOK, sResultMsg)
}
