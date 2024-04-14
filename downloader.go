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

func downloadEpisode(dlWrapper DLWrapper, wg *sync.WaitGroup) {
        defer wg.Done()

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

func main() {
        scanner := bufio.NewScanner(os.Stdin)

        fmt.Print("Enter the anime name: ")
        scanner.Scan()
        anime := scanner.Text()

        fmt.Print("Enter the anime entry id: ")
        scanner.Scan()
        animeEntryId, err := strconv.Atoi(scanner.Text())
        if err != nil {
                fmt.Printf("Invalid anime entry id: %v\n", err)
                return
        }

        fmt.Print("Enter the range start: ")
        scanner.Scan()
        rangeStart, err := strconv.Atoi(scanner.Text())
        if err != nil {
                fmt.Printf("Invalid range start: %v\n", err)
                return
        }

        fmt.Print("Enter the range end: ")
        scanner.Scan()
        rangeEnd, err := strconv.Atoi(scanner.Text())
        if err != nil {
                fmt.Printf("Invalid range end: %v\n", err)
                return
        }

        fmt.Print("Enter the quality (ex. 480/720/1080): ")
        scanner.Scan()
        quality, err := strconv.Atoi(scanner.Text())
        if err != nil {
                fmt.Printf("Invalid quality: %v\n", err)
                return
        }
        
        animeDir := filepath.Join(".", anime)

        if _, err := os.Stat(animeDir); os.IsNotExist(err) {
                err := os.MkdirAll(animeDir, 0755)
                if err != nil {
                        fmt.Printf("Error creating directory: %v\n", err)
                        return
                }
        }

        var wg sync.WaitGroup

        wg.Add(rangeEnd - rangeStart + 1)

        counter := 0

        for i := rangeStart; i <= rangeEnd; i++ {
                dlWrapper := DLWrapper{anime, animeEntryId, i, quality, animeDir}
                go downloadEpisode(dlWrapper, &wg)
                counter++
                if (counter >= 10) {
                        wg.Wait()
                }
        }

        wg.Wait()

        fmt.Println("All episodes downloaded")
}

