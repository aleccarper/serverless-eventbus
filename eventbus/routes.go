package eventbus

import (
	"os"

	"github.com/gin-gonic/gin"
)

func MountAuthorizedRoute(path string, method string, fn gin.HandlerFunc) *gin.Engine {
	engine := buildEngine()
	group := engine.Group("/")
	group.Use(authorizedHandler())
	setMethodHandlerForGroup(method, path, fn, group)
	return engine
}

func buildEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	return engine
}

func setMethodHandlerForGroup(method string, path string, fn gin.HandlerFunc, group *gin.RouterGroup) {
	switch method {
	case "post":
		{
			group.POST(path, fn)
		}
	case "delete":
		{
			group.DELETE(path, fn)
		}
	}
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}

func authorizedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			respondWithError(401, "Authorization token required", c)
			return
		}

		if token != os.Getenv("AUTHENTICATION_TOKEN") {
			respondWithError(401, "Invalid Authorization token", c)
			return
		}

		c.Next()
	}
}
