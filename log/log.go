package log

import "github.com/charmbracelet/log"

func GetOptions(level string) log.Options {
	opts := log.Options{
		ReportTimestamp: false,
		Prefix:          "FS",
	}
	switch level {
	case "debug":
		opts.Level = log.DebugLevel
	case "info":
		opts.Level = log.InfoLevel
	default:
		opts.Level = log.WarnLevel
	}
	return opts
}
