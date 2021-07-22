package app //nolint: typecheck

type config struct {
	Output string
	Input  string
}

var Config config //nolint: gochecknoglobals
