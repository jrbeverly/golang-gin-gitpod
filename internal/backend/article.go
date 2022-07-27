package backend

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jrbeverly/cobra-cmd-with-docs/internal/data"
)

func showIndexPage(c *gin.Context) {
	articles := data.ListArticles(db)

	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func showArticleCreationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func getArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := data.GetArticleByID(db, articleID); err == nil {
			render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")

		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func createArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := data.NewArticle(db, title, content); err == nil {
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
