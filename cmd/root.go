package cmd

import (
	"fmt"
	"os"

	"github.com/fabiotavarespr/Planilha-MTG/export"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "planilha",
	Short: "Gerador de planilha de edições de Magic: The Gathering",
	Long:  "Geradora uma planilha de edições de Magic: The Gathering que você solicita",
	Run: func(cmd *cobra.Command, args []string) {
		export.ApresentaListagem("ELD")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
