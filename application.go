package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/polyamanita/polyamanita-server/src/lib"
	"github.com/polyamanita/polyamanita-server/src/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Polyamanita API
// @version		1.0
// @description	API for Polyamanita server functions
func main() {
	godotenv.Load()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "5000"
	}

	RunServer(port)
}

func RunServer(port string) {
	r := gin.New()

	l := lib.NewLogger(os.Stdout)

	c, err := routes.NewController(l)
	if err != nil {
		log.Fatalf("error creating client: %s", err.Error())
	}

	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/session", c.Login)                 // login
	r.DELETE("/session", c.Logout)              // logout
	r.POST("/auth", c.PostAuths)                // email account, get auth code for register
	r.POST("/authGen", c.PostAuthsGen)          // email account, get auth code for general
	r.POST("/auth/refresh", c.RefreshAuthToken) // refresh auth token
	r.POST("/users", c.RegisterUser)            // register user

	users := r.Group("/users")
	users.Use(c.AuthenticateUser()) // check user is logged in
	{
		users.GET("", c.SearchUser)              // search for a user
		users.GET("/captures", c.GetAllCaptures) // get captures from all users
		users.GET("/:UserID", c.GetUser)         // get user

		user := users.Group("/:UserID")
		userAuth := users.Group("/:UserID")
		{
			userAuth.Use(c.AuthorizeUser())                // check logged in ID == UserID
			userAuth.PUT("", c.UpdateUser)                 // update user info
			userAuth.DELETE("", c.DeleteUser)              // delete user
			userAuth.POST("/images", c.UploadCaptureImage) // upload any image (usually captures)

			user.GET("/followers", c.GetUserFollowers) // get user followers
			user.GET("/feed", c.GetUserFeed)           // get follows feed

			captures := user.Group("/captures")

			capturesAuth := userAuth.Group("/captures")
			{
				captures.GET("", c.GetCapturesList)  // get list of captures
				capturesAuth.POST("", c.AddCaptures) // add captures

				capture := captures.Group("/:CaptureID")
				{
					capture.GET("", c.GetCapture) // get capture details
				}
			}

			followsAuth := userAuth.Group("/follows")
			follows := user.Group("/follows")
			{
				follows.GET("", c.GetUserFollows)                // get user's follows
				followsAuth.POST("", c.AddFollow)                // add a follow
				followsAuth.DELETE("/:FollowID", c.DeleteFollow) // unfollow a user
			}
		}
	}

	r.GET("/", c.TestHello)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + port)
}
