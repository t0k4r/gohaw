package pages

// import (
// 	"gohaw/db"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// )

// func Filtered(title string) func(echo.Context) error {
// 	return func(c echo.Context) error {
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			return c.NoContent(400)
// 		}
// 		anime, err := db.AnimesFromInfoId(id)
// 		if err != nil {
// 			return err
// 		}
// 		if isHx(c.Request()) {
// 			return c.Render(200, "Home", anime)
// 		}
// 		return c.Render(200, "pageFiltered.html", anime)
// 	}
// }
