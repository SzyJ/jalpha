package mapping

var katakanaMap = map[string]string {
    "a"  : "ア",
    "i"  : "イ",
    "u"  : "ウ",
    "e"  : "エ",
    "o"  : "オ",
    "ka" : "カ",
    "ki" : "キ",
    "ku" : "ク",
    "ke" : "ケ",
    "ko" : "コ",
    "ga" : "ガ",
    "gi" : "ギ",
    "gu" : "グ",
    "ge" : "ゲ",
    "go" : "ゴ",
    "sa" : "サ",
    "shi": "シ",
    "su" : "ス",
    "se" : "セ",
    "so" : "ソ",
    "za" : "ザ",
    "ji" : "ジ",
    "zu" : "ズ",
    "ze" : "ゼ",
    "zo" : "ゾ",
    "ta" : "タ",
    "chi": "チ",
    "tsu" :"ツ",
    "te" : "テ",
    "to" : "ト",
    "da" : "ダ",
    "ji ": "ヂ",
    "zu ": "ヅ",
    "de" : "デ",
    "do" : "ド",
    "na" : "ナ",
    "ni" : "ニ",
    "nu" : "ヌ",
    "ne" : "ネ",
    "no" : "ノ",
    "ha" : "ハ",
    "hi" : "ヒ",
    "fu" : "フ",
    "he" : "ヘ",
    "ho" : "ホ",
    "ba" : "バ",
    "bi" : "ビ",
    "bu" : "ブ",
    "be" : "ベ",
    "bo" : "ボ",
    "pa" : "パ",
    "pi" : "ピ",
    "pu" : "プ",
    "pe" : "ペ",
    "po" : "ポ",
    "ma" : "マ",
    "mi" : "ミ",
    "mu" : "ム",
    "me" : "メ",
    "mo" : "モ",
    "ya" : "ヤ",
    "yu" : "ユ",
    "yo" : "ヨ",
    "ra" : "ラ",
    "ri" : "リ",
    "ru" : "ル",
    "re" : "レ",
    "ro" : "ロ",
    "wa" : "ワ",
    "wo" : "ヲ",
    "n"  : "ン",
}

/*
 * Maps sequence of roman characters to a katakana symbol
 *
 * Params
 * romanji string sequence or roman characters equivilent to a katakana symbol
 *
 * returns Empty string if invalid romanji given. Katakana symbol otherwise
 */
func GetKatakanaSymbol(romanji string) string {
	return katakanaMap[romanji]
}
