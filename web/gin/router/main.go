package main

import "github.com/gin-gonic/gin"

func main() {
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	oldVersion := engine.GET("/v1")
	oldVersion.GET("/student", boy)

	newVersion := engine.GET("/v2")
	newVersion.GET("/boy", boy)
}

func boy(engine *gin.Context) {

}
