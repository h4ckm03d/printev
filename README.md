[![GoDoc](https://godoc.org/github.com/lumochift/printev?status.svg)](https://godoc.org/github.com/lumochift/printev)
[![Build Status](https://github.com/lumochift/printev/workflows/Go%20workflow/badge.svg)](https://github.com/lumochift/printev/actions)
[![codecov](https://codecov.io/gh/lumochift/printev/branch/master/graph/badge.svg)](https://codecov.io/gh/lumochift/printev)
[![Go Report Card](https://goreportcard.com/badge/github.com/lumochift/printev)](https://goreportcard.com/report/github.com/lumochift/printev)

# printev

## Features

- Read env var from go and ruby codes
- Output served on the sorted format

```bash
NAME:
   printev - Generate env variable from given codes

USAGE:
   printev [global options] command [command options] [arguments...]

VERSION:
   1.0.1

DESCRIPTION:
   Go Env Printer

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --source value, -s value  [Optional] Target source code. (default: current dir)
   --mute, -m                [Optional] Hide preview and log. (default: false)
   --write, -w               [Optional] Write environment variables found. (default: false)
   --output value, -o value  [Optional] Output location of generated env files, by default write to printev.sample (default: "printev.sample")
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)

COPYRIGHT:
   Lumochift™ © 2020
```

## Installation

```bash
go install github.com/lumochift/printev/cmd/printev
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
