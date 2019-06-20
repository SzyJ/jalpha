package handlers

import (
	"fmt"
	"net/http"
	"../mapping"
	"../utils"
	"github.com/gorilla/mux"
	"strings"
)

var rowTableHeader []string
var colTableHeader [15]string
var symbolTable [15][5]string

var tableSeperator string = "\t"

func init() {
    rowTableHeader = []string{"a" , "i" , "u" , "e" , "o", "n"}
	colTableHeader = [15]string {"-", "k", "g", "s", "z", "t", "d", "n", "h", "b", "p", "m", "y", "r", "w"}
    symbolTable = [15][5]string {
        {"a" , "i" , "u" , "e" , "o" },
        {"ka", "ki", "ku", "ke", "ko"},
        {"ga", "gi", "gu", "ge", "go"},
        {"sa", "shi", "su", "se", "so"},
        {"za", "ji", "zu", "ze", "zo"},
        {"ta", "chi", "tu", "te", "to"},
        {"da", "ji ", "du", "de", "do"},
        {"na", "ni", "nu", "ne", "no"},
        {"ha", "hi", "fu", "he", "ho"},
        {"ba", "bi", "bu", "be", "bo"},
        {"pa", "pi", "pu", "pe", "po"},
        {"ma", "mi", "mu", "me", "mo"},
        {"ya", "", "yu", "", "yo"},
        {"ra", "ri", "ru", "re", "ro"},
        {"wa", "", "", "", "wo"},
    }
}

func PrintHiraTable(w http.ResponseWriter, r *http.Request) {
	// Print Header
    fmt.Fprintf(w, term_format.HEADER + term_format.BOLD + tableSeperator + strings.Join(rowTableHeader, tableSeperator) + term_format.ENDC + "\n")

	for row, cols := range symbolTable {
		fmt.Fprintf(w, term_format.HEADER + term_format.BOLD + colTableHeader[row] + term_format.ENDC + tableSeperator)
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
		fmt.Fprintf(w, "Could not find hiragana for: " + term_format.FAIL + term_format.BOLD + searchTerm + term_format.ENDC + "\n")
		return
	}

	fmt.Fprintf(w, term_format.HEADER + term_format.BOLD + searchTerm + ": " + term_format.ENDC + symbol + "\n")
}

func SearchHiraganaUsage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Usage: " + term_format.BOLD + "<url>/hira/s/{enter search here}" + term_format.ENDC + "\n")
}
