package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yo it's working!\n")
	fmt.Fprintf(w, "Request sent at %s", time.Now())
}

func main() {
	port := "8080"
	fmt.Println("Server running on port", port)
	http.HandleFunc("/", greet)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
