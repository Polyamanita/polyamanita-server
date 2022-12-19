package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/polyamanita/polyamanita-server/src/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Polyamanita API
//	@version		1.0
//	@description	API for Polyamanita server functions
//	@host			<some url>
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

	c, err := routes.NewController()
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

	r.POST("/session")      // login
	r.DELETE("/session")    // logout
	r.POST("/auth")         // email account, get auth code
	r.POST("/auth/refresh") // refresh auth token
	r.POST("/users")        // register user

	users := r.Group("/users")
	users.Use(c.AuthenticateUser()) // check user is logged in
	{
		users.GET("")         // search for a user
		users.GET("/:UserID") // get user

		user := users.Group("/:UserID")
		userAuth := users.Group("/:UserID")
		userAuth.Use(c.AuthorizeUser()) // check logged in ID == UserID
		{
			user.GET("/followers") // get user followers
			user.GET("/feed")      // get follows feed
			userAuth.PUT("")       // update user info
			userAuth.DELETE("")    // delete user

			captures := user.Group("/captures")
			capturesAuth := userAuth.Group("/captures")
			{
				captures.GET("")        // get list of captures
				capturesAuth.POST("")   // add captures
				capturesAuth.DELETE("") // delete list of captures

				capture := captures.Group("/:MushroomCaptureID")
				{
					capture.GET("")                   // get capture details
					capture.GET("/image/:InstanceID") // download image
				}
			}

			followsAuth := userAuth.Group("/follows")
			follows := user.Group("/follows")
			{
				follows.GET("")                  // get user's follows
				followsAuth.POST("")             // add a follow
				followsAuth.DELETE("/:FollowID") // unfollow a user
			}

		}

	}

	r.GET("/", c.TestHello)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + port)
}
