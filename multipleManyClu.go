package reloaded

import (
	"regexp"
	"unicode"
)

func MultipleChars(re *regexp.Regexp, str string, reLowMany *regexp.Regexp, reUpMany *regexp.Regexp, reCapMany *regexp.Regexp) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		connector := ""
		for i := len(arr) - 1; i > len(arr)-10; i-- {
			switch arr[i] {
			case ' ':
				connector = " " + connector
			case '\n':
				connector = "\n" + connector
			}
		}
		isCaseUp := 0
		count := 0
		twoFuncsInd := 0
		threeFuncsInd := 0
		result := ""
		for i := len(arr) - 1; i >= 0; i-- {
			if arr[i] == '(' {
				if arr[i+1:i+4] == "cap" || arr[i+1:i+4] == "low" || arr[i+1:i+3] == "up" {
					count++
					if count == 3 {
						if unicode.IsDigit(rune(arr[i+7])) || unicode.IsDigit(rune(arr[i+6])) {
							threeFuncsInd = i
							break
						}
						count--
					}
					if twoFuncsInd == 0 && count == 2 {
						twoFuncsInd = i
					}
				}
			}
		}
		switch count {
		case 3:
			switch arr[threeFuncsInd+1] {
			case 'l':
				result = SetCharsMany(reLowMany, arr[:threeFuncsInd+8], "low")
			case 'L':
				result = SetCharsMany(reLowMany, arr[:threeFuncsInd+8], "low")
			case 'u':
				isCaseUp--
				result = SetCharsMany(reUpMany, arr[:threeFuncsInd+8], "up")
			case 'U':
				isCaseUp--
				result = SetCharsMany(reUpMany, arr[:threeFuncsInd+8], "up")
			case 'c':
				result = SetCharsMany(reCapMany, arr[:threeFuncsInd+8], "cap")
			case 'C':
				result = SetCharsMany(reCapMany, arr[:threeFuncsInd+8], "cap")
			}
			arr = arr[:threeFuncsInd] + arr[threeFuncsInd+9+isCaseUp:]
			if isCaseUp != 0 {
				isCaseUp = 0
			}
			switch arr[threeFuncsInd+1] {
			case 'l':
				result = SetCharsMany(reLowMany, result+arr[threeFuncsInd:threeFuncsInd+8], "low")
			case 'L':
				result = SetCharsMany(reLowMany, result+arr[threeFuncsInd:threeFuncsInd+8], "low")
			case 'u':
				isCaseUp--
				result = SetCharsMany(reUpMany, result+arr[threeFuncsInd:threeFuncsInd+8], "up")
			case 'U':
				isCaseUp--
				result = SetCharsMany(reUpMany, result+arr[threeFuncsInd:threeFuncsInd+8], "up")
			case 'c':
				result = SetCharsMany(reCapMany, result+arr[threeFuncsInd:threeFuncsInd+8], "cap")
			case 'C':
				result = SetCharsMany(reCapMany, result+arr[threeFuncsInd:threeFuncsInd+8], "cap")
			}
			arr = arr[:threeFuncsInd] + arr[threeFuncsInd+9+isCaseUp:]
			if isCaseUp != 0 {
				isCaseUp = 0
			}
			switch arr[threeFuncsInd+1] {
			case 'l':
				result = SetCharsMany(reLowMany, result+arr[threeFuncsInd:threeFuncsInd+8], "low")
			case 'L':
				result = SetCharsMany(reLowMany, result+arr[threeFuncsInd:threeFuncsInd+8], "low")
			case 'u':
				isCaseUp--
				result = SetCharsMany(reUpMany, result+arr[threeFuncsInd:threeFuncsInd+8], "up")
			case 'U':
				isCaseUp--
				result = SetCharsMany(reUpMany, result+arr[threeFuncsInd:threeFuncsInd+8], "up")
			case 'c':
				result = SetCharsMany(reCapMany, result+arr[threeFuncsInd:threeFuncsInd+8], "cap")
			case 'C':
				result = SetCharsMany(reCapMany, result+arr[threeFuncsInd:threeFuncsInd+8], "cap")
			}
		case 2:
			switch arr[twoFuncsInd+1] {
			case 'l':
				result = SetCharsMany(reLowMany, arr[:twoFuncsInd+8], "low")
			case 'L':
				result = SetCharsMany(reLowMany, arr[:twoFuncsInd+8], "low")
			case 'u':
				isCaseUp--
				result = SetCharsMany(reUpMany, arr[:twoFuncsInd+8], "up")
			case 'U':
				isCaseUp--
				result = SetCharsMany(reUpMany, arr[:twoFuncsInd+8], "up")
			case 'c':
				result = SetCharsMany(reCapMany, arr[:twoFuncsInd+8], "cap")
			case 'C':
				result = SetCharsMany(reCapMany, arr[:twoFuncsInd+8], "cap")
			}
			arr = arr[:threeFuncsInd] + arr[threeFuncsInd+9+isCaseUp:]
			switch arr[twoFuncsInd+1] {
			case 'l':
				result = SetCharsMany(reLowMany, result+arr[twoFuncsInd:twoFuncsInd+8], "low")
			case 'L':
				result = SetCharsMany(reLowMany, result+arr[twoFuncsInd:twoFuncsInd+8], "low")
			case 'u':
				isCaseUp--
				result = SetCharsMany(reUpMany, result+arr[twoFuncsInd:twoFuncsInd+8], "up")
			case 'U':
				isCaseUp--
				result = SetCharsMany(reUpMany, result+arr[twoFuncsInd:twoFuncsInd+8], "up")
			case 'c':
				result = SetCharsMany(reCapMany, result+arr[twoFuncsInd:twoFuncsInd+8], "cap")
			case 'C':
				result = SetCharsMany(reCapMany, result+arr[twoFuncsInd:twoFuncsInd+8], "cap")
			}
		}
		return result + connector
	})
}