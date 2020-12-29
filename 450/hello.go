package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is ...
type Hello struct {
	l *log.logger
}

func newHello(l *log.logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println("Hello")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s", d)
}
