package handlers

import (
	"errors"
	"net/http"
	"test-1/internal/dtos"
	"test-1/internal/services"
	"test-1/pkg/message"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userService *services.UserService
}

func Constructor(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var createUserRequest dtos.CreateUserRequest

	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				out[fe.Field()] = message.MsgForTag(fe)
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})

			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err := h.userService.CreateUser(createUserRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
