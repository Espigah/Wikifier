package generate //nolint: typecheck

import (
	"hash/fnv"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"bitbucket.org/git-fsrg/wikifier/internal/app"
	"github.com/gomarkdown/markdown"
)

func checkFileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func fileNameWithoutExtension(fileName string) string {
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}

func parseFile(file []byte, format string) []byte {
	return markdown.ToHTML(file, nil, nil)
}

func writeContent(content []byte, file string) {
	err := ioutil.WriteFile(file, content, 0644)
	app.Check(err)
}

func makePath(outputPath string, file string, hash string, format string) string {
	fileName := filepath.Base(file)
	var fileParsed strings.Builder
	fileParsed.WriteString(fileNameWithoutExtension(fileName))
	fileParsed.WriteString("__")
	fileParsed.WriteString(hash)
	fileParsed.WriteString(".")
	fileParsed.WriteString(format)

	return filepath.Join(outputPath, fileParsed.String())
}
