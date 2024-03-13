package reader

import (
	"context"
	"fmt"
	"github.com/richardlehane/mscfb"
	"github.com/sjunepark/hwp/internal/reader/model"
	"os"
	"strings"
)

func Read(filePath string) (doc *model.Document, encrypted bool, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return &model.Document{}, false, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	reader, err := mscfb.New(file)
	if err != nil {
		return &model.Document{}, false, err
	}

	doc = &model.Document{}

	documentData, err := getDocumentData(reader)
	if err != nil {
		return &model.Document{}, false, err
	}

	header, err := getHeader(documentData.header)
	if err != nil {
		return &model.Document{}, false, err
	}
	doc.Header = header
	ctx := context.Background()
	ctx = setVersion(ctx, header.Version)

	// todo: test if this works for encrypted documents
	// Early return when document is ed
	if header.Attributes1.Encrypted {
		return &model.Document{}, false, err
	}

	docInfo, err := getDocInfo(documentData.docInfo)
	if err != nil {
		return &model.Document{}, false, err
	}
	doc.DocInfo = docInfo

	sections, err := getSections(documentData.bodyText, ctx)
	if err != nil {
		return &model.Document{}, false, err
	}
	doc.BodyText = sections

	return doc, false, err
}

type documentData struct {
	header   []byte
	docInfo  []byte
	bodyText []sectionData
}

type sectionData []byte

func getDocumentData(reader *mscfb.Reader) (*documentData, error) {
	dd := &documentData{}

	for entry, err := reader.Next(); err == nil; entry, err = reader.Next() {
		entryName := entry.Name
		switch {
		case entryName == "FileHeader":
			data, err := getData(reader, entry.Size)
			if err != nil {
				return nil, err
			}
			dd.header = data
		case entryName == "DocInfo":
			data, err := getData(reader, entry.Size)
			if err != nil {
				return nil, err
			}
			dd.docInfo = data
		// Starts with Section
		case strings.HasPrefix(entryName, "Section"):
			data, err := getData(reader, entry.Size)
			if err != nil {
				return nil, err
			}
			dd.bodyText = append(dd.bodyText, data)
		}
	}
	return dd, nil
}

func getData(reader *mscfb.Reader, size int64) ([]byte, error) {
	data := make([]byte, size)
	_, err := reader.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getHeader(data []byte) (*model.HWPHeader, error) {
	header, err := model.NewHWPHeader(data)
	if err != nil {
		return nil, err
	}

	supportedSignature := "HWP Document File"
	if header.Signature != supportedSignature {
		return nil, fmt.Errorf("unsupported signature: %s", header.Signature)
	}

	supportedVersion := model.HWPVersion{Major: 5}
	if !header.Version.IsCompatible(supportedVersion) {
		return nil, fmt.Errorf("unsupported version: %s", header.Version)
	}
	return header, nil
}

func getDocInfo(data []byte) (*model.DocInfo, error) {
	deCompressedData, err := DecompressDeflate(data)
	if err != nil {
		return nil, err
	}

	docInfoReader, err := NewDocInfoReader(deCompressedData)
	if err != nil {
		return nil, err
	}

	di, err := docInfoReader.Read()
	if err != nil {
		return nil, err
	}
	return di, nil
}

func getSections(data []sectionData, ctx context.Context) ([]*model.Section, error) {
	sections := make([]*model.Section, 0, len(data))
	for _, sectionData := range data {
		deCompressedData, err := DecompressDeflate(sectionData)
		if err != nil {
			return nil, err
		}

		sectionReader, err := NewSectionReader(deCompressedData)
		if err != nil {
			return nil, err
		}

		s, err := sectionReader.Read(ctx)
		if err != nil {
			return nil, err
		}

		sections = append(sections, s)
	}
	return sections, nil
}
