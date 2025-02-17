package main

import (
	"bufio"
	"fmt"
	"go-reloaded/utilites"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: go run . <input file> <output file>")
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Открываем файл для чтения
	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inFile.Close()

	// Открываем файл для записи
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)

	// Обрабатываем файл построчно
	for scanner.Scan() {
		line := scanner.Text()
		processedLine := utilites.ProcessTextUntilStable(line)
		_, err := writer.WriteString(processedLine + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing output file:", err)
		return
	}

	fmt.Println("Text processing complete. Output saved to", outputFile)
}
