package routes

import (
	"github.com/gin-gonic/gin"
)

//MainRouter is the main Gin router
// instance
var MainRouter = gin.Default()

const (
	//APIV is the current version string for
	// our api
	APIV = "/api/v2.0/"
)

func init() {
	// TODO:
}
func index(c *gin.Context) {
	c.JSON(200, gin.H{"index route": "main index route not handle"})
}

// StartServer will start the server on
// port 8080 and add all the sub routes
func StartServer() {
	addToMainRouter()
	addUserRoutes()
	MainRouter.GET("/", index)
	MainRouter.Run(":8080")
}
