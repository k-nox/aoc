package cli

import (
	"fmt"
	"time"

	"github.com/k-nox/aoc/gen"
	"github.com/urfave/cli/v2"
)

func App(registry Registry, moduleName string) *cli.App {

	app := cli.NewApp()
	app.Name = "aoc"
	app.Usage = "A CLI to make a life a little easier when solving Advent of Code puzzles"
	app.DefaultCommand = "run"
	app.Suggest = true
	app.EnableBashCompletion = true

	dayFlag := &cli.IntFlag{
		Name:        "day",
		Aliases:     []string{"d"},
		Usage:       "Specify day",
		DefaultText: "today",
		Value:       time.Now().Day(),
	}

	app.Commands = []*cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Run one day's solutions",
			Action:  run(registry),
			Flags: []cli.Flag{
				dayFlag,
				&cli.BoolFlag{
					Name:        "sample",
					Aliases:     []string{"s"},
					Usage:       "Use sample input",
					DefaultText: "false",
					Value:       false,
				},
			},
		},
		{
			Name:    "gen",
			Aliases: []string{"g"},
			Usage:   "Generate files for specified day",
			Action:  generate(registry, moduleName),
			Flags: []cli.Flag{
				dayFlag,
				&cli.BoolFlag{
					Name:        "force",
					Aliases:     []string{"f"},
					Usage:       "Force generation - may overwrite existing files",
					DefaultText: "false",
					Value:       false,
				},
			},
		},
	}

	return app
}

func generate(registry Registry, moduleName string) func(*cli.Context) error {
	return func(c *cli.Context) error {
		day := c.Int("day")
		force := c.Bool("force")
		if _, ok := registry[day]; ok {
			if !force {
				return cli.Exit(fmt.Sprintf("day %02d already exists; use --force if you really want to overwrite existing files", day), 1)
			}
			fmt.Printf("force applied - may overwrite existing files")
		}
		fmt.Printf("generating files for day %02d\n", day)
		err := gen.Generate(day, moduleName)
		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Println("generated files")
		return nil
	}
}

func run(registry Registry) func(*cli.Context) error {
	return func(c *cli.Context) error {
		day := c.Int("day")
		useSample := c.Bool("sample")

		d, ok := registry[day]
		if !ok {
			return cli.Exit(fmt.Sprintf("unregistered day requested: %d", day), 1)
		}

		fmt.Printf("solution for day %d part one: %d\n", day, d.PartOne(useSample))
		fmt.Printf("solution for day %d part two: %d\n", day, d.PartTwo(useSample))

		return nil
	}
}
