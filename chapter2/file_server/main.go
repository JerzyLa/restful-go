package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main()  {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("/Users/jerzylasyk/GolandProjects/restful-go/static"))
	log.Fatal(http.ListenAndServe(":8000", router))
}