package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	url(engine)
	post(engine)
	log.Fatal(engine.Run(":8888"))
}

func url(engine *gin.Engine) {
	engine.GET("/api/:token", func(ctx *gin.Context) {
		if ctx.Param("token") != "xReadGroupArgs" {
			ctx.String(http.StatusBadRequest, " not authenticated ")
			return
		}

		command := `./api.sh .`
		cmd := exec.Command("/bin/bash", "-c", command)

		err := cmd.Run()
		if err != nil {
			fmt.Println("Execute Command failed:" + err.Error())
			return
		}
		ctx.String(http.StatusOK, "start api")
	})
}

func post(engine *gin.Engine) {
	engine.POST("/api1", func(ctx *gin.Context) {

		if ctx.PostForm("token") != "xReadGroupArgs" {
			ctx.String(http.StatusBadRequest, " not authenticated ")
			return
		}
		command := `./api.sh .`
		cmd := exec.Command("/bin/bash", "-c", command)

		err := cmd.Run()
		if err != nil {
			fmt.Println("Execute Command failed:" + err.Error())
			return
		}
		ctx.String(http.StatusOK, " live in ")
	})
}
