package gen

const (
	PartTemplate = `
package day{{ printf "%02d" .Day }}

import (
	"github.com/k-nox/aoc/util"
)

func Part{{ .Part }}(useSample bool) int {
	f := util.NewScannerForInput({{ .Day }}, useSample)
	defer f.Close()

	for f.Scan() {

	}
	
	return 0
}

`
	RegistryTemplate = `
// Code generated; DO NOT EDIT.
// This file was generated at
// {{ .Timestamp }}
package cli

import (
{{- range .Days }}
	"{{$.ModuleName}}/{{printf "day%02d" .}}"
{{- end }}
)

type part func(useSample bool) int
type day struct {
	PartOne part
	PartTwo part
}

var registry = map[int]day{
{{- range .Days }}
	{{ . }}: {day{{ printf "%02d" .}}.PartOne, day{{printf "%02d" .}}.PartTwo},
{{- end }}
}	
`
)
