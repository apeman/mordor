package main

import (
    "net/http"
    "log"
	"html/template"

    "github.com/julienschmidt/httprouter"
)

const PORT = "10000"
const maxUploadSize = 10 * 1024 * 1024 // 8 mb
const uploadPath = "./uploads"
var userpicPath = "./userpic"


var tmpl = template.Must(template.ParseGlob("_includes/*.html"))

func main() {

    router := httprouter.New()
	
    router.GET("/", UploadFileHandler)
    router.POST("/upload", UploadFileHandler)
	
	router.GET("/v/:PostId", ViewPost)
    router.GET("/viewall", ViewAllFiles)
    router.DELETE("/del/:PostId", DeleteFile)
	
    router.GET("/favicon.ico", Ignore)
	/*
	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))
	*/
	static := httprouter.New()
	static.ServeFiles("/files/*filepath", http.Dir(uploadPath))
	static.ServeFiles("/userpic/*filepath", http.Dir(userpicPath))
	router.NotFound = static	

	log.Print("Server started on localhost:4000")
	log.Fatal(http.ListenAndServe(PORT, router))
}

