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

	// initialization.
	userController := controllers.NewUserController(NewMySQLDb())
	postDB, postStorage := controllers.NewPostController(NewMySQLDb(), NewStorageHandler())
	followController := controllers.NewFollowController(NewMySQLDb())

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
	r.POST("/update-user/:user_id", func(r echo.Context) error { return userController.UpdateUser(r) })
	r.DELETE("/delete-users/:id", func(r echo.Context) error { return userController.DeleteUser(r) })
	r.GET("/my", func(r echo.Context) error {
		following, follower, _ := followController.FfNumImpl(r, "")
		r.Set("following", following)
		r.Set("follower", follower)
		short, long, _ := postDB.GetPostsImpl(r, "")
		r.Set("short", &short)
		r.Set("long", &long)
		return userController.GetInfo(r, "")
	})
	r.GET("/get-user/:user_id", func(r echo.Context) error {
		userID := r.Param("user_id")
		following, follower, _ := followController.FfNumImpl(r, userID)
		r.Set("following", following)
		r.Set("follower", follower)
		short, long, _ := postDB.GetPostsImpl(r, userID)
		r.Set("short", &short)
		r.Set("long", &long)
		return userController.GetInfo(r, userID)
	})

	// post
	r.GET("/get-federation", func(r echo.Context) error { return postStorage.GetFederation(r) })
	r.GET("/get-presign-url", func(r echo.Context) error {
		num, _ := postDB.GetUserPostNumImpl(r)
		r.Set("num", num)
		return postStorage.GetPresignedURL(r)
	})
	r.GET("/postup-num/:user_id", func(r echo.Context) error { return postDB.GetPostNum(r) })
	r.GET("/get-post/:user_id", func(r echo.Context) error { return postDB.GetUserPost(r) })
	r.POST("/postup", func(r echo.Context) error { return postDB.CreatePost(r) })
	r.POST("/watch", func(r echo.Context) error { return postDB.WatchPost(r) })
	r.POST("/good", func(r echo.Context) error { return postDB.GoodPost(r) })

	// follow
	r.GET("/followed", func(r echo.Context) error {
		follows, _ := followController.FollowedGetImpl(r)
		r.Set("follows", follows)
		return userController.ResFollows(r)
	})
	r.GET("/following", func(r echo.Context) error {
		follows, _ := followController.FollowingGetImpl(r)
		r.Set("follows", follows)
		return userController.ResFollows(r)
	})
	r.POST("/follow", func(r echo.Context) error { return followController.Follow(r) })
	r.DELETE("/unfollow", func(r echo.Context) error { return followController.UnFollow(r) })
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
