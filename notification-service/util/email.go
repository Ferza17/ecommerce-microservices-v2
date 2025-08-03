package util

import (
	"bytes"
	"text/template"
	"time"
)

func ParseEmailTemplate(htmlTemplateString string, data map[string]any) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	tmpl := template.Must(template.New("template").Funcs(
		template.FuncMap{
			"formatDate": func(t time.Time, layout string) string {
				return t.Format("January 2, 2006 at 3:04 PM")
			},
		},
	).Parse(htmlTemplateString))

	if err := tmpl.Execute(buf, data); err != nil {
		return nil, err
	}

	return buf, nil
}
