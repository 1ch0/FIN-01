package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	xml(engine)
	html(engine)
	redirect(engine)
	log.Fatal(engine.Run(":8080"))
}

func xml(engine *gin.Engine) {
	var stu struct {
		Name string
		Addr string
	}
	stu.Name = "zcy"
	stu.Addr = "bj"
	engine.GET("/user/xml", func(ctx *gin.Context) {
		ctx.XML(http.StatusOK, gin.H{"name": "zvy", "addr": "bj"})
	})
}

func html(engine *gin.Engine) {
	engine.LoadHTMLFiles("gin/reponse/static/template.html")
	engine.GET("/user/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template.html", gin.H{"title": "用户信息", "name": "dev", "addr": "sh"})
	})
}

func redirect(engine *gin.Engine) {
	engine.GET("/not_exists", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/user/html")
	})
}
