package downloader

import (
	"ani-cli-dw/logger"
	"ani-cli-dw/utils"
	"fmt"
	"os/exec"
	"sort"
	"strings"
	"sync"
)

type DLWrapper struct {
        title       string
        entryId     int
        episode     int
        quality     int
        downloadDir string
}

func Download(input utils.InputWrapper, destination string) {
        logger.Info.Println("Setting up downloads")

        var wg sync.WaitGroup
        wg.Add(input.RangeEnd - input.RangeStart + 1)
        semaphore := make(chan struct{}, input.ThreadCount)

        logger.Info.Println("Starting download routines")

        var failedDownloads []DLWrapper

        for episode := input.RangeStart; episode <= input.RangeEnd; episode++ {

                semaphore <- struct{}{}

                wrapper := DLWrapper{
                        input.Title,
                        input.EntryId,
                        episode,
                        input.Quality,
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
                logger.Warning.Printf("Episodes %v failed to download!", episodes)
        }

        logger.Info.Println("Downloading finished")
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

        logger.Info.Printf("%s download start", prefix)
        err := cmd.Run()
        if err != nil {
                logger.Warning.Printf("%s download fail! :%v", prefix, err)
                return err
        }

        logger.Info.Printf("%s downloaded success", prefix)
        return nil
}
