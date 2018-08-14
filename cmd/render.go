package cmd

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/Masterminds/sprig"
)

type inputVars interface{}

func Render(cfg *Config, inputJSON []byte) ([]byte, error) {
	tmpl, err := loadTemplate(cfg.TemplatePath)
	if err != nil {
		return nil, err
	}
	return doRender(inputJSON, tmpl)
}

func doRender(inputJSON []byte, tmpl *template.Template) ([]byte, error) {
	out := new(bytes.Buffer)
	var vars inputVars
	if err := json.Unmarshal(inputJSON, &vars); err != nil {
		return nil, err
	}
	if err := tmpl.Execute(out, vars); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func loadTemplate(templatePath string) (*template.Template, error) {
	return template.New(templatePath).Funcs(sprig.FuncMap()).ParseFiles(templatePath)
}
