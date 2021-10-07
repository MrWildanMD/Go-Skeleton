package v1resources

//User Response struct
type UserResponse struct {
	ID    uint   `json"id"`
	Name  string `json"name,omitempty"`
	Email string `json"email,omitempty" validate:"required,email"`
}
