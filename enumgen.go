package enumgen

import (
	"fmt"
	"os"
	"strings"
)

func GenerateFiles(path string, data *TemplateData) ([]string, error) {
	// check if the path exists and that it contains a trailing slash
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("path does not exist: %w", err)
	}

	// check if the path is empty or only contains a slash
	if path == "" || path == "/" {
		return nil, &PathError{Path: path}
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	filename, errFn := generateFilename(data.EnumTypeName)
	if errFn != nil {
		return nil, fmt.Errorf("error generating filename: %w", errFn)
	}

	genFiles := []string{path + filename + ".go", path + filename + "_test.go"}

	// .go file
	if err := GenerateFile(genFiles[0], "enum.go.tmpl", data); err != nil {
		return nil, fmt.Errorf("error generating file: %w", err)
	}

	// _test.go file
	if err := GenerateFile(path+filename+"_test.go", "enum_test.go.tmpl", data); err != nil {
		return []string{genFiles[0]}, fmt.Errorf("error generating file: %w", err)
	}

	return genFiles, nil
}

func GenerateFile(filename, templateFile string, data *TemplateData) error {
	formatted, errG := generateFromTmpl(templateFile, data)
	if errG != nil {
		return fmt.Errorf("error generating from template: %w", errG)
	}

	if errW := writeToFile(filename, formatted); errW != nil {
		return fmt.Errorf("error writing to file: %w", errW)
	}

	return nil
}
