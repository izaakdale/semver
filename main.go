package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	log.Printf("up and running\n")
	http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")), mux)
}
