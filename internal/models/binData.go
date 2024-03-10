package models

type binDataType int

const (
	binDataTypeLink binDataType = iota
	binDataTypeEmbedding
	binDataTypeStorage
)

type binDataCompress int

const (
	binDataCompressDefault binDataCompress = iota
	binDataCompressCompress
	binDataCompressNotCompress
)

type binDataStatus int

const (
	binDataStatusInitial binDataStatus = iota
	binDataStatusSuccess
	binDataStatusError
	binDataStatusIgnore
)

type binProperties struct {
	binDataType binDataType
	compress    binDataCompress
	status      binDataStatus
}

type BinData struct {
	properties binProperties
	extension  string
	payload    []byte
}
