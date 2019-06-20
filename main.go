package main

import (
	"fmt"
	"log"
	"net/http"
)

func hiraTable(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Print hiragana table here\n")
}

func kataTable(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Print katakana table here\n")
}

func main() {
    http.HandleFunc("/hira", hiraTable)
    http.HandleFunc("/kata", kataTable)

	fmt.Println("Server start...")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
