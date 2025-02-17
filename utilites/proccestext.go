package utilites

// ProcessText применяет все нужные преобразования к входящему тексту.
func ProcessText(text string) string {
	// Заменяем метки (hex) и (bin)
	text = ProcessCaseCommands(text)
	text = ReplaceHexToDec(text)
	text = ReplaceBinToDec(text)
	// Исправляем пунктуацию (объединяем многоточия, корректируем пробелы и т.д.)
	text = FixPunctuation(text)
	// Применяем команды изменения регистра (up, low, cap) к предыдущим словам.
	text = ProcessCaseCommands(text)
	// Исправляем артикли "a" / "an"
	text = FixAAnCorrectness(text)
	// Удаляем оставшиеся метки (например, если вдруг что-то осталось)
	text = DelAlltheStuff(text)
	return text
}

func ProcessTextUntilStable(text string) string {
	for {
		newText := ProcessText(text)
		if newText == text {
			break
		}
		text = newText
	}
	return text
}
