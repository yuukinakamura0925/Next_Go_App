package handlers

import (
	"Next_Go_App/internal/usecases"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
}

// NewUserHandler は新しい UserHandler を作成し、ルーティングを設定します。
func NewUserHandler(router *gin.Engine, userUsecase usecases.UserUsecase) {
	handler := &UserHandler{userUsecase: userUsecase}

	// ユーザー登録エンドポイント
	router.POST("/users/sign_up", handler.SignUp)
	// ユーザーログインエンドポイント
	router.POST("/users/sign_in", handler.SignIn)
}

// SignUp は新しいユーザーを登録します。
func (h *UserHandler) SignUp(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := h.userUsecase.SignUp(context.Background(), req.Email, req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "sign up failed"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// SignIn はユーザーの認証を行い、JWTトークンを返します。
func (h *UserHandler) SignIn(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := h.userUsecase.SignIn(context.Background(), req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "sign in failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
