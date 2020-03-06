package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.Default()
	g.GET("/:name", handle)
	g.Run()
}

func handle(c *gin.Context) {
	name := c.Param("name")
	fmt.Println(name)
	c.JSON(http.StatusOK, name)
}
