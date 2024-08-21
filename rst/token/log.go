package token

import (

	"github.com/spf13/viper"
	"github.com/charmbracelet/log"
	alog "github.com/avoceteditors/avocet/log"
)

func init() {
	viper.SetDefault("avocet.log.rst.token", "debug")
}

var lgr *log.Logger

func initLogger(){
	lgr = log.NewWithOptions(
		os.Stderr,
		alog.GetOptions(viper.GetString("avocet.log.rst.token"))

}