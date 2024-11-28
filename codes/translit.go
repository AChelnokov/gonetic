package codes

import (
	"unicode"
)

var cyrillicBase = map[int32]string{
	1040: "A",   // А
	1041: "B",   // Б
	1042: "V",   // В
	1043: "G",   // Г
	1044: "D",   // Д
	1045: "E",   // Е
	1025: "E",   // Ё
	1046: "Zh",  // Ж
	1047: "Z",   // З
	1048: "I",   // И
	1049: "J",   // Й
	1050: "K",   // К
	1051: "L",   // Л
	1052: "M",   // М
	1053: "N",   // Н
	1054: "O",   // О
	1055: "P",   // П
	1056: "R",   // Р
	1057: "S",   // С
	1058: "T",   // Т
	1059: "U",   // У
	1060: "F",   // Ф
	1061: "H",   // Х
	1062: "C",   // Ц
	1063: "Ch",  // Ч
	1064: "Sh",  // Ш
	1065: "Sch", // Щ
	1066: "'",   // Ъ
	1067: "Y",   // Ы
	1068: "'",   // Ь
	1069: "E",   // Э
	1070: "Ju",  // Ю
	1071: "Ja",  // Я
	1072: "a",   // а
	1073: "b",   // б
	1074: "v",   // в
	1075: "g",   // г
	1076: "d",   // д
	1077: "e",   // е
	1105: "e",   // ё
	1078: "zh",  // ж
	1079: "z",   // з
	1080: "i",   // и
	1081: "j",   // й
	1082: "k",   // к
	1083: "l",   // л
	1084: "m",   // м
	1085: "n",   // н
	1086: "o",   // о
	1087: "p",   // п
	1088: "r",   // р
	1089: "s",   // с
	1090: "t",   // т
	1091: "u",   // у
	1092: "f",   // ф
	1093: "h",   // х
	1094: "c",   // ц
	1095: "ch",  // ч
	1096: "sh",  // ш
	1097: "sch", // щ
	1098: "",    // ъ
	1099: "y",   // ы
	1100: "",    // ь
	1101: "e",   // э
	1102: "ju",  // ю
	1103: "ja",  // я
	32:   " ",   // space
}

func Translit(str string) string {
	var newStr string

	for _, ch := range str {
		newCh, ok := cyrillicBase[ch]
		if ok {
			newStr = newStr + newCh
		} else {
			newStr = newStr + string(ch)
		}
	}

	return newStr
}

func NeedTranslit(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Cyrillic, r) {
			return true
		}
	}
	return false
}
