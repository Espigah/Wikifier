package jira //nolint: typecheck

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"bitbucket.org/git-fsrg/wikifier/internal/app"
	"github.com/imroc/req"
)

func (j *jira) Create(metaData *app.MetaData) bool {
	var ancestors []ancestor
	if metaData.PageTitle == "page2" {
		fmt.Printf("%+v\n", metaData)
	}
	if metaData.Parent != "" {
		var parenMetaData app.MetaData
		parenMetaData.Load(metaData.Parent)
		parenId := getIdFromMetaData(&parenMetaData)
		// ancestors = append(ancestors, parenId)
		// ancestors = append(ancestors, &ancestor{Id: parenId})
		ancestors = []ancestor{ancestor{Id: parenId}}
	}
	var contentFile string
	if metaData.ContentFile != "" {
		file, err := ioutil.ReadFile(metaData.ContentFile) // just pass the file name
		if err != nil {
			panic(err)
		}
		contentFile = string(file)
	}

	ja := &jira{
		Type:      "page",
		Title:     metaData.PageTitle,
		Ancestors: ancestors,
		Space:     &space{Key: spaceKey},
		Body: &body{
			Storage: &storage{Value: contentFile, Representation: "storage"},
		},
	}

	r := req.New()
	authHeader := req.Header{
		"Accept":        "application/json",
		"Authorization": "Basic " + apiKey,
	}

	result, err := r.Post(apiURL, req.BodyJSON(&ja), authHeader)
	if err != nil {
		panic(err)
	}

	body, err := result.ToBytes() // ioutil.ReadAll(resultList.Response().Body)
	if err != nil {
		panic(err)
	}

	var obj page

	if err := json.Unmarshal(body, &obj); err != nil {
		panic(err)
	}

	if obj.Id == "" {
		panic(errors.New("Id not found"))
	}

	metaData.ID = obj.Id
	metaData.AutoSave()
	info.pageMap[obj.Title] = obj.Id
	return true
}
