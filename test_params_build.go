package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Mobile   string `form:"mobile"`
	Address  string `form:"address"`
}

func register(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBind(&user)
	if err != nil {
		log.Fatal(err)
	}
	//ctx.String(200, "userData:%s", user)
	ctx.JSON(200, user)

}

func main() {
	engine := gin.Default()
	engine.POST("/register", register)
	engine.Run()
}
