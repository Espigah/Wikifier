package generate //nolint: typecheck

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"bitbucket.org/git-fsrg/wikifier/internal/app"
	"bitbucket.org/git-fsrg/wikifier/internal/tree"
)

var root = "wikifier"

type Model struct {
	page         *tree.Page
	metafilePath string
	contentPath  string
}

func generate() {
	pageMap := tree.Create(path)
	modelMap := createModelMap(pageMap)
	createMetaFiles(modelMap)
	syncRemovedFiles()
}

func createModelMap(pageMap map[string]*tree.Page) map[string]*Model {
	modelMap := make(map[string]*Model)
	for _, page := range pageMap {
		var hashkey string
		if page.IsFolder {
			hashkey = page.Key
		} else {
			hashkey = page.File
		}
		fileName := page.Title
		nameSuffix := fmt.Sprint(hash(hashkey))
		metafilePath := makePath(root, fileName, nameSuffix, "wkf")
		contentPath := makePath(root, fileName, nameSuffix, format)

		modelMap[fileName] = &Model{
			page:         page,
			metafilePath: metafilePath,
			contentPath:  contentPath,
		}

	}
	return modelMap
}

func createMetaFiles(modelMap map[string]*Model) {
	for _, s := range modelMap {
		page := s.page
		file := page.File

		var metafile app.MetaData
		metafileLoadError := metafile.Load(s.metafilePath)
		var fileSum string
		var contentPath string

		if !page.IsFolder {
			fileBytes, err := ioutil.ReadFile(file)
			app.Check(err)
			//
			content := parseFile(fileBytes, format)
			writeContent(content, s.contentPath)
			//
			h := sha256.New()
			r := bytes.NewReader(fileBytes)
			io.Copy(h, r)
			fileSum = string(h.Sum(nil))

			contentPath = s.contentPath
		}

		parentSeila := modelMap[page.GetParentTitle()]
		var parentmetafile string
		if parentSeila != nil {
			parentmetafile = parentSeila.metafilePath
		}
		ei := app.MetaData{
			PageTitle:   page.Title,
			Origin:      file,
			Timestamp:   time.Now().String(),
			ContentFile: contentPath,
			Format:      format,
			Sum:         fileSum,
			Parent:      parentmetafile,
			MetaFile:    s.metafilePath,
			IsFolder:    page.IsFolder,
		}

		if metafileLoadError == nil || (metafile.Sum == ei.Sum && metafile.Origin == ei.Origin) {
			continue
		}

		ei.Status = app.STATUS_PENDING
		ei.Save(s.metafilePath)
	}
}

func syncRemovedFiles() {
	for _, filepath := range tree.Find(root, ".wkf") {
		var c app.MetaData
		c.Load(filepath)
		if !checkFileExists(c.Origin) {
			c.Status = app.STATUS_DELETED
			c.Save(filepath)
			return
		}
	}
}
