package template

// FileType depends on what operation work on the template file
type FileType uint32

const (
	// FileTypeNew means create a all new template file
	FileTypeNew = iota

	// FileTypeAppend means template file's data should be append data to existed file
	FileTypeAppend
)

// File describe the file included in manifest
type File struct {
	Name    string
	Type    FileType
	Context Context
}
