package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nicholasjackson/working/handlers"
)

func main() {

	l := log.New(os.Stdout, "product.api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()

	sm.handle("/", hh)

	s := http.Server()

	http.ListenAndServe(":9090", nil)
}
