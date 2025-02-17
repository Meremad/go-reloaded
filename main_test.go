package main

import (
	"go-reloaded/utilites"
	"os"
	"testing"
)

// Тест для проверки обработки текста
func TestProcessText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Multiple hex conversions",
			input:    "FF (hex)(hex)(hex)",
			expected: "255",
		},
		{
			name:     "Mixed conversions",
			input:    "101 (bin)(hex)",
			expected: "5",
		},
		{
			name:     "Hex and bin conversions",
			input:    "a (hex)(bin)",
			expected: "2",
		},
		{
			name:     "Mixed conversions",
			input:    "1010 (bin)(hex)",
			expected: "16",
		},
		{
			name:     "Punctuation and correctness",
			input:    "A hour has passed. It is a honor to meet you.",
			expected: "An hour has passed. It is an honor to meet you.",
		},
		{
			name:     "Hex and Bin conversion",
			input:    "1E (hex) files were added. It has been 10 (bin) years.",
			expected: "30 files were added. It has been 2 years.",
		},
		{
			name:     "Up, Low, Cap tags",
			input:    "Ready, set, go (up) ! I should stop SHOUTING (low). Welcome to the Brooklyn bridge (cap).",
			expected: "Ready, set, GO! I should stop shouting. Welcome to the Brooklyn Bridge.",
		},
		{
			name:     "Punctuation and quotes",
			input:    "I was sitting over there ,and then BAMM !! I was thinking ... You were right. I am exactly how they describe me: ' awesome '.",
			expected: "I was sitting over there, and then BAMM!! I was thinking... You were right. I am exactly how they describe me: 'awesome'.",
		},
		{
			name:     "A/An correction",
			input:    "There it was. A amazing rock! A hour has passed.",
			expected: "There it was. An amazing rock! An hour has passed.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utilites.ProcessText(tt.input)
			if result != tt.expected {
				t.Errorf("Test %s failed:\nExpected: %s\nGot: %s", tt.name, tt.expected, result)
			}
		})
	}
}

// Тест для проверки записи и чтения файлов
func TestFileProcessing(t *testing.T) {
	inputFile := "test_input.txt"
	outputFile := "test_output.txt"

	// Создаем тестовый входной файл
	inputContent := "1E (hex) files were added. It has been 10 (bin) years."
	err := os.WriteFile(inputFile, []byte(inputContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}
	defer os.Remove(inputFile)

	// Запускаем программу
	os.Args = []string{"cmd", inputFile, outputFile}
	main()

	// Читаем результат
	outputContent, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	defer os.Remove(outputFile)

	expected := "30 files were added. It has been 2 years."
	if string(outputContent) != expected {
		t.Errorf("File processing failed:\nExpected: %s\nGot: %s", expected, string(outputContent))
	}
}
