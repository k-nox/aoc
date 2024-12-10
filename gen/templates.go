package gen

const (
	PartTemplate = `package day{{ printf "%02d" .Day }}

import (
	"github.com/k-nox/aoc/util"
)

func Part{{ .Part }}(useSample bool) int {
	f := util.NewScannerForInput({{ .Year }}, {{ .Day }}, useSample)
	defer f.Close()

	for f.Scan() {

	}
	
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
