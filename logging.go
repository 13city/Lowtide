package main

import (
    "log"
    "os"
    "path/filepath"

    "gopkg.in/natefinch/lumberjack.v2"
)

func SetupLogger() *log.Logger {
    logDirPath := "./logging"
    err := os.MkdirAll(logDirPath, os.ModePerm)
    if err != nil {
        log.Fatalf("Failed to create log directory: %v", err)
    }

    logFilePath := filepath.Join(logDirPath, "Lowtide.log")
    logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
    logger.SetOutput(&lumberjack.Logger{
        Filename:   logFilePath,
        MaxSize:    5, // in MB
        MaxBackups: 3,
        MaxAge:     28, //days
        Compress:   true,
    })

    return logger
}