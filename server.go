package main

import (
	"log"
	"net/http"

	"github.com/hjkelly/zbbapi/handlers/v1"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {
	router := httprouter.New()
	v1.RegisterHandlers(router)

	n := negroni.Classic()
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(":8080", n))
}
