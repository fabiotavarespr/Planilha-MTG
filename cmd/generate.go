package cmd

import (
	"fmt"
	"github.com/fabiotavarespr/Planilha-MTG/export"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Generate the spreadsheet",
	Long: `Generate the spreadsheet according to the code of the informed edition
	`,
	Run: func(cmd *cobra.Command, args []string) {
		runGenerate(args)
	},
}

func runGenerate(args []string) {
	if len(args) == 1 {
		code := args[0]
		fmt.Println("Starting spreadsheet generation - " + code)
		export.ApresentaListagem(code)
		fmt.Println("Spreadsheet generation finished - " + code)
	}
}

func init() {
	RootCmd.AddCommand(generateCmd)
}
