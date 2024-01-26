package utility

import (
	"bytes"
	"text/template"
)

func ParseTemplateHTML(tmplPathFile string, data interface{}) (string, error) {
	template, err := template.ParseFiles(tmplPathFile)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err = template.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
