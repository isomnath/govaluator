package utilities

import (
	"bytes"
	"text/template"
)

func (utils *Utilities) parseExp(tmpResult map[string]interface{}, expression string) (string, error) {
	var b bytes.Buffer
	temp, err := template.New("").Parse(expression)
	if err != nil {
		return "", err
	}
	err = temp.Execute(&b, tmpResult)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
