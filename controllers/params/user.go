package params

type UserUpdateParams struct {
	Name  string `json:"name" binding:"required,max=255"`
	Email string `json:"email" binding:"required,max=255"`
}
