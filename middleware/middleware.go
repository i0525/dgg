package middleware

import "github.com/gin-gonic/gin"

func MyMiddleware() gin.HandlerFunc {//中间件

	// Do some initialization logic here

	// Foo()


	return func(c *gin.Context) {

		c.Next()

	}

}