package apply //nolint: typecheck

import "bitbucket.org/git-fsrg/wikifier/internal/app"

func applyChanges(changes []app.MetaData) {
	executionMap := make(map[string]*model)

	getExecution := func(key string) *model {
		if executionNoe, ok := executionMap[key]; ok {
			return executionNoe
		} else {
			page := &model{}
			executionMap[key] = page
			return page
		}
	}

	// Organize execution order
	for _, md := range changes {
		parentNode := getExecution(md.Parent)

		executionNode := &model{
			metaData:  md,
			dependsOn: parentNode,
		}

		parentNode.triggers = append(parentNode.triggers, executionNode)

		executionMap[md.MetaFile] = executionNode
	}

	for _, executionNode := range executionMap {
		executeChanges(executionNode)
	}
}

func executeChanges(executionNode *model) {
	if executionNode.dependsOn != nil {
		executeChanges(executionNode.dependsOn)
	}
	if executionNode.IsRoot() || executionNode.metaData.Status == app.STATUS_CREATED {
		executionNode.hasResult = true
	}
	if executionNode.hasResult {
		return
	}

	execute(&executionNode.metaData)

	executionNode.hasResult = true
	for _, trigger := range executionNode.triggers {
		executeChanges(trigger)
	}
}

func execute(metaData *app.MetaData) {
	switch metaData.Status {
	case app.STATUS_DELETED:

		if ok := wiki.Delete(metaData); ok {
			metaData.AutoDelete()
		}

	case app.STATUS_PENDING:

		if ok := wiki.Create(metaData); ok {
			metaData.Status = app.STATUS_CREATED
			metaData.AutoSave()
		}

	}
}
