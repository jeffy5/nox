package template

// FileType depends on what operation work on the template file.
type FileType uint32

const (
	// FileTypeCreate means create a new file.
	FileTypeCreate = iota

	// FileTypeModify means modify existed file.
	FileTypeModify
)

// File describe the file included in manifest.
type File struct {
	// Name should be create of modify by template.
	Name string

	// Type marks what operation(create or modify) should be process.
	Type FileType

	// Template include file content's template.
	Template string

	// Advance describe the file advance operation, you can insert after or before specific given regex line.
	Advance FileAdvance
}

// FileAdvance describe the file advance operation, you can insert after or before specific given regex line.
type FileAdvance struct {
	Before string
	After  string
}
