package parser

import (
	"fmt"
	"github.com/richardlehane/mscfb"
	"github.com/sjunepark/gohwp/internal/docInfo"
	"github.com/sjunepark/gohwp/internal/hwp"
	"github.com/sjunepark/gohwp/internal/models"
	"os"
)

var signature = "Hwp Document File"
var supportedVersion = models.HWPVersion{Major: 5, Minor: 1}

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

	doc := &hwp.HWPDocument{}

	for entry, err := reader.Next(); err == nil; entry, err = reader.Next() {
		switch entry.Name {
		case "FileHeader":
			headerData := make([]byte, entry.Size)
			_, err := reader.Read(headerData)
			if err != nil {
				return err
			}

			header, err := models.NewHWPHeader(headerData)
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

			docInfoParser, err := docInfo.NewParser(deCompressedData)
			if err != nil {
				return err
			}

			di, err := docInfoParser.Parse()
			if err != nil {
				return err
			}

			//	todo: remove
			fmt.Println(di)
		}

	}

	return nil
}
