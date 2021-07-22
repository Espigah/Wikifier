package jira //nolint: typecheck

import "bitbucket.org/git-fsrg/wikifier/internal/adapters"

type page struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type response struct {
	Results []*page `json:"results"`
}

type jiraInfo struct {
	pageMap       map[string]string
	lastPageIndex int
}

type storage struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}

type body struct {
	Storage *storage `json:"storage"`
}

type space struct {
	Key string `json:"key"`
}

type ancestor struct {
	Id string `json:"id"`
}

type jira struct {
	Type      string     `json:"type"`
	Title     string     `json:"title"`
	Ancestors []ancestor `json:"ancestors"`
	Space     *space     `json:"space"`
	Body      *body      `json:"body"`
}

func New() adapters.Wiki {
	return &jira{}
}
