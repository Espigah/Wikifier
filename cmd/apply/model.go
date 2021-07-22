package apply //nolint: typecheck

import "bitbucket.org/git-fsrg/wikifier/internal/app"

type model struct {
	metaData  app.MetaData
	triggers  []*model
	dependsOn *model
	hasResult bool
}

func (e *model) IsRoot() bool {
	if e == nil {
		return false
	}
	return e.metaData.IsRoot()
}
