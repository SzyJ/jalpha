package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"./handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/hira", handlers.PrintHiraTable)
	router.HandleFunc("/hira/s", handlers.SearchHiraganaUsage)
	router.HandleFunc("/hira/s/{romanji}", handlers.SearchHiraganaCharacter)

    router.HandleFunc("/kata", handlers.PrintKataTable)
	router.HandleFunc("/kata/s", handlers.SearchKatakanaUsage)
	router.HandleFunc("/kata/s/{romanji}", handlers.SearchKatakanaCharacter)

	fmt.Println("Server start...")
    log.Fatal(http.ListenAndServe(":8081", router))
}
