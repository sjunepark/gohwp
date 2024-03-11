package models

type CharShape struct {
	FontId           SupportedLocaleOptions
	FontScale        SupportedLocaleOptions
	FontSpacing      SupportedLocaleOptions
	FontRatio        SupportedLocaleOptions
	FontLocation     SupportedLocaleOptions
	FontBaseSize     float64
	Attr             int
	Shadow           ColorRef
	Shadow2          ColorRef
	Color            ColorRef
	UnderLineColor   ColorRef
	ShadeColor       ColorRef
	ShadowColor      ColorRef
	FontBackgroundId *int
	StrikeColor      *ColorRef
}

func NewCharShape(
	fontId SupportedLocaleOptions,
	fontScale SupportedLocaleOptions,
	fontSpacing SupportedLocaleOptions,
	fontRatio SupportedLocaleOptions,
	fontLocation SupportedLocaleOptions,
	fontBaseSize float64,
	attr int,
	shadow int,
	shadow2 int,
	color int,
	underLineColor int,
	shadeColor int,
	shadowColor int,
) CharShape {
	return CharShape{
		FontId:           fontId,
		FontScale:        fontScale,
		FontSpacing:      fontSpacing,
		FontRatio:        fontRatio,
		FontLocation:     fontLocation,
		FontBaseSize:     fontBaseSize / 100,
		Attr:             attr,
		Shadow:           getRGB(shadow),
		Shadow2:          getRGB(shadow2),
		Color:            getRGB(color),
		UnderLineColor:   getRGB(underLineColor),
		ShadeColor:       getRGB(shadeColor),
		ShadowColor:      getRGB(shadowColor),
		FontBackgroundId: nil,
		StrikeColor:      nil,
	}
}

type SupportedLocaleOptions = [7]int
