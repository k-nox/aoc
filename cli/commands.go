package cli

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/k-nox/aoc/gen"
	"github.com/urfave/cli/v2"
)

func generate(c *cli.Context) error {
	generatorOptions := []gen.Option{}
	if c.IsSet(path) {
		generatorOptions = append(generatorOptions, gen.WithPath(c.Path(path)))
	}
	if c.IsSet(force) {
		generatorOptions = append(generatorOptions, gen.WithForce(c.Bool(force)))
	}
	generator, err := gen.New(generatorOptions...)
	if err != nil {
		return cli.Exit(err, 1)
	}

	dayVal := c.Int(day)
	yearVal := c.Int(year)
	fmt.Printf("generating files for day %02d, year %d\n", dayVal, yearVal)
	err = generator.Generate(dayVal, yearVal)
	if err != nil {
		return cli.Exit(err, 1)
	}
	fmt.Println("generated files")
	return nil
}

func runStandalone(c *cli.Context) error {
	day := c.Int(day)
	year := c.Int(year)
	useSample := c.Bool(sample)
	path := c.Path(path)

	dayArg := fmt.Sprintf("--day=%d", day)
	useSampleArg := fmt.Sprintf("--sample=%t", useSample)

	mainFile := filepath.Join(path, strconv.Itoa(year), "main.go")
	cmd := exec.Command("go", "run", mainFile, dayArg, useSampleArg)
	out, err := cmd.Output()
	if err != nil {
		return cli.Exit(err, 1)
	}

	fmt.Println("running:", cmd.String())

	fmt.Print(string(out))

	return nil
}

func run(registry Registry) func(*cli.Context) error {
	return func(c *cli.Context) error {
		day := c.Int(day)
		useSample := c.Bool(sample)

		d, ok := registry[fmt.Sprintf("day%02d", day)]
		if !ok {
			return cli.Exit(fmt.Sprintf("unregistered day requested: %d", day), 1)
		}

		fmt.Printf("solution for day %d part one: %d\n", day, d.PartOne(useSample))
		fmt.Printf("solution for day %d part two: %d\n", day, d.PartTwo(useSample))

		return nil
	}
}
