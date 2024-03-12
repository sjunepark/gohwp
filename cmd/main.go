package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sjunepark/gohwp/internal/parser"
	"github.com/sjunepark/gohwp/internal/reader"
	"log/slog"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {
	initSlog()

	raw, encrypted, err := reader.Read("data/example.hwp")
	if err != nil {
		fmt.Println(err)
	}
	if encrypted {
		fmt.Println("Document is encrypted")
	}

	doc := parser.Parse(raw)
	fmt.Println(doc)
}

func initSlog() {
	logLevel := new(slog.LevelVar)
	logLevel.Set(getLogLevel())
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)
}

func getLogLevel() slog.Level {
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	}
	panic("invalid log level")
}
