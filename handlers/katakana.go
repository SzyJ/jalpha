package handlers

import (
	"fmt"
	"net/http"
	"../mapping"
	"github.com/gorilla/mux"
	"strings"
	"../utils"
)

func PrintKataTable(w http.ResponseWriter, r *http.Request) {
	// Print Header
    fmt.Fprintf(w, headerStyle + tableSeperator + strings.Join(rowTableHeader, tableSeperator) + defaultStyle + "\n")

	for row, cols := range symbolTable {
		fmt.Fprintf(w, headerStyle + colTableHeader[row] + defaultStyle + tableSeperator)
		for col, _ := range cols {
			fmt.Fprintf(w, mapping.GetKatakanaSymbol(symbolTable[row][col]) + tableSeperator)
			if (row == 0 && col == (len(cols) - 1)) {
				fmt.Fprintf(w, mapping.GetKatakanaSymbol("n"))
			}
		}
		fmt.Fprintf(w, "\n")
	}
}

func SearchKatakanaCharacter(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	searchTerm := vars["romanji"]

	var symbol string

	if (searchTerm == "ji") {
		symbol = mapping.GetKatakanaSymbol("ji") + " or " + mapping.GetKatakanaSymbol("ji ")
	} else if (searchTerm == "zu") {
		symbol = mapping.GetKatakanaSymbol("zu") + " or " + mapping.GetKatakanaSymbol("zu ")
	} else {
		symbol = mapping.GetKatakanaSymbol(searchTerm)
	}

	if (symbol == "") {
		fmt.Fprintf(w, "Could not find katakana for: " + failStyle + searchTerm + defaultStyle + "\n")
		return
	}

	fmt.Fprintf(w, headerStyle + searchTerm + ": " + defaultStyle + symbol + "\n")
}

func SearchKatakanaUsage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Usage: " + term_format.BOLD + "<url>/kata/s/{enter search here}" + defaultStyle + "\n")
}
