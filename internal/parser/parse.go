package parser

import (
	"fmt"
	"github.com/richardlehane/mscfb"
	"github.com/sjunepark/gohwp/internal/docInfo"
	"github.com/sjunepark/gohwp/internal/hwp"
	"github.com/sjunepark/gohwp/internal/models"
	"os"
	"strings"
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

	documentData, err := getDocumentData(reader)

	h, err := getHeader(documentData.header)
	if err != nil {
		return err
	}
	doc.Header = h

	di, err := getDocInfo(documentData.docInfo)
	if err != nil {
		return err
	}
	doc.DocInfo = di

	s, err := getSections(documentData.bodyText)
	if err != nil {
		return err
	}
	doc.BodyText = s

	return nil
}

type documentData struct {
	header   []byte
	docInfo  []byte
	bodyText []sectionData
}

type sectionData []byte

func getDocumentData(reader *mscfb.Reader) (*documentData, error) {
	d := &documentData{}

	for entry, err := reader.Next(); err == nil; entry, err = reader.Next() {
		entryName := entry.Name
		switch {
		case entryName == "FileHeader":
			data, err := getData(reader, entry.Size)
			if err != nil {
				return nil, err
			}
			d.header = data
		case entryName == "DocInfo":
			data, err := getData(reader, entry.Size)
			if err != nil {
				return nil, err
			}
			d.docInfo = data
		// Starts with Section
		case strings.HasPrefix(entryName, "Section"):
			data, err := getData(reader, entry.Size)
			if err != nil {
				return nil, err
			}
			d.bodyText = append(d.bodyText, data)
		}
	}
	return d, nil
}

func getData(reader *mscfb.Reader, size int64) ([]byte, error) {
	data := make([]byte, size)
	_, err := reader.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getHeader(data []byte) (*models.HWPHeader, error) {
	header, err := models.NewHWPHeader(data)
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

func getDocInfo(data []byte) (*docInfo.DocInfo, error) {
	deCompressedData, err := DecompressDeflate(data)
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

func getSections(data []sectionData) ([]*models.Section, error) {
	sections := make([]*models.Section, len(data))
	for _, sectionData := range data {
		// todo: parse sectionData
		// todo: append section to sections
		fmt.Println(sectionData)
	}
	return sections, nil
}
