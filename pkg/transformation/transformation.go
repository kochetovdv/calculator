package transformation

import (
	"regexp"
)

// SymbolDeleting deletes symbols from string using patterns
func SymbolDeleting(pattern, s string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(s, "")
}

// ReplaceAllSymbols replaces symbols from string using patterns
func ReplaceAllSymbols(pattern, srcS string, toS string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(srcS, toS)
}

// StringsConcatination adds the second string to the end of the first one
func StringsConcatination(s1, s2 string) string {
	return s1 + s2
}

// AddSymbolBeforeAndAfterAnotherSymbol adds symbol before and after symbols. Returns string
func AddSymbolBeforeAndAfterAnotherSymbol(r []rune, symbolToAdd rune, symbols ...rune) string {
	newR := make([]rune, len(r))
	for _, elMainArr := range r {
		addSymbol := false
		for _, elSymbols := range symbols {
			if elMainArr == elSymbols {
				addSymbol = true
				break
			}
		}
		if addSymbol {
			newR = append(newR, symbolToAdd, elMainArr, symbolToAdd)
		} else {
			newR = append(newR, elMainArr)
		}
	}
	return string(newR)
}

// AddSymbolBeforeAnotherSymbol adds symbol before symbols. Returns string
func AddSymbolBeforeAnotherSymbol(r []rune, symbolToAdd rune, symbols ...rune) string {
	newR := make([]rune, len(r))
	for _, elMainArr := range r {
		addSymbol := false
		for _, elSymbols := range symbols {
			if elMainArr == elSymbols {
				addSymbol = true
				break
			}
		}
		if addSymbol {
			newR = append(newR, symbolToAdd, elMainArr)
		} else {
			newR = append(newR, elMainArr)
		}
	}
	return string(newR)
}

// AddSymbolAfterAnotherSymbol adds symbol after symbols. Returns string
func AddSymbolAfterAnotherSymbol(r []rune, symbolToAdd rune, symbols ...rune) string {
	newR := make([]rune, len(r))
	for _, elMainArr := range r {
		addSymbol := false
		for _, elSymbols := range symbols {
			if elMainArr == elSymbols {
				addSymbol = true
				break
			}
		}
		if addSymbol {
			newR = append(newR, elMainArr, symbolToAdd)
		} else {
			newR = append(newR, elMainArr)
		}
	}
	newS := string(newR)
	return newS
}

// Split splits string by delimiter and returns slice of strings
func Split(s string, delimiter string) []string {
	return regexp.MustCompile(delimiter).Split(s, -1)
}
