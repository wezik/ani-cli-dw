package main

import (
	"ani-cli-dw/downloader"
	"ani-cli-dw/logger"
	"ani-cli-dw/utils"
	"log"
	"os"
)

func main() {
        input := utils.ReadInput()
        path := utils.SetupDirectory(input)
        downloader.Download(input, path)
}

func init() {
        logger.Info = log.New(os.Stderr, "INFO: ", log.Ldate | log.Ltime)
        logger.Warning = log.New(os.Stderr, "WARN: ", log.Ldate | log.Ltime)
        logger.Error = log.New(os.Stderr, "ERROR: ", log.Ldate | log.Ltime)
        logger.Debug = log.New(os.Stderr, "DEBUG: ", log.Ldate | log.Ltime)
}
