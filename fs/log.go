package fs

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"

	alog "github.com/avoceteditors/avocet/log"
)

var lgr *log.Logger

func init() {

	// Initialize Defaults
	viper.SetDefault("avocet.buffers.rune", 20)
	viper.SetDefault("avocet.log.fs", "debug")
}

func initLogger() {

	// Configure Logger
	switch viper.GetString("avocet.log.fs") {
	lgr = log.NewWithOptions(
		os.Stderr, 
		alog.GetOptions(viper.GetString("avocet.log.fs")))

}
