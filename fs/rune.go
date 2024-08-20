package fs

import (
	"github.com/spf13/viper"
)

func runeHandler(rs []rune) chan rune {
	ch := make(chan rune, viper.GetInt("avocet.buffers.rune")+1)
	go func() {
		for _, r := range rs {
			ch <- r
		}
		close(ch)
	}()
	return ch
}
