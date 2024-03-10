package parser

import (
	"fmt"
	"github.com/richardlehane/mscfb"
	"github.com/sjunepark/gohwp/internal/types"
	"os"
)

var signature = "Hwp Document File"
var supportedVersion = types.HWPVersion{Major: 5, Minor: 1}

func Parse(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	reader, err := mscfb.New(file)
	if err != nil {
		return err
	}

	doc := &types.HWPDocument{}

	for entry, err := reader.Next(); err == nil; entry, err = reader.Next() {
		switch entry.Name {
		case "FileHeader":
			headerData := make([]byte, entry.Size)
			_, err := reader.Read(headerData)
			if err != nil {
				return err
			}

			header, err := types.NewHWPHeader(headerData)
			if err != nil {
				return err
			}

			doc.Header = header
		case "DocInfo":
			docInfoData := make([]byte, entry.Size)
			_, err := reader.Read(docInfoData)
			if err != nil {
				return err
			}

			// Raw deflate al
			deCompressedData, err := DecompressDeflate(docInfoData)
			if err != nil {
				return err
			}

			docInfoParser, err := NewDocInfoParser(deCompressedData)
			if err != nil {
				return err
			}

			docInfo, err := docInfoParser.Parse()
			if err != nil {
				return err
			}

			//	todo: remove
			fmt.Println(docInfo)
		}

	}

	return nil
}
