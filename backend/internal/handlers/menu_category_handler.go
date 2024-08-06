package handlers

import (
	"Next_Go_App/ent"
	"Next_Go_App/internal/usecases"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuCategoryHandler struct {
	menuCategoryUsecase usecases.IMenuCategoryUsecase
}

func NewMenuCategoryHandler(api *gin.RouterGroup, menuCategoryUsecase usecases.IMenuCategoryUsecase) {
	handler := &MenuCategoryHandler{menuCategoryUsecase: menuCategoryUsecase}

	// メニューカテゴリー登録エンドポイント
	api.POST("/menu_categories", handler.CreateMenuCategory)
	// メニューカテゴリー一覧取得エンドポイント
	api.GET("/menu_categories", handler.GetMenuCategoryList)
	// メニューカテゴリー詳細取得エンドポイント
	api.GET("/menu_categories/:id", handler.GetMenuCategoryDetail)
	// メニューカテゴリー更新エンドポイント
	api.PATCH("/menu_categories/:id", handler.UpdateMenuCategory)
	// メニューカテゴリー削除エンドポイント
	api.DELETE("/menu_categories/:id", handler.DeleteMenuCategory)
}

func (h *MenuCategoryHandler) CreateMenuCategory(c *gin.Context) {

	// コンテキストから user_id を取得
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found"})
		return
	}

	var req ent.MenuCategory

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()}) // エラー詳細をレスポンスに含める
		return
	}

	menuCategory, err := h.menuCategoryUsecase.CreateMenuCategory(context.Background(), userID.(int), req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "create menu category failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"menu_category": menuCategory})
}

func (h *MenuCategoryHandler) GetMenuCategoryList(c *gin.Context) {
	// コンテキストから user_id を取得
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found"})
		return
	}

	menuCategories, err := h.menuCategoryUsecase.GetMenuCategoryList(context.Background(), userID.(int))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "get all menu categories failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"menu_categories": menuCategories})
}

func (h *MenuCategoryHandler) GetMenuCategoryDetail(c *gin.Context) {
	menuCategoryIDStr := c.Param("id")
	menuCategoryID, err := strconv.Atoi(menuCategoryIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Menu Category ID"})
		return
	}

	menuCategory, err := h.menuCategoryUsecase.GetMenuCategoryDetail(context.Background(), menuCategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "get menu category detail failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"menu_category": menuCategory})
}

func (h *MenuCategoryHandler) UpdateMenuCategory(c *gin.Context) {
	menuCategoryIDStr := c.Param("id")
	menuCategoryID, err := strconv.Atoi(menuCategoryIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Menu Category ID"})
		return
	}

	var req ent.MenuCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()}) // エラー詳細をレスポンスに含める
		return
	}

	menuCategory, err := h.menuCategoryUsecase.UpdateMenuCategory(context.Background(), menuCategoryID, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "update menu category failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"menu_category": menuCategory})
}

func (h *MenuCategoryHandler) DeleteMenuCategory(c *gin.Context) {
	menuCategoryIDStr := c.Param("id")
	menuCategoryID, err := strconv.Atoi(menuCategoryIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Menu Category ID"})
		return
	}

	err = h.menuCategoryUsecase.DeleteMenuCategory(context.Background(), menuCategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "delete menu category failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Menu category deleted successfully"})
}
