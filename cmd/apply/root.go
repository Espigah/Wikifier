package apply //nolint: typecheck

import (
	"bitbucket.org/git-fsrg/wikifier/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{ //nolint: gochecknoglobals
	Use:   "apply",
	Short: "create more commands",
	Long:  `create more commands`,
	Run: func(cmd *cobra.Command, args []string) {
		apply()
	},
}

func init() { //nolint: gochecknoinits
	cmd.AddCommand(rootCmd)
}
