package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jrbeverly/cobra-cmd-with-docs/internal/data"
)

var router *gin.Engine
var db *data.InMemoryDatabase

func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

func InitializeRoutes() *gin.Engine {
	router = gin.Default()
	db = data.NewDatabase()

	router.Use(setUserStatus())

	router.GET("/", showIndexPage)
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)
		userRoutes.GET("/logout", ensureLoggedIn(), logout)
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", getArticle)
		articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)
		articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	}

	return router
}
