package views

import "github.com/jottsu/sns-sample-api/models"

type SignupView struct {
	Token string   `json:"token"`
	User  UserView `json:"user"`
}

func ToSignupView(token string, user *models.User) *SignupView {
	return &SignupView{
		Token: token,
		User: UserView{
			Name:  user.Name,
			Email: user.Email,
		},
	}
}
