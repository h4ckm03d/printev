[![Go Report Card](https://goreportcard.com/badge/github.com/h4ckm03d/printev)](https://goreportcard.com/report/github.com/h4ckm03d/printev)

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

## How to get

```bash
go get go.lumochift.org/printev
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
