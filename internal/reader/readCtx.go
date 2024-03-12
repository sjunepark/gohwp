package reader

import (
	"context"
	"github.com/sjunepark/gohwp/internal/reader/model"
)

const versionKey = "hwpVersion"

func setVersion(ctx context.Context, version model.HWPVersion) context.Context {
	return context.WithValue(ctx, versionKey, version)
}

func getVersion(ctx context.Context) (model.HWPVersion, bool) {
	v, ok := ctx.Value(versionKey).(model.HWPVersion)
	return v, ok
}
