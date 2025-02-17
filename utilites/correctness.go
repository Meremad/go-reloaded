package utilites

import (
	"regexp"
	"strings"
)

// processQuotes удаляет лишние пробелы внутри кавычек (как одинарных, так и двойных)
func processQuotes(text, quote string) string {
	// Экранируем кавычку, чтобы её использовать в регулярном выражении
	pattern := regexp.QuoteMeta(quote) + `\s*(.*?)\s*` + regexp.QuoteMeta(quote)
	re := regexp.MustCompile(pattern)
	// Заменяем на кавычку без пробелов внутри
	return re.ReplaceAllString(text, quote+"$1"+quote)
}

func FixPunctuation(text string) string {
	// 1. Объединение многоточий (убираем пробелы между точками)
	re := regexp.MustCompile(`\.(\s*\.)+`)
	text = re.ReplaceAllStringFunc(text, func(s string) string {
		return strings.ReplaceAll(s, " ", "")
	})

	// 2. Объединение групп ! и ? (удаляем пробелы между ними)
	re = regexp.MustCompile(`([!?](?:\s*[!?])+)`)
	text = re.ReplaceAllStringFunc(text, func(s string) string {
		return strings.ReplaceAll(s, " ", "")
	})

	// 3. Удаление пробелов вокруг одиночных знаков препинания (.,!?:;)
	re = regexp.MustCompile(`\s*([.,!?:;])\s*`)
	text = re.ReplaceAllString(text, "$1")

	// 4. Добавление пробела после знаков препинания, если следующий символ не является пробелом или знаком
	re = regexp.MustCompile(`([.,!?:;])([^ .,!?:;])`)
	text = re.ReplaceAllString(text, "$1 $2")

	// 5. Исправление апострофов внутри слов (например, can't, I'm)
	re = regexp.MustCompile(`(\w)\s*'\s*(\w)`)
	text = re.ReplaceAllString(text, "$1'$2")

	// 6. Обработка двойных кавычек – убираем лишние пробелы внутри них
	text = processQuotes(text, `"`)
	// 7. Обработка одинарных кавычек – убираем лишние пробелы внутри них
	text = processQuotes(text, `'`)

	// 8. Финальная очистка: убираем лишние пробелы
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")

	return strings.TrimSpace(text)
}

// FixAAnCorrectness исправляет использование артиклей "a" и "an" в зависимости от
// того, с какой буквы начинается следующее слово.
func FixAAnCorrectness(input string) string {
	silentH := map[string]struct{}{
		"honest": {}, "heir": {}, "honorific": {}, "honor": {}, "herb": {}, "hour": {}, "homage": {},
	}
	exceptions := map[string]struct{}{
		"for": {}, "and": {}, "nor": {}, "but": {}, "or": {}, "so": {}, "yet": {},
	}
	words := strings.Fields(input)
	for i := 0; i < len(words)-1; i++ {
		lowered := strings.ToLower(words[i])
		if lowered != "a" && lowered != "an" {
			continue
		}
		next := words[i+1]

		// Если следующее слово состоит из одного символа, считаем его отдельной буквой – не заменяем артикль.
		if len(next) == 1 {
			continue
		}

		// Если слово из списка исключений — пропускаем замену
		if _, exists := exceptions[strings.ToLower(next)]; exists {
			continue
		}

		// Если слово начинается с гласной или входит в список слов с беззвучным "h"
		shouldBeAn := false
		if _, exists := silentH[strings.ToLower(next)]; exists || strings.ContainsRune("aeiouAEIOU", rune(next[0])) {
			shouldBeAn = true
		}

		if shouldBeAn {
			if words[i][0] == 'A' {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		} else {
			if words[i][0] == 'A' {
				words[i] = "A"
			} else {
				words[i] = "a"
			}
		}
	}
	return strings.Join(words, " ")
}
