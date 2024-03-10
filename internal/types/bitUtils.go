package types

func getBitValue(value int, start int, end int) uint8 {
	return uint8((value >> start) & ((1 << (end - start + 1)) - 1))
}

func getRGB(colorRef int) ColorRef {
	return ColorRef{
		getBitValue(colorRef, 0, 7),
		getBitValue(colorRef, 8, 15),
		getBitValue(colorRef, 16, 23),
	}
}

func getFlag(bits int, position int) bool {
	mask := 1 << position
	return (bits & mask) == mask
}
