package reloaded

import (
	"regexp"
	"strconv"
	"strings"
	// "golang.org/x/text/cases"
	// "golang.org/x/text/language"
)

func setNums(re *regexp.Regexp, str string, numTypeInt int) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		connector := " "
		if arr[0] == '\n' {
			connector = "\n"
		}
		arr = theTrimSpace(arr)
		for i, v := range arr {
			if v == '(' {
				arr = arr[:i]
			}
		}
		decDigit, _ := strconv.ParseInt(arr, numTypeInt, 64)
		result := strconv.Itoa(int(decDigit))
		return connector + result
	})
}

func fixPunc(re *regexp.Regexp, str string) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		arr = theTrimSpace(arr)
		return arr + " "
	})
}

func theTrimSpace(str string) string {
	result := ""
	for _, v := range str {
		if v == ' ' || v == '\n' {
			continue
		}
		result += string(v)
	}
	return result
}

func fixPunc2(re *regexp.Regexp, str string) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		result := theTrimSpace(arr)
		return result + " "
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
	if str[len(str)-1] == ' ' || str[len(str)-1] == '\n' {
		return tailSpacesCut(str[:len(str)-1])
	}
	return str
}

func CorrectAll(str string) string {
	reHex := regexp.MustCompile(`\s+[a-fA-F0-9]+[\s,!.\[\]{}():;']*\(hex\)`)
	reBin := regexp.MustCompile(`\s+[0-1]+[\s,!.\[\]{}():;']*\(bin\)`)
	reCap := regexp.MustCompile(`[a-zA-Z'\[\](){}]+[\s,!.:;]*\((cap|Cap|CAP)\)`)
	reLow := regexp.MustCompile(`[a-zA-Z'\[\](){}]+[\s,!.:;]*\((low|Low|LOW)\)`)
	reUp := regexp.MustCompile(`[a-zA-Z\'[\](){}]+[\s,!.:;]*\((up|UP|Up)\)`)
	reMultipleChars := regexp.MustCompile(`.*(\((cap|CAP|Cap|UP|up|Up|Low|LOW|low),\s(\d+)\)\s*){2,}`)
	reCapMany := regexp.MustCompile(`(.|\n)*\((cap|Cap|CAP),\s(\d+)\)`)
	reUpMany := regexp.MustCompile(`(.|\n)*\((up|UP|Up),\s(\d+)\)`)
	reLowMany := regexp.MustCompile(`(.|\n)*\((low|Low|LOW),\s(\d+)\)`)
	rePunc := regexp.MustCompile(`[\s^.?!]*[.,,,!,?,:;]\s*`)
	rePunc2 := regexp.MustCompile(`[?!.]\s*[?!.]\s*[?!.]\s*`)
	reQuotes := regexp.MustCompile(`'\s*[^']*\s*'`)
	reAn := regexp.MustCompile(`\s[Aa]\s+\w\w+`)
	//reCluMinus := regexp.MustCompile(`((cap|low|up),\s*-(\d+)\)`)

	words := strings.Split(str, " ")
	switch words[0] {
	case "(low,":
		str = str[9:]
	case "(up,":
		str = str[8:]
	case "(cap,":
		str = str[9:]
	case "(low)":
		str = str[9:]
	case "(up)":
		str = str[8:]
	case "(cap)":
		str = str[9:]
	}
	lowCount := 0
	upCount := 0
	capCount := 0
	for _, v := range words {
		switch v {
		case "(low,":
			lowCount++
		case "(up,":
			upCount++
		case "(cap,":
			capCount++
		}
	}

	result := MultipleChars(reMultipleChars, str, reLowMany, reUpMany, reCapMany)
	result = setNums(reBin, result, 2)
	result = setNums(reHex, result, 16)
	result = SetChars(reCap, result, "cap")
	result = SetChars(reLow, result, "low")
	result = SetChars(reUp, result, "up")
	result = fixPunc(rePunc, result)
	result = fixQuote(reQuotes, result)
	result = FixAn(reAn, result)
	for i := 0; i <= upCount; i++ {
		result = SetCharsMany(reUpMany, result, "up")
	}
	for j := 0; j <= capCount; j++ {
		result = SetCharsMany(reCapMany, result, "cap")
	}
	for k := 0; k <= lowCount; k++ {
		result = SetCharsMany(reLowMany, result, "low")
	}
	result = fixPunc2(rePunc2, result)
	result = EmptyCheck(result)

	return result
}
