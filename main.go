package main

import (
	"./handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

var DEFAULT_PORT string = "8080"

func getValidPort() string {
	var port string
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		port = DEFAULT_PORT
	} else {
		port = arguments[0]
		validPort, err := strconv.Atoi(port)

		if err != nil || validPort < 1 || validPort > 65535 {
			fmt.Println("Invalid port received. Using default.")
			port = DEFAULT_PORT
		}
	}

	return port
}

func main() {
	port := getValidPort()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/hira", handlers.PrintHiraTable)
	router.HandleFunc("/hira/s", handlers.SearchHiraganaUsage)
	router.HandleFunc("/hira/s/{romanji}", handlers.SearchHiraganaCharacter)

	router.HandleFunc("/kata", handlers.PrintKataTable)
	router.HandleFunc("/kata/s", handlers.SearchKatakanaUsage)
	router.HandleFunc("/kata/s/{romanji}", handlers.SearchKatakanaCharacter)

	fmt.Print("Server started on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
