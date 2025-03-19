package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/timeout", timeoutHandler)

	port := 8080
	fmt.Printf("Timeout test server running at http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// timeoutHandler never sends a response, causing client timeout
func timeoutHandler(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w)

	fmt.Println("Received request to /api/timeout - This will never respond")

	select {}
}

func setupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}
