package tree //nolint: typecheck

var (
	index     = "index.md"
	extension = ".md"
)

type Page struct {
	Title    string
	File     string
	Key      string
	IsFolder bool
	Parent   *Page
	Children []*Page
}

func (page *Page) GetParentTitle() string {
	if page == nil || page.Parent == nil {
		return ""
	}
	return page.Parent.Title
}
