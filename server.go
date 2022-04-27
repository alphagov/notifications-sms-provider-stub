package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/mmg", MmgEndpoint)
	http.HandleFunc("/firetext", FiretextEndpoint)
	port := getenv("PORT", "6300")
	log.Printf("Listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
