package cmd

import (
	"github.com/fabiotavarespr/Planilha-MTG/sets"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var setsCmd = &cobra.Command{
	Use:   "sets",
	Short: "Display alls sets of Magic Editions",
	Run: func(cmd *cobra.Command, args []string) {
		sets.ListarSets()
	},
}

func init() {
	rootCmd.AddCommand(setsCmd)
}
