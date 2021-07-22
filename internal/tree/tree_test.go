package tree

import (
	"sort"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	pageMap := Create("../../wiki")
	var titles []string
	for _, page := range pageMap {
		titles = append(titles, page.Title)
	}

	sort.Strings(titles)
	if titles[0] != "home" {
		t.Errorf(`titles[0] = %q, should be %q`, titles[0], "home")
	}

	if titles[1] != "page1" {
		t.Errorf(`titles[1] = %q, should be %q`, titles[1], "page1")
	}

	if titles[2] != "page2" {
		t.Errorf(`titles[2] = %q, should be %q`, titles[2], "page2")
	}

	if titles[3] != "page3" {
		t.Errorf(`titles[3] = %q, should be %q`, titles[3], "page3")
	}

	if titles[4] != "page4" {
		t.Errorf(`titles[4] = %q, should be %q`, titles[4], "page4")
	}

	if titles[5] != "page5" {
		t.Errorf(`titles[5] = %q, should be %q`, titles[5], "page4")
	}
}

func TestFiles(t *testing.T) {
	pageMap := Create("wiki")
	var files []string
	for _, page := range pageMap {
		files = append(files, page.File)
	}

	sort.Strings(files)

	if files[0] != "" {
		t.Errorf(`files[0] = %q, should be %q`, files[0], "empty")
	}

	if files[1] != "" {
		t.Errorf(`files[1] = %q, should be %q`, files[1], "empty")
	}

	if files[2] != "wiki/home.md" {
		t.Errorf(`files[2] = %q, should be %q`, files[2], "wiki/home.md")
	}

	if files[3] != "wiki/page1/index.md" {
		t.Errorf(`files[3] = %q, should be %q`, files[3], "wiki/page1/index.md")
	}

	if files[4] != "wiki/page1/page2/page3/index.md" {
		t.Errorf(`files[4] = %q, should be %q`, files[4], "wiki/page1/page2/page3/index.md")
	}

	if files[5] != "wiki/page1/page2/page3/page4/page5.md" {
		t.Errorf(`files[5] = %q, should be %q`, files[5], "wiki/page1/page2/page3/page4/page5.md")
	}
}
