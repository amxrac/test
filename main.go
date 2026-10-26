package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	name := os.Getenv("NAME")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s from :8000", name)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
