package main

import (
	"log"
	"os"

	"github.com/k-nox/aoc/cli"
)

func main() {
	app := cli.StandaloneApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
