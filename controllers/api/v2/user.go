package v2

import (
	"encoding/json"
	apihelper "github.com/MrWildanMD/Go-Skeleton/apihelpers"
	v1services "github.com/MrWildanMD/Go-Skeleton/services/api/v1"
	"github.com/gin-gonic/gin"
)

//Func to give list of users
func UserList(c *gin.Context) {
	var userService v1services.UserService

	//decode the request body into struct and failed if any error occured
	err := json.NewDecoder(c.Request.Body).Decode(&userService.User)
	if err != nil {
		apihelper.Respond(c.Writer, apihelper.Message(1, "Invalid request"))
		return
	}

	//calling service
	resp := userService.UserList()

	//return response using apihelper
	apihelper.Respond(c.Writer, resp)
}
