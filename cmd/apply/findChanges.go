package apply //nolint: typecheck

import (
	"bitbucket.org/git-fsrg/wikifier/internal/app"
	"bitbucket.org/git-fsrg/wikifier/internal/tree"
)

func findChanges(root string) []app.MetaData {
	var changes []app.MetaData
	for _, filepath := range tree.Find(root, ".wkf") {
		var md app.MetaData
		md.Load(filepath)

		changes = append(changes, md)
	}
	return changes
}
