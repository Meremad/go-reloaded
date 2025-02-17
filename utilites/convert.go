package utilites

import (
	"fmt"
	"regexp"
	"strconv"
)

// ReplaceHexToDec ищет шестнадцатеричные числа, за которыми следует метка (hex),
// и заменяет их десятичным представлением.
func ReplaceHexToDec(text string) string {
	re := regexp.MustCompile(`\b([0-9A-Fa-f]+)\b(?:\s*\(hex\))+`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match
		}
		hexNum := submatches[1]
		decNum, err := strconv.ParseInt(hexNum, 16, 64)
		if err != nil {
			return match
		}
		return fmt.Sprintf("%d", decNum)
	})
}

// ReplaceBinToDec ищет двоичные числа, за которыми следует метка (bin),
// и заменяет их десятичным представлением.
func ReplaceBinToDec(text string) string {
	re := regexp.MustCompile(`\b([01]+)\b(?:\s*\(bin\))+`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match
		}
		binNum := submatches[1]
		decNum, err := strconv.ParseInt(binNum, 2, 64)
		if err != nil {
			return match
		}
		return fmt.Sprintf("%d", decNum)
	})
}

// DelAlltheStuff удаляет оставшиеся служебные метки: (hex), (bin), (up), (low), (cap)
// с необязательным параметром, если таковой указан.
func DelAlltheStuff(text string) string {
	re := regexp.MustCompile(`\(\s*(hex|bin|up|low|cap)(?:,\s*\d+)?\s*\)`)
	return re.ReplaceAllString(text, "")
}
