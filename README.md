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

```
NAME:
   aoc gen - Generate files for specified day

USAGE:
   aoc gen [command options]

OPTIONS:
   --day value, -d value   Specify day (default: today)
   --year value, -y value  specify year (default: current year)
   --force, -f             Force generation - may overwrite existing files (default: false)
   --path value, -p value  Path to your advent of code solutions directory (default: current working directory)
   --partTemplate value    Use custom template for part files [or set $AOC_PART_TEMPLATE in your env]
   --mainTemplate value    Use custom template for main file [or set $AOC_MAIN_TEMPLATE in your env]
   --session value         Advent of Code session string, used for downloading input files [or set $AOC_SESSION in your env]
   --help, -h              show help
```

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

You can see an example of a custom template file being used in [my solutions repo](https://github.com/k-nox/advent-of-code-solutions/blob/main/templates/part.go.tmpl)

### Running Solutions

```
NAME:
   aoc run - Run solutions for specified day

USAGE:
   aoc run [command options]

OPTIONS:
   --day value, -d value   Specify day (default: today)
   --year value, -y value  specify year (default: current year)
   --path value, -p value  Path to your advent of code solutions directory (default: current working directory)
   --sample, -s            Use sample input (default: false)
   --help, -h              show help
```

`aoc run` is just a convenience wrapper that executes `go run <year>/main.go`.
If you want to run your main file directly, it will also accept the `--sample` and `--day` flags.
