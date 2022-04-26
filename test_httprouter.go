package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func testGin(ctx *gin.Context) {
	name := ctx.Query("name")
	password := ctx.DefaultQuery("password", "test_password")
	if password != "" && name != "" {
		ctx.JSON(200, map[string]interface{}{"name:": name, "password:": password})
	}
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "test httpRouter")
}

func hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "test echo params :%v\n", params.ByName("name"))
}

func main() {
	engine := gin.Default()
	engine.GET("/test", testGin)
	err := engine.Run(":8888")
	if err != nil {
		log.Fatal(err)
	}
	//router := httprouter.New()
	//router.GET("/", index)
	//router.GET("/echo/:name", hello)
	//err := http.ListenAndServe(":8080", router)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
