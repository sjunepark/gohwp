package model

import (
	"bytes"
	"encoding/binary"
	"github.com/sjunepark/gohwp/internal/reader/constant"
	"github.com/sjunepark/gohwp/internal/validator"
)

type ByteReader struct {
	Data   []byte
	offset int
}

func (br *ByteReader) ReadRecord() (*Record, error) {
	header, err := br.readRecordHeader()
	if err != nil {
		return nil, err
	}

	data, err := br.ReadBytes(int(header.Size))
	if err != nil {
		return nil, err
	}

	record := &Record{RecordHeader: *header, Payload: data}
	return record, nil
}

func (br *ByteReader) readRecordHeader() (*RecordHeader, error) {
	dWord, err := br.readDword()
	if err != nil {
		return nil, err
	}
	header := &RecordHeader{
		TagID: constant.TagID(dWord & 0x3FF),
		Level: uint32((dWord >> 10) & 0x3FF),
		Size:  uint32((dWord >> 20) & 0xFFF),
	}

	err = validator.ValidateStruct(header)
	if err != nil {
		return nil, err
	}

	// Size : 데이터 영역의 길이를 바이트 단위로 나타낸다. 12개의 비트가 모두 1일 때는 데이터 영역의
	// 길이가 4095 바이트 이상인 경우로, 이때는 레코드 헤더에 연이어 길이를 나타내는 DWORD 가 추가된다.
	if header.Size == 0xFFF {
		additionalSize, err := br.readDword()
		if err != nil {
			return nil, err
		}
		header.Size = uint32(additionalSize)
	}

	return header, nil
}

func (br *ByteReader) ReadUint16() (uint16, error) {
	if br.offset+2 > len(br.Data) {
		return 0, &OutOfBoundsError{Requested: br.offset + 2, Max: len(br.Data)}
	}

	result := binary.LittleEndian.Uint16(br.Data[br.offset : br.offset+2])
	br.offset += 2
	return result, nil
}

func (br *ByteReader) ReadUint32() (uint32, error) {
	if br.offset+4 > len(br.Data) {
		return 0, &OutOfBoundsError{Requested: br.offset + 4, Max: len(br.Data)}
	}

	result := binary.LittleEndian.Uint32(br.Data[br.offset : br.offset+4])
	br.offset += 4
	return result, nil
}

func (br *ByteReader) ReadBytes(n int) ([]byte, error) {
	if br.offset+n > len(br.Data) {
		return nil, &OutOfBoundsError{Requested: br.offset + n, Max: len(br.Data)}
	}

	result := br.Data[br.offset : br.offset+n]
	br.offset += n
	return result, nil
}

func (br *ByteReader) readDword() (Dword, error) {
	if br.offset+4 > len(br.Data) {
		return 0, &OutOfBoundsError{Requested: br.offset + 4, Max: len(br.Data)}
	}

	result := binary.LittleEndian.Uint32(br.Data[br.offset : br.offset+4])
	br.offset += 4
	dword := Dword(result)
	return dword, nil
}

func (br *ByteReader) ReadStruct(data interface{}) error {
	size := binary.Size(data)
	if br.offset+size > len(br.Data) {
		return &OutOfBoundsError{Requested: br.offset + size, Max: len(br.Data)}
	}

	reader := bytes.NewReader(br.Data[br.offset : br.offset+size])
	err := binary.Read(reader, binary.LittleEndian, data)
	if err != nil {
		return err
	}

	br.offset += size
	return nil
}

func (br *ByteReader) Skip(n int) error {
	if br.offset+n > len(br.Data) {
		return &OutOfBoundsError{Requested: br.offset + n, Max: len(br.Data)}
	}
	br.offset += n
	return nil
}

func (br *ByteReader) IsEOF() bool {
	return br.offset >= len(br.Data)
}
