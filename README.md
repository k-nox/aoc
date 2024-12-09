# aoc

An opinonated CLI to make life a little bit easier when solving [Advent of Code](https://adventofcode.com/) with Go.
It generates daily files for solutions and can also be used to run the solutions.

## Installation

```sh
$ go install github.com/k-nox/aoc@latest
```

Currently the files generated by this tool depend on the `cli` package in this module, so you'll also need add the module to your dependencies:

```sh
$ go get github.com/k-nox/aoc
```

## Usage

### Generate

The file structure created by this tool looks like this:

```
|-- 2024
|   |-- day01
|   |   |-- partone.go
|   |   |-- parttwo.go
|   |-- day02
|   |   |-- partone.go
|   |   |-- partwo.go
|   |-- main.go
|-- 2025
|   |-- day01
|   |   |-- partone.go
|   |   |-- parttwo.go
|   |-- main.go
|-- input
|   |-- 2024
|   |   |-- day01
|   |   |   |-- input.txt
|   |   |   |-- sample.txt
```

Running `aoc gen` with no flags will assume you want to generate files for the current year & day. For example if it's the Dec 12, 2024, the tool will generate `day12` files under the `2024` folder. You can use the `--day` and `--year` commands to provide a different day & year.

The tool will look for a `go.mod` in your current directory and parse your module name from that to use in the `main.go` template.

If you want the tool to consider a different path your current directory, you can pass that in with a `--path` flag. This will also affect where the files end up being generated.

If the tool detects the requested files are already present, it will fail, unless you choose to pass in a `--force` flag.

The default templates used are in `gen/templates.go`. You can override these in one of two ways:

- a `--partTemplate` or `--mainTemplate` flag with a value of the path to the template file you want to use
- setting `AOC_PART_TEMPLATE` or `AOC_MAIN_TEMPLATE` in your environment to the path of the template file you want to use.

You can see an example of a custom template file being used in [my solutions repo](https://github.com/k-nox/advent-of-code-solutions/blob/main/templates/part.go.tmpl)

Currently the generated input files are empty, but it's planned to add an option to download them from Advent of Code.

### Running Solutions

There are two ways to run your solutions:

1. Run them directly using the generated `main.go` files: `go run 2024/main.go`
2. Use aoc to run it: `aoc run`

`aoc run` is just a convenience wrapper that executes `go run <year>/main.go`.

Both options allow you to pass some flags:

- `--sample` will use the `input/<year>/<day>/sample.txt` file instead of the `input.txt` file.
- `--day=<day>` to choose a day other than today to run

`aoc run` also accepts `--year` to choose the year to run.
