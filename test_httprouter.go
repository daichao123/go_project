package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "test httpRouter")
}

func hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "test echo params :%v\n", params.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/echo/:name", hello)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
