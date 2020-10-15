package template

// Template include template's infomation
type Template struct {
	Manifest Manifest
	Args     []Args
}

// NewTemplate return a new template instance
func NewTemplate() *Template {
	return &Template{}
}
