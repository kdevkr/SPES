package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func handler(c *gin.Context) {
	c.String(200, "%s\n", c.Request.UserAgent())
}

func main() {
	r := gin.Default()
	r.GET("/", handler)
	log.Fatal(r.Run(":8080"))
}
