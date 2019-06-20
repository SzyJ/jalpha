package handlers

import "../utils"

var rowTableHeader []string
var colTableHeader [15]string
var symbolTable [15][5]string

var tableSeperator string = "\t"

var defaultStyle string = term_format.ENDC
var headerStyle string = term_format.HEADER + term_format.BOLD
var failStyle string = term_format.FAIL + term_format.BOLD

func init() {
    rowTableHeader = []string{"a" , "i" , "u" , "e" , "o", "n"}
	colTableHeader = [15]string {"-", "k", "g", "s", "z", "t", "d", "n", "h", "b", "p", "m", "y", "r", "w"}
    symbolTable = [15][5]string {
        {"a" , "i" , "u" , "e" , "o" },
        {"ka", "ki", "ku", "ke", "ko"},
        {"ga", "gi", "gu", "ge", "go"},
        {"sa", "shi", "su", "se", "so"},
        {"za", "ji", "zu", "ze", "zo"},
        {"ta", "chi", "tsu", "te", "to"},
        {"da", "ji ", "zu ", "de", "do"},
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
