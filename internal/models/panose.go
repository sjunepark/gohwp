package models

type Panose struct {
	Family          int
	SerifStyle      int
	Weight          int
	Proportion      int
	Contrast        int
	StrokeVariation int
	ArmStyle        int
	LetterForm      int
	Midline         int
	XHeight         int
}

func (p Panose) GetFontFamily() string {
	if p.Family == 3 {
		return "cursive"
	}

	if p.Family == 2 {
		if p.SerifStyle > 1 && p.SerifStyle < 11 {
			return "sans"
		}

		if p.SerifStyle > 10 && p.SerifStyle < 14 {
			return "sans-serf"
		}
	}

	return ""
}
