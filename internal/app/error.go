package app //nolint: typecheck

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
