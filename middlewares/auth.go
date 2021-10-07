package middlewares

import "github.com/gin-gonic/gin"

/*
UserMiddlewares func to add auth
*/

func UserMiddlewares() gin.Handlerfunc {
	return func(c *gin.Context) {
		//Your middleware code

		c.Next()
	}
}
