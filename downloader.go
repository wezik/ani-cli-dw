package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
)

type DLWrapper struct {
        anime string
        animeEntryId int
        episode int
        quality int
        animeDir string
}

func downloadEpisode(dlWrapper DLWrapper, wg *sync.WaitGroup, semaphore chan struct{}) {
        defer wg.Done()
        defer func() { <-semaphore }()

        animeEntryIdArg := fmt.Sprintf("-S %d", dlWrapper.animeEntryId)
        rangeArg := fmt.Sprintf("-e %d", dlWrapper.episode)
        qualityArg := fmt.Sprintf("-q %dp", dlWrapper.quality)

        formatedCmd := fmt.Sprintf("ani-cli %s %s %s %s -d", dlWrapper.anime, animeEntryIdArg, rangeArg, qualityArg)

        fmt.Println("Executing:", formatedCmd)

        cmd := exec.Command("cmd.exe", "/C", formatedCmd)

        cmd.Dir = dlWrapper.animeDir

        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        err := cmd.Run()
        if err != nil {
                fmt.Printf("Error downloading episode %d: %v\n", dlWrapper.episode, err)
                return
        }

        fmt.Printf("Episode %d downloaded successfully\n", dlWrapper.episode)
}

type InputWrapper struct {
        anime string
        animeEntryId int
        rangeStart int
        rangeEnd int
        quality int
        threadCount int
}

func main() {
        inputWrapper := readInput()
        animeDir := filepath.Join(".", inputWrapper.anime, strconv.Itoa(inputWrapper.animeEntryId))

        if _, err := os.Stat(animeDir); os.IsNotExist(err) {
                err := os.MkdirAll(animeDir, 0755)
                if err != nil {
                        fmt.Printf("Error creating directory: %v\n", err)
                        return
                }
        }

        var wg sync.WaitGroup
        wg.Add(inputWrapper.rangeEnd - inputWrapper.rangeStart + 1)
        semaphore := make(chan struct{}, inputWrapper.threadCount)

        for episode := inputWrapper.rangeStart; episode <= inputWrapper.rangeEnd; episode++ {

                semaphore <- struct{}{}

                dlWrapper := DLWrapper{
                        inputWrapper.anime,
                        inputWrapper.animeEntryId,
                        episode,
                        inputWrapper.quality,
                        animeDir,
                }

                go downloadEpisode(dlWrapper, &wg, semaphore)
        }

        wg.Wait()

        fmt.Println("All episodes downloaded")
}

func readInput() InputWrapper {
        wrapper := InputWrapper{}
        scanner := bufio.NewScanner(os.Stdin)

        scan := func (label string, scanner *bufio.Scanner) string {
                fmt.Print(label)
                scanner.Scan()
                return scanner.Text()
        }

        wrapper.anime = scan("Enter the anime name: ", scanner)
        animeEntryIdString := scan("Enter the anime entry id: ", scanner)
        rangeStartString := scan("Enter the range start: ", scanner)
        rangeEndString := scan("Enter the range end: ", scanner)
        qualityString := scan("Enter the quality (ex. 480/720/1080): ", scanner)
        threadCountString := scan("How many simultaneous downloads do you want to hold: ", scanner)

        atoiWithHandling := func(value string) int {
                num, err := strconv.Atoi(value)
                if err != nil {
                        message := fmt.Sprintf("Invalid input for %s: %v", value, err)
                        fmt.Println(message)
                        os.Exit(1)
                }
                return num
        }

        wrapper.animeEntryId = atoiWithHandling(animeEntryIdString)
        wrapper.rangeStart = atoiWithHandling(rangeStartString)
        wrapper.rangeEnd = atoiWithHandling(rangeEndString)
        wrapper.quality = atoiWithHandling(qualityString)
        wrapper.threadCount = atoiWithHandling(threadCountString)

        return wrapper
}


