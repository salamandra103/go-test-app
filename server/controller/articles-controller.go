package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticlesController interface {
	GetArticles(context *gin.Context)
	SetArticles(context *gin.Context)
	DeleteArticles(context *gin.Context)
}

type articlesController struct {
}

func (c *articlesController) GetArticles(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"articles": []string{"text 1", "text 2", "text 3"},
	})
}
func (c *articlesController) SetArticles(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"articles": []string{"text 1", "text 2", "text 3"},
	})
}
func (c *articlesController) DeleteArticles(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"articles": []string{"text 1", "text 2", "text 3"},
	})
}

func NewArticlesController() ArticlesController {
	return &articlesController{}
}
