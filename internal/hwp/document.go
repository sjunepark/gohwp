package hwp

import (
	"github.com/sjunepark/gohwp/internal/docInfo"
	"github.com/sjunepark/gohwp/internal/models"
	"github.com/sjunepark/gohwp/internal/section"
)

type HWPDocument struct {
	Header   *models.HWPHeader  `validate:"required"`
	DocInfo  *docInfo.DocInfo   `validate:"required"`
	BodyText []*section.Section `validate:"required"`
}
