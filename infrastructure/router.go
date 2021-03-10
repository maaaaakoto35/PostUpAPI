package infrastructure

import (
	"io/ioutil"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/maaaaakoto35/PostUpAPI/interfaces/controllers"
)

// Init this func is initializing server.
func Init() {
	e := echo.New()

	userController := controllers.NewUserController(NewMySQLDb())
	postDB, postStorage := controllers.NewPostController(NewMySQLDb(), NewStorageHandler())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 認証なし
	e.POST("/setup", func(c echo.Context) error { return userController.CreateUser(c) })
	e.POST("/login", func(c echo.Context) error { return userController.LogIn(c) })

	// 認証あり
	r := e.Group("/api")
	config := setJwtConfig()
	r.Use(middleware.JWTWithConfig(config))

	// user
	r.GET("/get-users", func(r echo.Context) error { return userController.GetUsers(r) })
	r.GET("/get-user/:user_id", func(r echo.Context) error { return userController.GetUser(r) })
	r.POST("/update-user/:user_id", func(r echo.Context) error { return userController.UpdateUser(r) })
	r.DELETE("/delete-users/:id", func(r echo.Context) error { return userController.DeleteUser(r) })

	// post
	r.GET("/get-federation", func(r echo.Context) error { return postStorage.GetFederation(r) })
	r.GET("/get-presign-url/:num", func(r echo.Context) error { return postStorage.GetPresignedURL(r) })
	r.GET("/postup-num/:user_id", func(r echo.Context) error { return postDB.GetPostNum(r) })
	r.POST("/postup", func(r echo.Context) error { return postDB.CreatePost(r) })
	r.POST("/watch", func(r echo.Context) error { return postDB.WatchPost(r) })
	r.POST("/good", func(r echo.Context) error { return postDB.GoodPost(r) })

	e.Logger.Fatal(e.Start(":8080"))
}

func setJwtConfig() middleware.JWTConfig {
	// 公開鍵読み込み
	pubPath := os.Getenv("PUBLIC_KEY_PATH")
	pubKeyData, err := ioutil.ReadFile(pubPath)
	if err != nil {
		panic(err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyData)

	return middleware.JWTConfig{
		Claims:        &controllers.JwtCustomClaims{},
		SigningKey:    pubKey,
		ContextKey:    "jwt",
		SigningMethod: "RS256",
	}
}
