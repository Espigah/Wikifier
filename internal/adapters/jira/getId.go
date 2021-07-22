package jira //nolint: typecheck

import (
	"encoding/json"
	"errors"
	"fmt"

	"bitbucket.org/git-fsrg/wikifier/internal/app"
	"github.com/imroc/req"
)

func getIdFromTitle(title string) (string, error) {
	if id, ok := info.pageMap[title]; ok {
		return id, nil
	}

	if info.lastPageIndex > 1000 {
		return "", errors.New("Max ")
	}

	r := req.New()
	authHeader := req.Header{
		"Accept":        "application/json",
		"Authorization": "Basic " + apiKey,
	}

	url := fmt.Sprintf("%s?start=%v&limit=10", apiURL, info.lastPageIndex)

	resultList, err := r.Get(url, authHeader)
	if err != nil {
		return "", err
	}

	body, err := resultList.ToBytes()
	if err != nil {
		return "", err
	}

	var obj response

	if err := json.Unmarshal(body, &obj); err != nil {
		return "", err
	}

	if len(obj.Results) == 0 {
		return "", errors.New("Id by title not found " + title)
	}

	for _, result := range obj.Results {
		id := result.Id
		title := result.Title
		info.pageMap[title] = id
		info.lastPageIndex++
	}

	return getIdFromTitle(title)
}

func getIdFromMetaData(metaData *app.MetaData) string {
	if metaData.ID == "" {
		id, err := getIdFromTitle(metaData.PageTitle)
		if err != nil {
			panic(err)
		}
		metaData.ID = id
		metaData.AutoSave()
	}
	return metaData.ID
}
