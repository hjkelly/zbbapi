package main

import (
	"log"
	"net/http"

	"github.com/hjkelly/zbbapi/handlers/v1"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func main() {
	router := httprouter.New()
	v1.RegisterHandlers(router)

	n := negroni.Classic()
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(":8080", n))
}
