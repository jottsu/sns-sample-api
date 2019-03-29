package params

type PostCreateParams struct {
	Text string `json:"text" binding:"required,max=255"`
}
