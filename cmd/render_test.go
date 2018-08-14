package cmd

import (
	"bytes"
	"html/template"
	"log"
	"testing"

	"github.com/Masterminds/sprig"
)

func TestDoRender(t *testing.T) {
	tests := []struct {
		name     string
		vars     []byte
		expected string
		wantErr  bool
		template string
	}{
		{
			name: "basic",
			vars: bytes.NewBufferString(`{"Name": "basic", "Slice": [1, 3, 5]}`).Bytes(),
			expected: `hello basic
0: 1
1: 3
2: 5`,
			template: `hello {{.Name}}
{{- range $i, $v := .Slice}}
{{$i}}: {{$v}}
{{- end}}`,
		},
		{
			name:     "json error",
			vars:     bytes.NewBufferString(`{`).Bytes(),
			wantErr:  true,
			template: `hello {{.Name}}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tmpl, err := parseTemplate(test.template)
			if err != nil {
				log.Fatal(err)
			}
			out, err := doRender(test.vars, tmpl)
			if test.wantErr {
				if err == nil {
					t.Errorf("expected to return error")
				}
			} else {
				if err != nil {
					log.Fatal(err)
				}
				if string(out) != test.expected {
					t.Errorf("expected:\n%s\n-----\noutput:\n%s", test.expected, out)
				}
			}
		})
	}
}

func parseTemplate(tmpl string) (*template.Template, error) {
	return template.New("test").Funcs(sprig.FuncMap()).Parse(tmpl)
}
