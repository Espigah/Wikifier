package version

import (
	"fmt"

	"bitbucket.org/git-fsrg/wikifier/cmd"
	"github.com/spf13/cobra"
)

var (
	Commit    = "" //nolint: gochecknoglobals
	BuildDate = "" //nolint: gochecknoglobals
	Version   = "" //nolint: gochecknoglobals
)

func init() { //nolint: gochecknoinits
	cmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Shows current version of CLI",
		Long:  `Shows current version of CLI`,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s - %s Commit: %s\n", Version, BuildDate, Commit)
		},
	})
}
