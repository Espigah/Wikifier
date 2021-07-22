package jira //nolint: typecheck

import (
	"bitbucket.org/git-fsrg/wikifier/internal/app"
	"github.com/imroc/req"
)

func (j *jira) Delete(metaData *app.MetaData) bool {
	r := req.New()
	authHeader := req.Header{
		"Accept":        "application/json",
		"Authorization": "Basic " + apiKey,
	}

	_, err := r.Delete(apiURL+metaData.ID, authHeader)
	if err != nil {
		panic(err)
	}
	return true
}
