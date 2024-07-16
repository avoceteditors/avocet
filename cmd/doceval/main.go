package main

import (
	"github.com/avoceteditors/avocet"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "rev",
	Short: "A Highlighter of Common Mistakes",
	Run: func(cmd *cobra.Command, args []string) {
		avocet.DocEval(args)
	},
}

func main() {
	cmd.Execute()

}
