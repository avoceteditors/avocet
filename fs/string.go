package fs

// ReadString takes a string and retunrs a rune channel. Internally,
// it also starts a goroutine that passes the runes of the string
// into the channel.
func ReadString(s string) chan rune {
	return runeHandler([]rune(s))
}
