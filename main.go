package main

import (
	"log"
	"os"
)

func main() {
        inputWrapper := readInput()
        path := setupDirectory(inputWrapper)
        Download(inputWrapper, path)
}

func init() {
        InfoLogger = log.New(os.Stderr, "INFO: ", log.Ldate | log.Ltime)
        WarningLogger = log.New(os.Stderr, "WARN: ", log.Ldate | log.Ltime)
        ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate | log.Ltime)
        DebugLogger = log.New(os.Stderr, "DEBUG: ", log.Ldate | log.Ltime)
}
