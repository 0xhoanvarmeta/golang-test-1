package dtos

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}
