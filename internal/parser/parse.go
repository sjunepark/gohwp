package parser

import (
	"fmt"
	"github.com/richardlehane/mscfb"
	"github.com/sjunepark/gohwp/internal/docInfo"
	"github.com/sjunepark/gohwp/internal/hwp"
	"github.com/sjunepark/gohwp/internal/models"
	"os"
)

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

	h, err := getHeader(reader)
	if err != nil {
		return err
	}
	doc.Header = h

	di, err := getDocInfo(reader)
	if err != nil {
		return err
	}
	doc.DocInfo = di

	return nil
}

func getHeader(reader *mscfb.Reader) (*models.HWPHeader, error) {
	for entry, err := reader.Next(); err == nil; entry, err = reader.Next() {
		if entry.Name == "FileHeader" {
			headerData := make([]byte, entry.Size)
			_, err := reader.Read(headerData)
			if err != nil {
				return nil, err
			}

			header, err := models.NewHWPHeader(headerData)
			if err != nil {
				return nil, err
			}

			supportedSignature := "HWP Document File"
			if header.Signature != supportedSignature {
				return nil, fmt.Errorf("unsupported signature: %s", header.Signature)
			}

			supportedVersion := models.HWPVersion{Major: 5, Minor: 0}
			if header.Version.IsCompatible(supportedVersion) == false {
				return nil, fmt.Errorf("unsupported version: %s", header.Version)
			}
			return header, nil
		}
	}
	return nil, fmt.Errorf("FileHeader not found")
}

func getDocInfo(reader *mscfb.Reader) (*docInfo.DocInfo, error) {
	for entry, err := reader.Next(); err == nil; entry, err = reader.Next() {
		if entry.Name == "DocInfo" {
			docInfoData := make([]byte, entry.Size)
			_, err := reader.Read(docInfoData)
			if err != nil {
				return nil, err
			}

			// Raw deflate algorithm
			deCompressedData, err := DecompressDeflate(docInfoData)
			if err != nil {
				return nil, err
			}

			docInfoParser, err := docInfo.NewParser(deCompressedData)
			if err != nil {
				return nil, err
			}

			di, err := docInfoParser.Parse()
			if err != nil {
				return nil, err
			}
			return di, nil
		}
	}
	return nil, fmt.Errorf("DocInfo not found")
}
