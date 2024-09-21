package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Credential struct {
	Username string
	Password string
}

func BasicAuth(credentials []Credential) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		log.Println("---basic auth---")
		log.Println(username, password, ok)
		if !ok {
			c.Writer.Header().Set("WWW-Authenticate", "Basic")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		for _, v := range credentials {
			if v.Username == username && v.Password == password {
				c.Next()
				return
			}
		}
		c.Writer.Header().Set("WWW-Authenticate", "Basic")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
