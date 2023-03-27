package handlers

import (
	"net/http"

	"Meow-fi/internal/auth"
	controllers "Meow-fi/internal/services"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func NoticeAuthHandler(noticeController *controllers.NoticeController, c echo.Context) {
	id := c.Param("id")
	notice := noticeController.GetNotice(id)
	if notice.Client.UserId == c.Get("user").(*jwt.Token).Claims.(*auth.JwtCustomClaims).Id {
		c.JSON(http.StatusOK, "You added notice: \""+notice.Containing+"\" at "+notice.CreatedAt.Format("02-Jan-2006"))
	} else {
		c.JSON(http.StatusOK, "Somebody added notice: \""+notice.Containing+"\"")
	}
}
func NoticeGuestHandler(noticeController *controllers.NoticeController, c echo.Context) {
	c.JSON(http.StatusOK, "Not allowed for guests")
}
