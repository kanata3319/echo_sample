package main

// go modコマンドで必要なライブラリを導入
import (
	"net/http"

	"github.com/labstack/echo"
	"rsc.io/quote"
)

// main関数が最初に呼び出される
func main() {
	e := echo.New()

	// http://localhost:1323/hello にアクセスすると表示
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n"+quote.Hello())
	})

	e.Logger.Fatal(e.Start(":1323"))
}
