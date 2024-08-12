package main

import (
	"github.com/avoceteditors/avocet"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmd = &cobra.Command{
	Use:   "avocet",
	Short: "Document Processor",
}

var lex = &cobra.Command{
	Use:   "lex",
	Short: "String lexer (for internal/debugging use only)",
	Run: func(_ *cobra.Command, args []string) {
		avocet.RunLexer(args)
	},
}

func init() {
	viper.SetDefault("avocet.default_lexer", "rst")

	lex.PersistentFlags().StringP(
		"lexer", "L",
		viper.GetString("avocet.default_lexer"),
		"Specify the lexer you want to use")
	viper.BindPFlag("avocet.default_lexer", lex.PersistentFlags().Lookup("lexer"))

}

func main() {
	cmd.AddCommand(lex)
	cmd.Execute()
}
