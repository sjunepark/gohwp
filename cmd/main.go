package main

import (
	"fmt"
	"github.com/sjunepark/gohwp/internal/parser"
)

func main() {
	err := parser.Parse("data/example.hwp")
	if err != nil {
		fmt.Println(err)
	}
}
