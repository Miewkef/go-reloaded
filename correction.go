package reloaded

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
	// "golang.org/x/text/cases"
	// "golang.org/x/text/language"
)

func setNums(re *regexp.Regexp, str string, numTypeInt int) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		last6Runes := len(arr) - 6
		arr = arr[:last6Runes]
		arr = strings.Trim(arr, " ")
		arr = strings.Trim(arr, ",")
		arr = strings.Trim(arr, ".")
		arr = strings.Trim(arr, "(")
		arr = strings.Trim(arr, ")")
		arr = strings.Trim(arr, "[")
		arr = strings.Trim(arr, "]")
		arr = strings.Trim(arr, "!")
		arr = strings.Trim(arr, ":")
		arr = strings.Trim(arr, ";")
		arr = strings.Trim(arr, "'")
		decDigit, _ := strconv.ParseInt(arr, numTypeInt, 64)
		result := strconv.Itoa(int(decDigit))
		return result
	})
}

func setChars(re *regexp.Regexp, str string, charType string) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		lastRunes := len(arr) - 6
		if len(charType) == 2 {
			lastRunes++
		}
		arr = arr[:lastRunes]
		// var caser cases.Caser //{without deprication}
		switch charType {
		case "cap":
			// caser = cases.Title(language.English) //{without deprication}
			arr = strings.ToLower(arr)
			arr = strings.ToUpper(string(arr[0])) + arr[1:]
		case "low":
			// caser = cases.Title(language.English) //{without deprication}
			arr = strings.ToLower(arr)
		case "up":
			// caser = cases.Title(language.English) //{without deprication}
			arr = strings.ToUpper(arr)
		}
		// arr = caser.String(arr) //{without deprication}
		return arr
	})
}

func fixPunc(re *regexp.Regexp, str string) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		connector := " "
		if arr[len(arr)-1] == 10 {
			connector = ""
		}
		arr = strings.Trim(arr, " ")
		return arr + connector
	})
}

func fixQuote(re *regexp.Regexp, str string) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		headCuttedStr := headSpacesCut(arr[1:])
		return "'" + tailSpacesCut(headCuttedStr[:len(headCuttedStr)-1]) + "'"
	})
}

func headSpacesCut(str string) string {
	if str[0] == ' ' {
		return headSpacesCut(str[1:])
	}
	return str
}

func tailSpacesCut(str string) string {
	if str[len(str)-1] == ' ' {
		return headSpacesCut(str[:len(str)-1])
	}
	return str
}

func fixAn(re *regexp.Regexp, str string) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		letter := arr[len(arr)-1]
		switch letter {
		case 'a':
			arr = arr[:2] + "n" + arr[2:]
		case 'e':
			arr = arr[:2] + "n" + arr[2:]
		case 'i':
			arr = arr[:2] + "n" + arr[2:]
		case 'o':
			arr = arr[:2] + "n" + arr[2:]
		case 'u':
			arr = arr[:2] + "n" + arr[2:]
		case 'h':
			arr = arr[:2] + "n" + arr[2:]
		case 'A':
			arr = arr[:2] + "n" + arr[2:]
		case 'E':
			arr = arr[:2] + "n" + arr[2:]
		case 'I':
			arr = arr[:2] + "n" + arr[2:]
		case 'O':
			arr = arr[:2] + "n" + arr[2:]
		case 'U':
			arr = arr[:2] + "n" + arr[2:]
		case 'H':
			arr = arr[:2] + "n" + arr[2:]
		}

		return arr
	})
}

func setCharsMany(re *regexp.Regexp, str string, charType string) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		wordsToChange := 0
		startFrom := 0
		countSpace := 0
		var arrToChange string
		var i int
		var cutHere int
		outOfRange := false
		for i = len(arr) - 1; i >= 0; i-- {
			if arr[i] == ' ' {
				wordsToChange, _ = strconv.Atoi(arr[i+1 : len(arr)-1])
				break
			}
		}
		for i = len(arr) - 1; i >= 0; i-- {
			if arr[i] == '(' {
				startFrom = i
				break
			}
		}
		for i = startFrom; i >= 0; i-- {
			if unicode.IsLetter(rune(arr[i])) {
				startFrom = i
				break
			}
		}
		for i = startFrom; i > 0; i-- {
			if arr[i] == ' ' {
				countSpace++
				for i > 0 {
					if unicode.IsLetter(rune(arr[i])) {
						break
					}
					i--
				}
			}
			if i == 1 {
				if countSpace == wordsToChange {
					arrToChange = arr[i+1:]
					break
				}
				arrToChange = arr
				outOfRange = true
				break
			}
			if countSpace == wordsToChange {
				arrToChange = arr[i+1:]
				break
			}
		}
		switch charType {
		case "up":
			arrToChange = strings.ToUpper(arrToChange)
		case "cap":
			arrToChange = strings.ToLower(arrToChange)
			for i := 0; i < len(arrToChange); i++ {
				if unicode.IsLetter(rune(arrToChange[i])) {
					arrToChange = arrToChange[:i] + strings.ToUpper(string(arrToChange[i])) + arrToChange[i+1:]
					for i < len(arrToChange) {
						if arrToChange[i] == '\'' {
							i++
							continue
						}
						if !unicode.IsLetter(rune(arrToChange[i])) {
							break
						}
						i++
					}
				}
			}
		case "low":
			arrToChange = strings.ToLower(arrToChange)
		}
		for j := len(arrToChange) - 1; j >= 0; j-- {
			if arrToChange[j] == '(' {
				cutHere = j
				break
			}
		}
		if outOfRange {
			return arrToChange[:len(arrToChange)-9]
		}
		return arr[:i+1] + arrToChange[:cutHere-1]
	})
}

func CorrectAll(str string) string {
	reHex := regexp.MustCompile(`[a-fA-F0-9]+[\s,!.\[\]{}():;']*\(hex\)`)
	reBin := regexp.MustCompile(`[0-1]+[\s,!.\[\]{}():;']*\(bin\)`)
	reCap := regexp.MustCompile(`[a-zA-Z'\[\](){}]+[\s,!.:;]*\(cap\)`)
	reLow := regexp.MustCompile(`[a-zA-Z'\[\](){}]+[\s,!.:;]*\(low\)`)
	reUp := regexp.MustCompile(`[a-zA-Z\'[\](){}]+[\s,!.:;]*\(up\)`)
	reCapMany := regexp.MustCompile(`.*\(cap,\s(\d+)\)`)
	reUpMany := regexp.MustCompile(`.*\(up,\s(\d+)\)`)
	reLowMany := regexp.MustCompile(`.*\(low,\s(\d+)\)`)
	rePunc := regexp.MustCompile(`[\s^.?!]*[.,,,!,?,:;]\s*`)
	reQuotes := regexp.MustCompile(`'\s*[^']*\s*'`)
	reAn := regexp.MustCompile(`\s[Aa]\s+\w\w+`)

	result := setNums(reBin, str, 2)
	result = setNums(reHex, result, 16)
	result = setChars(reCap, result, "cap")
	result = setChars(reLow, result, "low")
	result = setChars(reUp, result, "up")
	result = fixPunc(rePunc, result)
	result = fixQuote(reQuotes, result)
	result = fixAn(reAn, result)
	result = setCharsMany(reUpMany, result, "up")
	result = setCharsMany(reCapMany, result, "cap")
	result = setCharsMany(reLowMany, result, "low")

	return result
}
