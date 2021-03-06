package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/lumochift/printev"
	"github.com/urfave/cli/v2"
)

var version = "1.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "printev"
	app.Description = "Go Env Printer"
	app.Usage = "Generate env variable from given codes"
	app.Copyright = "Moch Lutfi <kapten.lufi@gmail.com>™ © 2020"
	app.Version = version
	env := Env{}
	app.Flags = env.Flags()
	app.Action = env.Action

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

// Env is an env command implementation
type Env struct {
	verbose     bool
	writeToFile bool
	output      string
	source      string
}

// Action is a function command executor
func (e *Env) Action(c *cli.Context) error {
	e.output = c.String("output")
	e.verbose = !c.Bool("mute")
	e.writeToFile = c.Bool("write")
	e.source = c.String("source")
	e.opEnv()

	return nil
}

// Flags is a function to return all registered flag
func (e *Env) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "source",
			Aliases: []string{"s"},
			Usage:   "[Optional] Target source code. (default: current dir)",
		},
		&cli.BoolFlag{
			Name:    "mute",
			Aliases: []string{"m"},
			Usage:   "[Optional] Hide preview and log.",
		},
		&cli.BoolFlag{
			Name:    "write",
			Aliases: []string{"w"},
			Usage:   "[Optional] Write environment variables found.",
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "[Optional] Output location of generated env files, by default write to printev.sample",
			Value:   "printev.sample",
		},
	}
}

func (e *Env) opEnv() {
	source := e.source
	if source == "" {
		source, _ = os.Getwd()
	}

	envs := printev.FindEnv(source)

	e.showEnv(envs)

	if !e.writeToFile {
		return
	}
	// check if target exist
	if _, err := os.Stat(e.output); !os.IsNotExist(err) {
		if !e.confirmOverwrite(e.output) {
			return
		}
	}

	if err := e.writeEnvFile(envs, e.output); err != nil {
		e.Println("\033[0;31merror:", err, "\033[0m")
		return
	}

	e.Printf("\033[0;32m%s created \033[0m\n", e.output)
}

func (e *Env) showEnv(envs []string) {
	if len(envs) == 0 {
		return
	}

	e.Println("List env variable:")
	sort.Strings(envs)
	e.Println(strings.Join(envs, "\n"))
}

func (e *Env) writeEnvFile(envs []string, output string) error {
	var buff bytes.Buffer

	sort.Strings(envs)
	buff.WriteString(strings.Join(envs, "=\n") + "=")

	return ioutil.WriteFile(output, buff.Bytes(), 0644)
}

func (e *Env) confirmOverwrite(path string) bool {
	reader := bufio.NewReader(os.Stdin)

	e.Printf("\033[0;33mAre you sure overwrite %s (y/n)?\033[0m", path)

	confirmation, _ := reader.ReadString('\n')

	return strings.TrimSuffix(strings.ToLower(confirmation), "\n") == "y"
}

// Println is proxy fmt.Printf with verbosity condition.
func (e *Env) Printf(format string, a ...interface{}) {
	if e.verbose {
		fmt.Printf(format, a...)
	}
}

// Println is proxy fmt.Println with verbosity condition.
func (e *Env) Println(a ...interface{}) {
	if e.verbose {
		fmt.Println(a...)
	}
}
