package main

import (
	"fmt"
	"log"
	"net/http"
	"./mapping"
	"strings"
)

// Terminal Colour escapes
var HEADER string     = "\033[95m"
var OKBLUE string     = "\033[94m"
var OKGREEN string    = "\033[92m"
var WARNING string    = "\033[93m"
var FAIL string       = "\033[91m"
var ENDC string       = "\033[0m"
var BOLD string       = "\033[1m"
var UNDERLINE string  = "\033[4m"

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

func printHiraTable(w http.ResponseWriter, r *http.Request) {
	// Print Header
    fmt.Fprintf(w, HEADER + BOLD + tableSeperator + strings.Join(rowTableHeader, tableSeperator) + ENDC + "\n")

	for row, cols := range symbolTable {
		fmt.Fprintf(w, HEADER + BOLD + colTableHeader[row] + ENDC + tableSeperator)
		for col, _ := range cols {
			fmt.Fprintf(w, mapping.GetHiraganaSymbol(symbolTable[row][col]) + tableSeperator)
			if (row == 0 && col == (len(cols) - 1)) {
				fmt.Fprintf(w, mapping.GetHiraganaSymbol("n"))
			}
		}
		fmt.Fprintf(w, "\n")
	}
}

func searchHiraganaCharacter(w http.ResponseWriter, r *http.Request) {

}

func printKataTable(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Print katakana table here\n")
}

func main() {
    http.HandleFunc("/hira", printHiraTable)
    http.HandleFunc("/kata", printKataTable)

	http.HandleFunc("/hira/{id}", handler func(http.ResponseWriter, *http.Request))
	fmt.Println("Server start...")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
