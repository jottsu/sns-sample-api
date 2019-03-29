package params

type SignupParams struct {
	Name     string `json:"name" binding:"required,max=255"`
	Email    string `json:"email" binding:"required,max=255"`
	Password string `json:"password" binding:"required,max=255"`
}

type SigninParams struct {
	Email    string `json:"email" binding:"required,max=255"`
	Password string `json:"password" binding:"required,max=255"`
}
