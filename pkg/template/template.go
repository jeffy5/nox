package template

import (
	"bytes"
	"encoding/json"
	"errors"
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
	var tmpl Template

	// Open template manifest file
	templateFile := path.Join(templateURL, "manifest.json")
	if _, err := os.Stat(templateFile); os.IsNotExist(err) {
		return &tmpl, errors.New("template file does not exist")
	}
	file, err := os.OpenFile(templateFile, os.O_RDONLY, 0644)
	if err != nil {
		return &tmpl, fmt.Errorf("open template file failed, %s", err)
	}
	manifestRaw, err := ioutil.ReadAll(file)
	if err != nil {
		return &tmpl, fmt.Errorf("read file failed: %s", err)
	}

	// Parse manifest content
	err = tmpl.parseManifest(manifestRaw)
	if err != nil {
		return &tmpl, fmt.Errorf("parse manifest failed, %w", err)
	}

	// Parse option's values
	if len(options) != len(tmpl.Manifest.Options) {
		return &tmpl, fmt.Errorf("invalid options, required %d options, but only %d",
			len(tmpl.Manifest.Options), len(options))
	}
	for i := range tmpl.Manifest.Options {
		tmpl.Manifest.Options[i].Value = options[i]
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
		tmpl.Manifest.Files[i].Template = path.Join(templateURL, file.Template)
	}
	return &tmpl, nil
}

// Generate start the task to create and modify files.
func (tmpl *Template) Generate() error {
	var err error
	for _, f := range tmpl.Manifest.Files {
		err = f.Generate(tmpl.Manifest.OptionMap())
		if err != nil {
			return err
		}
	}
	return nil
}

// parseManifest parse the raw manifest raw data.
func (tmpl *Template) parseManifest(raw []byte) error {
	// Parse manifest
	var manifest Manifest
	err := json.Unmarshal(raw, &manifest)
	if err != nil {
		return fmt.Errorf("json unmarshal failed, %w", err)
	}
	tmpl.Manifest = manifest
	return nil
}

// interact with user to parse manifest's options
func (tmpl *Template) interact() error {
	return nil
}
