package apply //nolint: typecheck

import (
	"fmt"

	"bitbucket.org/git-fsrg/wikifier/internal/adapters"
	"bitbucket.org/git-fsrg/wikifier/internal/adapters/jira"
	"bitbucket.org/git-fsrg/wikifier/internal/app"
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

func execute(executionNode *model) {
	if executionNode.dependsOn != nil {
		execute(executionNode.dependsOn)
	}
	if executionNode.IsRoot() || executionNode.metaData.Status == app.STATUS_CREATED {
		executionNode.hasResult = true
	}
	if executionNode.hasResult {
		return
	}

	syncPage(&executionNode.metaData)

	executionNode.hasResult = true
	for _, trigger := range executionNode.triggers {
		execute(trigger)
	}
}

func syncPage(metaData *app.MetaData) {
	switch metaData.Status {
	case app.STATUS_DELETED:
		wiki.Delete(metaData)
		metaData.AutoDelete()
	case app.STATUS_PENDING:
		wiki.Create(metaData)
		metaData.Status = app.STATUS_CREATED
		metaData.AutoSave()
	}
}
