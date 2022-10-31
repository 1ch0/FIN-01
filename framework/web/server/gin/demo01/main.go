package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	url(engine)
	restful(engine)
	post(engine)
	upload_file(engine)
	upload_multi_file(engine)
	formBind(engine)
	jsonBind(engine)
	log.Fatal(engine.Run(":8080"))
}

func url(engine *gin.Engine) {
	engine.GET("/student", func(ctx *gin.Context) {
		name := ctx.Query("name")
		addr := ctx.DefaultQuery("addr", "China")
		ctx.String(http.StatusOK, name+" live in "+addr)
	})
}

func restful(engine *gin.Engine) {
	engine.GET("/student/:name/*addr", func(ctx *gin.Context) {
		name := ctx.Param("name")
		addr := ctx.Param("addr")
		ctx.String(http.StatusOK, name+" live in "+addr)
	})
}

func post(engine *gin.Engine) {
	engine.POST("/student", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		addr := ctx.DefaultPostForm("addr", "China")
		ctx.String(http.StatusOK, name+" live in "+addr)
	})
}

func upload_file(engine *gin.Engine) {
	engine.MaxMultipartMemory = 8 << 20
	engine.POST("/upload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			fmt.Printf("get file error %v\n", err)
			ctx.String(http.StatusInternalServerError, "upload file failed")
		} else {
			ctx.SaveUploadedFile(file, "./data/"+file.Filename)
			ctx.String(http.StatusOK, file.Filename)
		}
	})
}

func upload_multi_file(engine *gin.Engine) {
	engine.POST("/upload_files", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			files := form.File["files"]
			for _, file := range files {
				ctx.SaveUploadedFile(file, "./data/"+file.Filename)
			}
			ctx.String(http.StatusOK, "upload"+strconv.Itoa(len(files))+" files")
		}
	})
}

type Student struct {
	Name string `form:"form_name" json:"json_name" uri:"uri_name" xml:"xml_name" yaml:"yaml_name" binding:"required"`
	Addr string `form:"form_addr" json:"json_addr" uri:"uri_addr" xml:"xml_addr" yaml:"yaml_addr" binding:"required"`
}

func formBind(engine *gin.Engine) {
	engine.POST("/stu/form", func(ctx *gin.Context) {
		var stu Student
		if err := ctx.ShouldBind(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse paramter failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+"  "+stu.Addr)
		}
	})
}

func jsonBind(engine *gin.Engine) {
	engine.POST("/stu/json", func(ctx *gin.Context) {
		var stu Student
		if err := ctx.ShouldBindJSON(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse paramter failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+"  "+stu.Addr)
		}
	})
}
