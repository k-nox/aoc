package cli

import (
	"time"

	"github.com/urfave/cli/v2"
)

const (
	day    = "day"
	year   = "year"
	sample = "sample"
	force  = "force"
	path   = "path"
)

func flags() map[string]cli.Flag {
	return map[string]cli.Flag{
		day: &cli.IntFlag{
			Name:        "day",
			Aliases:     []string{"d"},
			Usage:       "Specify day",
			DefaultText: "today",
			Value:       time.Now().Day(),
		},
		year: &cli.IntFlag{
			Name:        "year",
			Aliases:     []string{"y"},
			Usage:       "specify year",
			DefaultText: "current year",
			Value:       time.Now().Year(),
		},
		sample: &cli.BoolFlag{
			Name:    "sample",
			Aliases: []string{"s"},
			Usage:   "Use sample input",
			Value:   false,
		},
		force: &cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "Force generation - may overwrite existing files",
			Value:   false,
		},
		path: &cli.PathFlag{
			Name:    "path",
			Aliases: []string{"p"},
			Usage:   "Path to your advent of code solutions directory",
		},
	}
}
