package template

// Manifest includes all files' information
type Manifest struct {
	Author Author

	// Files which should be created or modified
	Files []File
}

// Author infomations
type Author struct {
	Name  string
	Email string
}
