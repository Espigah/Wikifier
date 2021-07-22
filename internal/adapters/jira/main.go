package jira //nolint: typecheck

// @todo read the settings from the config file
var (
	apiURL   = "https://fsrg.atlassian.net/wiki/rest/api/content/"
	apiKey   = "mFicmljaW8uZ29uY2FsdmVzQG00dS5jb20uYnI6bkRZc3dXSDZm"
	spaceKey = "FSRG"
	info     = &jiraInfo{
		pageMap: make(map[string]string),
	}
)

// https://docs.atlassian.com/atlassian-confluence/REST/6.6.0/#content-getContentById
