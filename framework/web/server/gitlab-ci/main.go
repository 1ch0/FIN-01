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
	apiserver(engine)
	webserver(engine)
	log.Fatal(engine.Run(":18888"))
}

func apiserver(engine *gin.Engine) {
	engine.GET("/apiserver/:token", func(ctx *gin.Context) {
		if ctx.Param("token") != "xReadGroupArgs" {
			ctx.String(http.StatusBadRequest, " not authenticated ")
			return
		}

		command := `./apiserver.sh .`
		cmd := exec.Command("/bin/bash", "-c", command)

		err := cmd.Run()
		if err != nil {
			fmt.Println("Execute Command failed:" + err.Error())
			return
		}
		ctx.String(http.StatusOK, "restart apiserver...")
	})
}

func webserver(engine *gin.Engine) {
	engine.GET("/webserver/:token", func(ctx *gin.Context) {
		if ctx.Param("token") != "kSmallestPairs2" {
			ctx.String(http.StatusBadRequest, " not authenticated ")
			return
		}

		command := `./webserver.sh .`
		cmd := exec.Command("/bin/bash", "-c", command)

		err := cmd.Run()
		if err != nil {
			fmt.Println("Execute Command failed:" + err.Error())
			return
		}
		ctx.String(http.StatusOK, "restart webserver...")
	})
}
