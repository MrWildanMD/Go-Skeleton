package v2service

import (
	apihelper "github.com/MrWildanMD/Go-Skeleton/apihelpers"
	"github.com/MrWildanMD/Go-Skeleton/models"
	res "github.com/MrWildanMD/Go-Skeleton/resources/api/v2"
)

//UserService struct
type UserService struct {
	User models.User
}

//UserList func returns the list of users
func (us *UserService) UserList() map[string]interface{} {
	user := us.User

	userData := res.UserResponse{
		ID:    user.ID,
		Name:  "test",
		Email: "test@gmail.com",
	}

	response := apihelper.Message(0, "This is from version 2 api")
	response["data"] = userData
	return response
}
