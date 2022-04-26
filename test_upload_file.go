package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func upload(ctx *gin.Context) {
	//单文件上传
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	err1 := ctx.SaveUploadedFile(file, file.Filename)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func main() {
	engine := gin.Default()
	engine.POST("/upload", upload)
	engine.Run()
}
