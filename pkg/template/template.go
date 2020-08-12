package template

// FileType depends on what operation work on the template file
type FileType int

const (
	// FileTypeNew means create a all new template file
	FileTypeNew = iota

	// FileTypeAppend means template file's data should be append data to existed file
	FileTypeAppend
)

// ArgType determine how to interact with user
type ArgType int

const (
	// ArgTypeInput
	ArgTypeInput = iota

	ArgTypeCheckbox
)

// Template include template's infomation
type Template struct {
	Manifest Manifest
	Args     []Args
}

// Manifest include all files' information
type Manifest struct {
	Version string   `json:"version"`
	Files   []string `json:"files"`
	Remark  string   `json:"remark"`
}

// File describe the file included in manifest
type File struct {
	Name string   `json:"name"`
	Type FileType `json:"type"`
}

// Args describe template file
type Args struct {
	Type  ArgType `json:"type"`
	Value string  `json:"value"`
}
