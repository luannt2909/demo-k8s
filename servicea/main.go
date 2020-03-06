package main

import (
	"demo_kubernetes/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

func main() {
	//viper.SetDefault("baseurl", "http://localhost:8082")
	config.InitializeConfiguration("servicea", "servicea", "", "")
	baseurl := viper.GetString("baseurl")
	fmt.Println(baseurl)
	g := gin.Default()
	g.GET("/:name", handle)
	g.Run(":8081")
}

func handle(c *gin.Context) {
	name := c.Param("name")
	baseurl := viper.GetString("baseurl")
	fmt.Println(baseurl)
	rsp, err := http.Get(fmt.Sprintf("%s/%s", baseurl, name))
	if err != nil {
		c.AbortWithError(http.StatusOK, err)
		return
	}
	defer rsp.Body.Close()
	bytes, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		c.AbortWithError(http.StatusOK, err)
		return
	}
	fmt.Println(string(bytes))
	c.JSON(http.StatusOK, string(bytes))
}
