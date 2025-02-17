package utilites

import (
	"regexp"
	"strconv"
	"strings"
)

// ProcessCaseCommands ищет команды (up), (low) или (cap) (с необязательным числом) и
// преобразует указанное количество предыдущих слов.
func ProcessCaseCommands(text string) string {
	// Регулярное выражение ищет шаблон: (up) или (low) или (cap) с необязательным числом, например (up, 2)
	re := regexp.MustCompile(`\((up|low|cap)(?:,\s*(\d+))?\)`)
	for {
		loc := re.FindStringSubmatchIndex(text)
		if loc == nil {
			break
		}
		fullStart, fullEnd := loc[0], loc[1]
		op := text[loc[2]:loc[3]]
		count := 1
		if loc[4] != -1 {
			countStr := text[loc[4]:loc[5]]
			if n, err := strconv.Atoi(countStr); err == nil && n > 0 {
				count = n
			}
		}
		// Получаем текст до команды
		prefix := text[:fullStart]
		// Находим все слова в префиксе
		wordRe := regexp.MustCompile(`\b\w+\b`)
		words := wordRe.FindAllStringIndex(prefix, -1)
		if len(words) == 0 {
			// Если нет слов для обработки, просто удаляем команду
			text = prefix + text[fullEnd:]
			continue
		}
		if len(words) < count {
			count = len(words)
		}
		newPrefix := prefix
		// Обрабатываем последние count слов (идём по индексам слева направо)
		for i := len(words) - count; i < len(words); i++ {
			start := words[i][0]
			end := words[i][1]
			originalWord := prefix[start:end]
			var transformed string
			switch op {
			case "up":
				transformed = strings.ToUpper(originalWord)
			case "low":
				transformed = strings.ToLower(originalWord)
			case "cap":
				transformed = Capitalize(originalWord)
			}
			newPrefix = newPrefix[:start] + transformed + newPrefix[end:]
		}
		// Собираем строку без обработанного служебного выражения
		text = newPrefix + text[fullEnd:]
	}
	return text
}

// Capitalize приводит слово к виду: первая буква заглавная, остальные – строчные.
func Capitalize(text string) string {
	if len(text) == 0 {
		return text
	}
	return strings.ToUpper(text[:1]) + strings.ToLower(text[1:])
}
