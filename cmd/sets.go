package cmd

import (
	"github.com/fabiotavarespr/Planilha-MTG/sets"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var setsCmd = &cobra.Command{
	Use:     "sets",
	Aliases: []string{"s"},
	Short:   "Display alls sets of Magic Editions",
	Long: `Display alls sets of Magic Editions.
Separate by Name and Code
	`,
	Run: func(cmd *cobra.Command, args []string) {
		sets.ListarSets()
	},
}

func init() {
	RootCmd.AddCommand(setsCmd)
}
