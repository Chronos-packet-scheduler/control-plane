package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	log.Println("control-plane listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
