package model

import "fmt"

//User model for users
type User struct {
	ID          string     `json:"-"`
	Email       string     `json:"email"`		
	Password    string     `json:"password"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	AccessToken string     `json:"access_token"`
}

//UserRegister model for users
type UserRegister struct {
	Email 		string
	Message     string
}
//GetFullName returns full name of the user
func (u User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

//UserRequest for registration
type UserRequest struct {
	Email       string     `json:"email"`		
	Password    string     `json:"password"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
}
