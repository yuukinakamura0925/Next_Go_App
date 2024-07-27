package handlers

import (
	"Next_Go_App/internal/usecases"
	"context"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(router *gin.Engine, userUsecase usecases.UserUsecase) {
	handler := &UserHandler{userUsecase: userUsecase}

	router.POST("users/sign_up", handler.SignUp)
	router.POST("users/sign_in", handler.SignIn)
}

func (h *UserHandler) SignUp(c *gin.Context) {
	type SignUpRequest struct {
		Email    string `json:"email" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// reqという変数をSignUpRequest構造体型で宣言しています。この変数にリクエストデータがマッピングされます。
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := h.userUsecase.SignUp(context.Background(), req.Email, req.Name, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "message": "sign up failed"})
		return
	}
	c.JSON(201, gin.H{"user": user})
}

func (h *UserHandler) SignIn(c *gin.Context) {
	type SignInRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := h.userUsecase.SignIn(context.Background(), req.Email, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "message": "sign in failed"})
		return
	}
	c.JSON(200, gin.H{"user": user})
}
