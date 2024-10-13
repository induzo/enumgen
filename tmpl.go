package enumgen

import (
	"bytes"
	"fmt"
	"text/template"

	"mvdan.cc/gofumpt/format"
)

type TemplateData struct {
	PackageName       string
	EnumTypeName      string
	EnumTypeShortName string
	EnumValues        []string
}

func generateFromTmpl(templateFile string, data TemplateData) ([]byte, error) {
	tmpl, errP := template.New(templateFile).Funcs(template.FuncMap{"constMaker": constMaker}).ParseFiles(templateFile)
	if errP != nil {
		return nil, fmt.Errorf("error parsing template file: %w", errP)
	}

	// Execute the template
	var output bytes.Buffer

	if err := tmpl.Execute(&output, data); err != nil {
		return nil, fmt.Errorf("error executing template: %w", err)
	}

	// Format the output as Go code
	formatted, errS := format.Source(output.Bytes(), format.Options{
		ExtraRules: true,
	})
	if errS != nil {
		return nil, fmt.Errorf("error executing template: %w", errS)
	}

	return formatted, nil
}