// Hello world HTTP. Só existe para ter algo rodando atrás do HPA.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	fmt.Fprintf(w, "hello world from %s\n", host)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("ouvindo em :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
