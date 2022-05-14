package cookie

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// 쿠키 구조체, 내장되어 있으므로 참조.
// type Cookie struct {
// 	Name       string    `json:"name"`
// 	Value      string    `json:"value"`
// 	Path       string    `json:"path"`
// 	Domain     string    `json:"domain"`
// 	Expires    time.Time `json:"expires"`
// 	RawExpires string    `json:"rawexpires"`
// 	MaxAge     int       `json:"maxage"`
// 	Secure     bool      `json:"secure"`
// 	HttpOnly   bool      `json:"httponly"`
// 	Raw        string    `json:"raw"`
// 	Unparsed   []string  `json:"unparsed"`
// }

func Login(c echo.Context) error {
	cookie := new(http.Cookie)
	sMap := map[string]string{}
	sErr := c.Bind(&sMap)
	if sErr != nil {
		return c.String(http.StatusOK, "아이고 에러가 발생했네 : "+sErr.Error())
	}

	cookie.Name = "cookie_login"
	cookie.Value = sMap["id"] + "__" + sMap["pw"]
	cookie.Expires = time.Now().Add(30 * time.Second)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie("+sMap["id"]+"__"+sMap["pw"]+")")
}
