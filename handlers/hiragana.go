package handlers

import (
	"fmt"
	"net/http"
	"../mapping"
	"github.com/gorilla/mux"
	"strings"
	"../utils"
)

func PrintHiraTable(w http.ResponseWriter, r *http.Request) {
	// Print Header
    fmt.Fprintf(w, headerStyle + tableSeperator + strings.Join(rowTableHeader, tableSeperator) + term_format.ENDC + "\n")

	for row, cols := range symbolTable {
		fmt.Fprintf(w, headerStyle + colTableHeader[row] + term_format.ENDC + tableSeperator)
		for col, _ := range cols {
			fmt.Fprintf(w, mapping.GetHiraganaSymbol(symbolTable[row][col]) + tableSeperator)
			if (row == 0 && col == (len(cols) - 1)) {
				fmt.Fprintf(w, mapping.GetHiraganaSymbol("n"))
			}
		}
		fmt.Fprintf(w, "\n")
	}
}

func SearchHiraganaCharacter(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	searchTerm := vars["romanji"]

	var symbol string

	if (searchTerm == "ji") {
		symbol = mapping.GetHiraganaSymbol("ji") + " or " + mapping.GetHiraganaSymbol("ji ")
	} else if (searchTerm == "zu") {
		symbol = mapping.GetHiraganaSymbol("zu") + " or " + mapping.GetHiraganaSymbol("zu ")
	} else {
		symbol = mapping.GetHiraganaSymbol(searchTerm)
	}

	if (symbol == "") {
		fmt.Fprintf(w, "Could not find hiragana for: " + failStyle + searchTerm + defaultStyle + "\n")
		return
	}

	fmt.Fprintf(w, headerStyle + searchTerm + ": " + defaultStyle + symbol + "\n")
}

func SearchHiraganaUsage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Usage: " + term_format.BOLD + "<url>/hira/s/{enter search here}" + defaultStyle + "\n")
}
