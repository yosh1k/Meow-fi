package transport

import (
	"Meow-fi/internal/auth"
	"Meow-fi/internal/config"

	"Meow-fi/internal/database"
	controllers "Meow-fi/internal/services"
	"time"

	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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
		users := userController.GetAllUsers()
		ctx.Bind(&users)
		return ctx.JSON(http.StatusOK, users)
	})

	e.POST("/registrate", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		err := userController.Create(username, password)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusOK, "created")
	})

	type jwtCustomClaims struct {
		Name string `json:"name"`
		jwt.RegisteredClaims
	}
	// configjwt := echojwt.Config{
	// 	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	// 		return new(jwtCustomClaims)
	// 	},
	// 	SigningKey: []byte("secret"),
	// }
	e.POST("login", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Throws unauthorized error
		user, err := userController.GetUserByLogin(username)
		if err != nil {
			return echo.ErrUnauthorized
		}
		if user.Login != username || user.Password != auth.HashPass(password, user.Salt, config.LocalSalt) {
			return echo.ErrUnauthorized
		}
		// Set custom claims
		claims := &jwtCustomClaims{
			username,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(config.SecretKeyJwt))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	})

	e.Logger.Fatal(e.Start(config.ServerPort))
}
