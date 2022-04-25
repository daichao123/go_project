package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func testTemplate1() {
	name := "test"
	testTemplate := "test hello :{{.}}"
	parse, _ := template.New("testTemplate").Parse(testTemplate)
	err := parse.Execute(os.Stdout, name)
	if err != nil {
		log.Fatal(err)
	}
}

func templateHtml(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("test.html")
	if err != nil {
		log.Fatal(err)
	}
	err1 := files.Execute(w, "test output")
	if err1 != nil {
		log.Fatal(err1)
	}
}

func testServer() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/testTemp", templateHtml)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//testTemplate1()
	testServer()
}
