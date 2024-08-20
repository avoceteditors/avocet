package fs

func ReadString(s string) chan rune {
	return runeHandler([]rune(s))
}
