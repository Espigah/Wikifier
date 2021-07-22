package jira //nolint: typecheck

import "github.com/imroc/req"

func (j *jira) Setup() error {
	req.Debug = true
	return nil
}
