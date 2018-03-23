package main

import (
	"github.com/wcauchois/http-example/index"
	"log"
	"net/http"
	"strconv"
)

const port int = 8000

func main() {
	i, err := index.New()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", i)
	go func() {
		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
	}()
	log.Printf("Listening on port %v", port)
	select {} // Sleep forever
}
