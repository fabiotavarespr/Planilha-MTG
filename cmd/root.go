package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fabiotavarespr/Planilha-MTG/api"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   BinaryName,
	Short: "Generating a spreadsheet of Magic: The Gathering",
	Long:  "Generating a spreadsheet of Magic: The Gathering edits you request",

	SilenceUsage: true,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	export.ApresentaListagem("THB")
	// },
}

// Execute adds all child commands to the root command.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	BinaryName = os.Args[0]
	Out = os.Stdout
	Err = os.Stderr
	api.UserAgent = fmt.Sprintf("github.com/fabiotavarespr/Planilha-MTG v%s (%s/%s)", Version, runtime.GOOS, runtime.GOARCH)
}
