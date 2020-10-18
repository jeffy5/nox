package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

// Template include template's infomation.
type Template struct {
	Manifest Manifest
	Args     []Args
}

// NewTemplate return a new template instance.
func NewTemplate(templateURL string, options []string) (*Template, error) {
	// Open template manifest file
	file, err := os.OpenFile(path.Join(templateURL, "manifest.json"), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(fmt.Sprintf("open template file failed, %s", err))
	}
	manifestRaw, err := ioutil.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("read file failed: %s", err))
	}

	// Parse manifest content
	var tmpl Template
	err = tmpl.parseManifest(manifestRaw)
	if err != nil {
		return &tmpl, err
	}

	// Parse option's values
	if len(options) != len(tmpl.Manifest.Options) {
		return &tmpl, fmt.Errorf("invalid options, required %d options, but only %d",
			len(tmpl.Manifest.Options), len(options))
	}

	// Parse files' name
	for i, file := range tmpl.Manifest.Files {
		t := template.Must(template.New(file.Name).Parse(file.Name))

		buf := bytes.NewBuffer([]byte{})
		err = t.Execute(buf, tmpl.Manifest.OptionMap())
		if err != nil {
			return &tmpl, fmt.Errorf("parse file name failed, %s", err)
		}
		tmpl.Manifest.Files[i].Name = buf.String()
	}
	return &tmpl, nil
}

// parseManifest parse the raw manifest raw data.
func (template *Template) parseManifest(raw []byte) error {
	// Parse manifest
	var manifest Manifest
	err := json.Unmarshal(raw, &manifest)
	if err != nil {
		return fmt.Errorf("parse manifest failed -> %w", err)
	}
	return nil
}

// interact with user to parse manifest's options
func (template *Template) interact() error {
	return nil
}
