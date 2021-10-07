package models

//Your Any Model goes here
//this only example
type User struct {
	AnyModel
	Name  string `gorm:"type:varchar(50)" json:"name" validate:"required"`
	Email string `gorm:"type:varchar(50)" json:"email" validate:"required,email"`
}

//TableName return name of database table
func (u *User) TableName() string {
	return "user"
}
