package main

import "log"

type DLWrapper struct {
        title       string
        entryId     int
        episode     int
        quality     int
        downloadDir string
}

type InputWrapper struct {
        title       string
        entryId     int
        rangeStart  int
        rangeEnd    int
        quality     int
        threadCount int
}

var (
        InfoLogger *log.Logger
        ErrorLogger *log.Logger
        DebugLogger *log.Logger
        WarningLogger *log.Logger
)
