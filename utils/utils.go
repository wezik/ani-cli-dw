package utils

import (
	"ani-cli-dw/logger"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type InputWrapper struct {
        Title       string
        EntryId     int
        RangeStart  int
        RangeEnd    int
        Quality     int
        ThreadCount int
}


func ReadInput() InputWrapper {
        input := InputWrapper{}
        scanner := bufio.NewScanner(os.Stdin)

        scan := func (label string, scanner *bufio.Scanner) string {
                fmt.Print(label)
                scanner.Scan()
                return scanner.Text()
        }

        input.Title = scan("Enter the anime name: ", scanner)
        entryIdString := scan("Enter the anime entry id: ", scanner)
        rangeStartString := scan("Enter the range start: ", scanner)
        rangeEndString := scan("Enter the range end: ", scanner)
        qualityString := scan("Enter the quality (ex. 480/720/1080): ", scanner)
        threadCountString := scan("How many simultaneous downloads do you want to hold: ", scanner)

        atoiWithHandling := func(value string) int {
                num, err := strconv.Atoi(value)
                if err != nil {
                        message := fmt.Sprintf("Invalid input for %s: %v", value, err)
                        logger.Error.Fatal(message)
                }
                return num
        }

        input.EntryId = atoiWithHandling(entryIdString)
        input.RangeStart = atoiWithHandling(rangeStartString)
        input.RangeEnd = atoiWithHandling(rangeEndString)
        input.Quality = atoiWithHandling(qualityString)
        input.ThreadCount = atoiWithHandling(threadCountString)

        return input
}

func SetupDirectory(input InputWrapper) string {
        entryId := strconv.Itoa(input.EntryId)
        dir := strings.Join([]string{input.Title, entryId}, "-")
        path := filepath.Join(".", dir)
        
        logger.Info.Printf("Looking for \"%s\" directory", path)
        if _, err := os.Stat(path); os.IsNotExist(err) {
                logger.Info.Printf("\"%s\" not found, creating it", path)
                err := os.MkdirAll(path, 0755)
                if err != nil {
                        logger.Error.Fatalf("Error creating directory: %v\n", err)
                }
                logger.Info.Printf("\"%s\" created", path)
        }

        return path
}
