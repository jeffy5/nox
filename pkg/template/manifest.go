package template

// Manifest includes all files' information
type Manifest struct {
	// Version tag
	Version string

	// Author describes the template owner information
	Author Author

	Options []Option

	// Files which should be created or modified
	Files []File
}

// Author infomations
type Author struct {
	Name  string
	Email string
}
