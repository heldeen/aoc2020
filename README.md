# Advent of Code 2020

[![](https://github.com/heldeen/aoc2020/workflows/CI/badge.svg)](https://github.com/heldeen/aoc2020/actions) [![Coverage Status](https://coveralls.io/repos/github/heldeen/aoc2020/badge.svg?branch=master)](https://coveralls.io/github/heldeen/aoc2020?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/heldeen/aoc2020)](https://goreportcard.com/report/github.com/heldeen/aoc2020) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](./LICENSE)

Solutions for the 2020 Advent of Code

## Building

This project makes use of Go 1.15.

```bash
go mod download
go test ./...
```

## Running the Solutions

To run a solution, use the problem name followed by the path to an input file.

For example, to run problem `2a`:

```bash
$ go run ./main.go 2 a -i ./challenge/day2/input.txt
Answer: 9633
Took 999.4µs
```

## Adding New Solutions

A generator program is included in `gen/problem.go` that makes templates for each day. For
example, `go run gen/problem.go 9 a` will generate the following files:

* `challenge/cmd/importDay9.go`: A "glue" file combining commands for both of the day's problems to simplify wiring up subcommands
* `challenge/day9/a.go`: The main problem implementation, containing a cobra command `a` and the implementation `func A(*challenge.Input) int`
* `challenge/day9/a_test.go`: A basic test template
* `challenge/day9/b.go`: The main problem implementation, containing a cobra command `b` and the implementation `func B(*challenge.Input) int`
* `challenge/day9/b_test.go`: A basic test template
* `challenge/day9/input.txt`: The challenge input

## License

These solutions are licensed under the MIT License.

See [LICENSE](./LICENSE) for details.
