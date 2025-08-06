package template

import (
	"bytes"
	"text/template"
)

var funcMap = map[string]any{
	// преобразование плейсхолдеров
}

// Execute исполняет шаблон и в случае ошибки отдает пустую строку
func Execute(tmpl string, value any) string {
	t, err := template.New("").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return ""
	}

	var buff bytes.Buffer
	if err = t.Execute(&buff, value); err != nil {
		return ""
	}
	return buff.String()
}
