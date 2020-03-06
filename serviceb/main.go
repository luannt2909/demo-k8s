package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.Default()
	g.GET("/:name", handle)
	g.Run(":8082")
}

func handle(c *gin.Context) {
	name := c.Param("name")
	fmt.Println(name)
	c.JSON(http.StatusOK, fmt.Sprintf("i love %s", name))
}
