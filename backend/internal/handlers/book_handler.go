package handlers

import (
	"Next_Go_App/internal/usecases"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookUsecase usecases.BookUsecase
}

func NewBookHandler(api *gin.RouterGroup, bookUsecase usecases.BookUsecase) {
	handler := &BookHandler{bookUsecase: bookUsecase}

	api.POST("/books", handler.CreateBook)
	api.GET("/books", handler.GetAllBooks)
	api.GET("/books/:id", handler.GetBookByID)
	api.PATCH("/books/:id", handler.UpdateBook)
	api.DELETE("/books/:id", handler.DeleteBook)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	fmt.Println("CreateBook")
	fmt.Println(c)

	type NewBookRequest struct {
		Title  string `json:"title" binding:"required"`
		Body   string `json:"body" binding:"required"`
		UserID int    `json:"user_id" binding:"required"`
	}

	var req NewBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Error:", err)                                                               // エラー内容をログに出力
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()}) // エラー詳細をレスポンスに含める
		return
	}

	book, err := h.bookUsecase.CreateBook(context.Background(), req.Title, req.Body, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "create book failed"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"book": book})
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.bookUsecase.GetAllBooks(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "get all books failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"books": books})
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Book ID"})
		return
	}

	book, err := h.bookUsecase.GetBookByID(context.Background(), bookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "message": "Book with specified id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"book": book})
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	type UpdateBookRequest struct {
		Title  string `json:"title" binding:"required"`
		Body   string `json:"body" binding:"required"`
		UserID int    `json:"user_id" binding:"required"`
	}

	var req UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Book ID"})
		return
	}

	book, err := h.bookUsecase.UpdateBook(context.Background(), bookID, req.Title, req.Body, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "update book failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"book": book})
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Book ID"})
		return
	}

	err = h.bookUsecase.DeleteBook(context.Background(), bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "delete book failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
