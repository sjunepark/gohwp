package hwp

import (
	"github.com/sjunepark/gohwp/internal/docInfo"
	"github.com/sjunepark/gohwp/internal/models"
)

type HWPDocument struct {
	Header   *models.HWPHeader `validate:"required"`
	DocInfo  *docInfo.DocInfo  `validate:"required"`
	BodyText []*models.Section `validate:"required"`
}
