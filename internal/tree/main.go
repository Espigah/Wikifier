package tree //nolint: typecheck

import (
	"path/filepath"
	"strings"
)

type node struct {
	page     *Page
	pathList []string
	path     string
}

func Create(root string) map[string]*Page {
	pageMap := make(map[string]*Page)

	getPage := func(key string) *Page {
		if page, ok := pageMap[key]; ok {
			return page
		} else {
			page := &Page{
				Key: key,
			}
			pageMap[key] = page
			return page
		}
	}

	nodes := make([]*node, 0)
	for _, file := range Find(root, extension) {
		path := strings.Replace(file, root, "", 1)
		pathList := strings.Split(path, "/")
		pathList = pathList[1:] // remove root
		length := len(pathList)

		var leafnode *Page
		var key string

		if strings.Contains(file, index) {
			key = strings.Join(pathList[:length-1], "/")
			// if it's the index, get title from folder
			leafnode = &Page{
				File:  file,
				Title: pathList[length-2],
			}
		} else if strings.Contains(file, extension) {
			key = strings.Join(pathList[:length], "/")
			// if it's a file, get title from filename
			basename := pathList[length-1]
			leafnode = &Page{
				File:  file,
				Title: strings.TrimSuffix(basename, filepath.Ext(basename)),
			}
		}

		nodes = append(nodes, &node{
			page:     leafnode,
			pathList: pathList,
			path:     key,
		})
		pageMap[key] = leafnode
	}

	for _, node := range nodes {
		leafnode := node.page
		pathList := node.pathList
		length := len(pathList)

		// 1 - because is root index
		if length == 1 {
			continue
		}

		list := pathList[:length-1]

		lastPage := leafnode
		for i := len(list) - 1; i >= 0; i-- {
			key := strings.Join(list[:i+1], "/")
			page := getPage(key)

			if page == lastPage {
				continue
			}

			page.Title = pathList[i]

			page.Children = append(page.Children, lastPage)
			lastPage.Parent = page
			lastPage = page
			if page.File == "" {
				page.IsFolder = true
				page.File = filepath.Join(root, key)
			}
		}

	}

	return pageMap
}
