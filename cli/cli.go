package cli

import (
	"github.com/urfave/cli/v2"
)

func StandaloneApp() *cli.App {
	app := cli.NewApp()
	app.Name = "aoc"
	app.Usage = "A CLI to make a life a little easier when solving Advent of Code puzzles"
	app.Suggest = true
	app.EnableBashCompletion = true

	flagMap := flags()

	app.Commands = []*cli.Command{
		{
			Name:    "gen",
			Aliases: []string{"g"},
			Usage:   "Generate files for specified day",
			Flags: []cli.Flag{
				flagMap[day],
				flagMap[year],
				flagMap[force],
				flagMap[path],
				flagMap[partTemplate],
				flagMap[mainTemplate],
			},
			Action: generate,
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Run solutions for specified day",
			Flags: []cli.Flag{
				flagMap[day],
				flagMap[year],
				flagMap[path],
				flagMap[sample],
			},
			Action: runStandalone,
		},
	}

	return app
}

func App(registry Registry, moduleName string) *cli.App {

	app := cli.NewApp()
	app.Suggest = true
	app.EnableBashCompletion = true

	flagMap := flags()
	app.Flags = []cli.Flag{
		flagMap[day],
		flagMap[sample],
	}

	app.Usage = "Run given day's solutions"
	app.Action = run(registry)

	return app
}
