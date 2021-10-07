package routers

import (
	apiControllerV! "github.com/MrWildanMD/Go-Skeleton/controllers/api/v1"
	apiControllerV! "github.com/MrWildanMD/Go-Skeleton/controllers/api/v2"
	"github.com/MrWildanMD/Go-Skeleton/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {

	router := gin.Default()

	//Access to storage folders
	router.Static("/storage","storage")

	//Access to template folders
	router.Static("/templates","templates")
	router.LoadHTMLGlob("templates/*")

	router.Use(func(c *gin.Context) {
		//add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	//Route for api V1
	v1 := router.Group("/api/v1")

	//If you want to pass your route through spesific middlewares
	v1.Use(middlewares.UserMiddlewares())
	{
		v1.POST("user-list", apiControllerV!.UserList)
	}

	//Route for api V2
	v2 := router.Group("/api/v2")

	v2.POST("user-list", apiControllerV2.UserList)

	return r
}