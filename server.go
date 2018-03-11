package main

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloServer)

	n := negroni.Classic()
	n.UseHandler(mux)
	log.Fatal(http.ListenAndServe(":8080", n))
}
