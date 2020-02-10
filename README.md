[![GoDoc](https://godoc.org/github.com/lumochift/printev?status.svg)](https://godoc.org/github.com/lumochift/printev)
[![Build Status](https://github.com/lumochift/printev/workflows/Go%20workflow/badge.svg)](https://github.com/lumochift/printev/actions)
[![Coverage Status](https://badgen.net/codecov/c/github/lumochift/printev)](https://codecov.io/gh/lumochift/printev)
[![Go Report Card](https://goreportcard.com/badge/github.com/lumochift/printev)](https://goreportcard.com/report/github.com/lumochift/printev)

# printev

## Features

- Read env var from go and ruby codes
- Output served on sorted format

```bash
NAME:
   printev - Generate env variable from given codes

USAGE:
   printev [global options] command [command options] [arguments...]

DESCRIPTION:
   Go Env Printer

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --source value, -s value  [Optional] Target source code
   --mute, -m                [Optional] Hide preview and log.
   --write, -w               [Optional] Write environment variables found.
   --output value, -o value  [Optional] Output location of generated env files, by default write to env.sample
   --help, -h                show help
   --version, -v             print the version

COPYRIGHT:
   Lumochift™ © 2020
```

## Usage

```bash
go get github.com/lumochift/printev/cli/printev
```

## Sample run

### Source code example

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(os.Getenv("TEST_ENV_1"))
    fmt.Println(os.Getenv("TEST_ENV_2"))
}
```

### Output

```bash
➜   printev -s testdata
List env variable:
TEST_ENV_1
TEST_ENV_2
```
