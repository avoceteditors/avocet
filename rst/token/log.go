package token

import (
	"os"

	alog "github.com/avoceteditors/avocet/log"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("avocet.log.rst.token", "debug")
}

var lgr *log.Logger

func initLogger() {
	lgr = log.NewWithOptions(
		os.Stderr,
		alog.GetOptions(viper.GetString("avocet.log.rst.token")))

}
