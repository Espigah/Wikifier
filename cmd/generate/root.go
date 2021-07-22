package generate //nolint: typecheck

import (
	"bitbucket.org/git-fsrg/wikifier/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{ //nolint: gochecknoglobals
	Use:   "plan",
	Short: "create more commands",
	Long:  `create more commands`,
	Run: func(cmd *cobra.Command, args []string) {
		generate()
	},
}

func init() { //nolint: gochecknoinits
	rootCmd.Flags().StringVarP(&page, "page", "p", "", "Page name")
	rootCmd.Flags().StringVarP(&format, "format", "t", "html", "Output file")
	rootCmd.Flags().StringVarP(&path, "path", "a", "", "Path to file")
	cmd.AddCommand(rootCmd)
}
