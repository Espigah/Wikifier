package adapters //nolint: typecheck

import "bitbucket.org/git-fsrg/wikifier/internal/app"

type Wiki interface {
	Setup() error
	Create(*app.MetaData) bool
	Delete(*app.MetaData) bool
	Update(*app.MetaData) bool
}
