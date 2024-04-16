package main

import (
        "fmt"
        "os/exec"
        "sort"
        "strings"
        "sync"
)

func Download(input InputWrapper, destination string) {
        InfoLogger.Println("Setting up downloads")

        var wg sync.WaitGroup
        wg.Add(input.rangeEnd - input.rangeStart + 1)
        semaphore := make(chan struct{}, input.threadCount)

        InfoLogger.Println("Starting download routines")

        var failedDownloads []DLWrapper

        for episode := input.rangeStart; episode <= input.rangeEnd; episode++ {

                semaphore <- struct{}{}

                wrapper := DLWrapper{
                        input.title,
                        input.entryId,
                        episode,
                        input.quality,
                        destination, 
                }

                go func() {
                        err := downloadEpisode(wrapper, &wg, semaphore)
                        if err != nil {
                                failedDownloads = append(failedDownloads, wrapper)
                        }
                }()
        }

        wg.Wait()

        var episodes []int
        for _, dl := range failedDownloads {
                episodes = append(episodes, dl.episode)
        }

        sort.Ints(episodes)

        if len(episodes) > 0 {
                WarningLogger.Printf("Episodes %v failed to download!", episodes)
        }

        InfoLogger.Println("Downloading finished")
}

func downloadEpisode(wrapper DLWrapper, wg *sync.WaitGroup, semaphore chan struct{}) error {
        defer wg.Done()
        defer func() { <-semaphore }()
        prefix := fmt.Sprintf("\"%s\" episode %d", wrapper.title, wrapper.episode)

        args := []string{
                "ani-cli",
                wrapper.title,
                fmt.Sprintf("-S %d", wrapper.entryId),
                fmt.Sprintf("-e %d", wrapper.episode),
                fmt.Sprintf("-q %dp", wrapper.quality),
                "-d",
        }

        finalCmd := strings.Join(args, " ")

        // DebugLogger.Println("Executing:", finalCmd)

        cmd := exec.Command("cmd.exe", "/C", finalCmd) //For now windows specific

        cmd.Dir = wrapper.downloadDir

        InfoLogger.Printf("%s download start", prefix)
        err := cmd.Run()
        if err != nil {
                WarningLogger.Printf("%s download fail! :%v", prefix, err)
                return err
        }

        InfoLogger.Printf("%s downloaded success", prefix)
        return nil
}
