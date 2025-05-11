package util

import (
	"bytes"
	"text/template"
)

func ParseEmailTemplate(htmlTemplateString string, data map[string]any) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	tmpl := template.Must(template.New("template").Funcs(
		template.FuncMap{},
	).Parse(htmlTemplateString))

	if err := tmpl.Execute(buf, data); err != nil {
		return nil, err
	}

	return buf, nil
}
