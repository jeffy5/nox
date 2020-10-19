package template

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

// FileType depends on what operation work on the template file.
type FileType uint32

const (
	// FileTypeCreate means create a new file.
	FileTypeCreate = iota

	// FileTypeModify means modify existed file.
	FileTypeModify
)

var (
	presetFuncMap = template.FuncMap{
		"split": strings.Split,
		"sub": func(i, s int) int {
			return i - s
		},
		"lowerfirst": func(s string) string {
			if s == "" {
				return s
			}

			r := []rune(s)
			first := strings.ToLower(string(r[0]))
			res := make([]rune, 0)
			res = append(res, []rune(first)...)
			res = append(res, r[1:]...)
			return string(res)
		},
	}
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

// Generate file.
func (f *File) Generate(optionMap map[string]string) error {
	if _, err := os.Stat(f.Name); os.IsNotExist(err) {
		fmt.Printf("Create file %s.\n", f.Name)
		return f.create(optionMap)
	}
	fmt.Printf("Modify file: %s.\n", f.Name)
	return f.modify(optionMap)
}

func (f *File) create(optionMap map[string]string) error {
	// Mkdir for destination file
	err := os.MkdirAll(filepath.Dir(f.Name), 0755)
	if err != nil {
		return fmt.Errorf("mkdir failed, %w", err)
	}

	// Open destination file
	destFile, err := os.OpenFile(f.Name, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("open destination file failed, %w", err)
	}

	// Read template content
	tmplFile, err := os.OpenFile(f.Template, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("open template file failed, %w", err)
	}
	tmplContent, err := ioutil.ReadAll(tmplFile)
	if err != nil {
		return fmt.Errorf("read file content failed, %w", err)
	}

	// Execute template content and write to destination file
	tmpl := template.Must(template.New(f.Name).Funcs(presetFuncMap).Parse(string(tmplContent)))
	err = tmpl.Execute(destFile, optionMap)
	if err != nil {
		return fmt.Errorf("execute template failed, %w", err)
	}
	return nil
}

func (f *File) modify(optionMap map[string]string) error {
	// Open template file
	tmplFile, err := os.OpenFile(f.Template, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("open template file failed, %w", err)
	}
	tmplContent, err := ioutil.ReadAll(tmplFile)
	if err != nil {
		return fmt.Errorf("read template file content failed, %w", err)
	}

	// Parse template
	buf := bytes.NewBuffer([]byte{})
	tmpl := template.Must(template.New(f.Name).Funcs(presetFuncMap).Parse(string(tmplContent)))
	err = tmpl.Execute(buf, optionMap)
	if err != nil {
		return fmt.Errorf("template execute failed, %w", err)
	}

	// Check if advance is empty.
	if f.Advance.IsEmpty() {
		fmt.Println("\tAppend by default.")
		// Insert after the last line by default.
		destFile, err := os.OpenFile(f.Name, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return fmt.Errorf("open destination file failed, %w", err)
		}
		_, err = destFile.Write(buf.Bytes())
		if err != nil {
			return fmt.Errorf("append destination file failed, %w", err)
		}
	} else {
		// Modify content according to advance settings.
		destFile, err := os.OpenFile(f.Name, os.O_RDONLY, 0644)
		if err != nil {
			return fmt.Errorf("open destination file failed, %w", err)
		}
		destContent, err := ioutil.ReadAll(destFile)
		if err != nil {
			return fmt.Errorf("read destination file content failed, %w", err)
		}

		var newDestContent string
		if f.Advance.Before != "" {
			fmt.Println("\tInsert before specific line.")
			// Insert before pattern.
			reg := regexp.MustCompile(f.Advance.Before)
			newDestContent = reg.ReplaceAllString(string(destContent), fmt.Sprintf("%s\n${0}", buf.String()))
		} else if f.Advance.After != "" {
			fmt.Println("\tInsert after specific line.")
			// Insert after pattern.
			reg := regexp.MustCompile(f.Advance.After)
			newDestContent = reg.ReplaceAllString(string(destContent), fmt.Sprintf("${0}%s\n", buf.String()))
		}

		err = ioutil.WriteFile(f.Name, []byte(newDestContent), 0655)
		if err != nil {
			return fmt.Errorf("modify destination file content failed, %w", err)
		}
	}
	return nil
}

// FileAdvance describe the file advance operation, you can insert after or before specific given regex line.
type FileAdvance struct {
	Before string
	After  string
}

// IsEmpty return true if the content of before and after is empty
func (fileAdvance *FileAdvance) IsEmpty() bool {
	return fileAdvance.Before == "" && fileAdvance.After == ""
}
