package gen

import "text/template"

const (
	PartTemplate = `package day{{ printf "%02d" .Day }}

func Part{{ .Part }}(useSample bool) int {
	
	return 0
}
`

	MainTemplate = `// Code generated; DO NOT EDIT.
// This file was generated at
// {{ .Timestamp }}
package main

import (
	"log"
	"os"
	"github.com/k-nox/aoc/cli"
{{- range .Days }}
	"{{$.ModuleName}}/{{$.Year}}/{{.}}"
{{- end }}
)

var registry = cli.Registry{
{{- range .Days }}
	"{{ . }}": {PartOne: {{ . }}.PartOne, PartTwo: {{ . }}.PartTwo},
{{- end }}
}	

func main() {
	app := cli.App(registry, "{{ .ModuleName }}")
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
`
)

type templ struct {
	text string
	file string
	name string
}

func (t templ) parse() (*template.Template, error) {
	if t.file != "" {
		return template.ParseFiles(t.file)
	}
	return template.New(t.name).Parse(t.text)
}
