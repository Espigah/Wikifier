package main //nolint: typecheck

import (
	"bitbucket.org/git-fsrg/wikifier/cmd"
	_ "bitbucket.org/git-fsrg/wikifier/cmd/apply"
	_ "bitbucket.org/git-fsrg/wikifier/cmd/generate"
	_ "bitbucket.org/git-fsrg/wikifier/cmd/setup"
	"bitbucket.org/git-fsrg/wikifier/cmd/version"
	//%MAIN_IMPORT%
)

var (
	Commit    = "" //nolint: gochecknoglobals
	BuildDate = "" //nolint: gochecknoglobals
	Version   = "" //nolint: gochecknoglobals
)

func main() {
	version.Version = Version
	version.Commit = Commit
	version.BuildDate = BuildDate
	cmd.Execute()
}
