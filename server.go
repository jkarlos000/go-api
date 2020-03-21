package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":3600"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Servidor se esta ejecutando ...")
	})
	log.Println("Server listen on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
