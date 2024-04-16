package main

import (
        "bufio"
        "fmt"
        "os"
        "path/filepath"
        "strconv"
        "strings"
)

func readInput() InputWrapper {
        input := InputWrapper{}
        scanner := bufio.NewScanner(os.Stdin)

        scan := func (label string, scanner *bufio.Scanner) string {
                fmt.Print(label)
                scanner.Scan()
                return scanner.Text()
        }

        input.title = scan("Enter the anime name: ", scanner)
        entryIdString := scan("Enter the anime entry id: ", scanner)
        rangeStartString := scan("Enter the range start: ", scanner)
        rangeEndString := scan("Enter the range end: ", scanner)
        qualityString := scan("Enter the quality (ex. 480/720/1080): ", scanner)
        threadCountString := scan("How many simultaneous downloads do you want to hold: ", scanner)

        atoiWithHandling := func(value string) int {
                num, err := strconv.Atoi(value)
                if err != nil {
                        message := fmt.Sprintf("Invalid input for %s: %v", value, err)
                        ErrorLogger.Fatal(message)
                }
                return num
        }

        input.entryId = atoiWithHandling(entryIdString)
        input.rangeStart = atoiWithHandling(rangeStartString)
        input.rangeEnd = atoiWithHandling(rangeEndString)
        input.quality = atoiWithHandling(qualityString)
        input.threadCount = atoiWithHandling(threadCountString)

        return input
}

func setupDirectory(input InputWrapper) string {
        entryId := strconv.Itoa(input.entryId)
        dir := strings.Join([]string{input.title, entryId}, "-")
        path := filepath.Join(".", dir)
        
        InfoLogger.Printf("Looking for \"%s\" directory", path)
        if _, err := os.Stat(path); os.IsNotExist(err) {
                InfoLogger.Printf("\"%s\" not found, creating it", path)
                err := os.MkdirAll(path, 0755)
                if err != nil {
                        ErrorLogger.Fatalf("Error creating directory: %v\n", err)
                }
                InfoLogger.Printf("\"%s\" created", path)
        }

        return path
}
