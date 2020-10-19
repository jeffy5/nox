package template

import (
	"encoding/json"
)

type marshalManifest struct {
	// Version tag.
	Version string `json:"version"`

	// Author describes the template owner information.
	Author Author `json:"author"`

	// Options determines the interaction between the template and the user.
	Options []Option `json:"options"`

	// Files which should be created or modified.
	Files marshalManifestFiles `json:"files"`

	// Hooks is a command list which will execute after all template files generated.
	Hooks []string `json:"hooks"`
}

type marshalManifestFiles struct {
	Create []marshalManifestFile `json:"create"`
	Modify []marshalManifestFile `json:"modify"`
}

type marshalManifestFile struct {
	Template    string                     `json:"template"`
	Destination string                     `json:"destination"`
	Advance     marshalManifestFileAdvance `json:"advance"`
}

type marshalManifestFileAdvance struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

// Manifest includes all files' information
type Manifest struct {
	// Version tag.
	Version string

	// Author describes the template owner information.
	Author Author

	// Options determines the interaction between the template and the user.
	Options []Option

	// Files which should be created or modified.
	Files []File

	// Hooks is a command list which will execute after all template files generated.
	Hooks []string
}

// UnmarshalJSON override json unmarshal.
func (manifest *Manifest) UnmarshalJSON(b []byte) error {
	var mManifest marshalManifest
	err := json.Unmarshal(b, &mManifest)
	if err != nil {
		return err
	}

	manifest.Version = mManifest.Version
	manifest.Author = mManifest.Author
	manifest.Hooks = mManifest.Hooks
	manifest.Options = mManifest.Options

	files := make([]File, 0)
	for _, f := range mManifest.Files.Create {
		advance := FileAdvance{
			Before: f.Advance.Before,
			After:  f.Advance.After,
		}
		file := File{
			Name:     f.Destination,
			Type:     FileTypeCreate,
			Template: f.Template,
			Advance:  advance,
		}
		files = append(files, file)
	}
	for _, f := range mManifest.Files.Modify {
		advance := FileAdvance{
			Before: f.Advance.Before,
			After:  f.Advance.After,
		}
		file := File{
			Name:     f.Destination,
			Type:     FileTypeModify,
			Template: f.Template,
			Advance:  advance,
		}
		files = append(files, file)
	}
	manifest.Files = files
	return nil
}

// OptionMap return the option list with map type
func (manifest *Manifest) OptionMap() map[string]string {
	optionMap := make(map[string]string)
	for _, option := range manifest.Options {
		optionMap[option.Name] = option.Value
	}
	return optionMap
}

// Author infomations.
type Author struct {
	Name  string
	Email string
}
