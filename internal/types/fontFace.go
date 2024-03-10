package types

import "strings"

type FontFace struct {
	Name        string
	Alternative string
	Default     string
	Panose      *Panose
}

func (f FontFace) GetFontFamily() string {
	var result []string
	//Initialize result, not append
	result = []string{f.Name}
	if f.Alternative != "" {
		result = append(result, f.Alternative)
	}
	if f.Default != "" {
		result = append(result, f.Default)
	}
	if f.Panose != nil {
		panoseFontFamily := f.Panose.GetFontFamily()
		result = append(result, panoseFontFamily)
	}
	return strings.Join(result, ",")
}
