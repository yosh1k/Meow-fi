package transport

import (
	"Meow-fi/internal/config"
	"Meow-fi/internal/database"
	controllers "Meow-fi/internal/services"

	"net/http"

	"github.com/labstack/echo"
)

func Init() {
	e := echo.New()

	noticeController := controllers.NewNoticeController(database.NewSqlHandler())
	userController := controllers.NewUserController(database.NewSqlHandler())
	e.GET("/notices", func(ctx echo.Context) error {
		tasks := noticeController.GetAllNotices()
		ctx.Bind(&tasks)
		return ctx.JSON(http.StatusOK, tasks)
	})
	e.GET("/notices/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		str := noticeController.GetNoticeInfo(id)
		ctx.Bind(&str)
		return ctx.JSON(http.StatusOK, str)
	})
	e.POST("/notices", func(ctx echo.Context) error {
		noticeController.Create(ctx)
		return ctx.String(http.StatusOK, "created")
	})

	e.GET("/users", func(ctx echo.Context) error {
		users := userController.GetUser()
		ctx.Bind(&users)
		return ctx.JSON(http.StatusOK, users)
	})
	e.POST("/users", func(ctx echo.Context) error {
		userController.Create(ctx)
		return ctx.String(http.StatusOK, "created")
	})
	e.Logger.Fatal(e.Start(config.ServerPort))
}
