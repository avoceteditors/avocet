package fs

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

var lgr *log.Logger

func init() {

	// Initialize Defaults
	viper.SetDefault("avocet.buffers.rune", 20)

	// Configure Logger
	opts := log.Options{
		ReportTimestamp: false,
		Prefix:          "FS",
	}
	switch viper.GetString("avocet.log.fs") {
	case "debug":
		opts.Level = log.DebugLevel
	case "info":
		opts.Level = log.InfoLevel
	default:
		opts.Level = log.WarnLevel
	}
	lgr = log.NewWithOptions(os.Stderr, opts)

}
