package template

import (
	"encoding/json"
	"fmt"
)

// Template include template's infomation.
type Template struct {
	Manifest Manifest
	Args     []Args
}

// NewTemplate return a new template instance.
func NewTemplate(manifestRaw []byte) (*Template, error) {
	var template Template

	err := template.parseManifest(manifestRaw)
	if err != nil {
		return &template, err
	}
	return &template, nil
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
