package apply //nolint: typecheck

import (
	"fmt"

	"bitbucket.org/git-fsrg/wikifier/internal/adapters"
	"bitbucket.org/git-fsrg/wikifier/internal/adapters/jira"
)

//@ TODO load from config file
var (
	root = "wikifier"
	wiki adapters.Wiki
)

func apply() {
	wiki = jira.New() //@ TODO update to use DI
	wiki.Setup()
	changes := findChanges(root)
	toApply := prompt(changes)
	fmt.Printf("To Apply: %v\n", toApply)
	if !toApply {
		return
	}
	applyChanges(changes)
}
