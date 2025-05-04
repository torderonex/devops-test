package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	ip := r.RemoteAddr
	author := os.Getenv("AUTHOR")

	fmt.Fprintf(w, "Host: %s\nIP: %s\nAuthor: %s\n", host, ip, author)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
