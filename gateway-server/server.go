package main

import (
	"fmt"

	"github.com/cqroot/garden/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func setDefaultConfig() {
	viper.SetDefault("bind_ip", "0.0.0.0")
	viper.SetDefault("bind_port", "6000")
}

func main() {
	setDefaultConfig()
	common.InitConfig("gateway-server")
	common.InitLogger()

	r := gin.Default()

	v1Group := r.Group("/v1")
	{
		v1Group.GET("/foo", func(c *gin.Context) { c.JSON(200, gin.H{"foo": true}) })
		v1Group.GET("/bar", func(c *gin.Context) { c.JSON(200, gin.H{"bar": true}) })
	}

	r.Run(fmt.Sprintf("%s:%s", viper.GetString("bind_ip"), viper.GetString("bind_port")))
}
