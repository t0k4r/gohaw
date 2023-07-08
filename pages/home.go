package pages

import (
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	if isHx(c.Request()) {
		return c.Render(200, "Home", nil)
	}
	return c.Render(200, "pageHome.html", nil)
}
