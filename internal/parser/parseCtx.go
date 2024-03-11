package parser

import (
	"context"
	"github.com/sjunepark/gohwp/internal/models"
)

const versionKey = "hwpVersion"

func WithVersion(ctx context.Context, version models.HWPVersion) context.Context {
	return context.WithValue(ctx, versionKey, version)
}

func VersionFromContext(ctx context.Context) (models.HWPVersion, bool) {
	version, ok := ctx.Value(versionKey).(models.HWPVersion)
	return version, ok
}
