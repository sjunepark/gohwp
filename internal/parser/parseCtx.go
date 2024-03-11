package parser

import (
	"context"
	"github.com/sjunepark/gohwp/internal/models"
)

const versionKey = "hwpVersion"

func setVersion(ctx context.Context, version models.HWPVersion) context.Context {
	return context.WithValue(ctx, versionKey, version)
}

func getVersion(ctx context.Context) (models.HWPVersion, bool) {
	v, ok := ctx.Value(versionKey).(models.HWPVersion)
	return v, ok
}
