package main

import "github.com/sjunepark/gohwp/internal/parser"

func main() {
	err := parser.Parse("data/example.hwp")
	if err != nil {
		return
	}
}
