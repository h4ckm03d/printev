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

	"github.com/lumochift/printev/env"
	"github.com/urfave/cli"
)

var (
	Version string
)

func main() {
	app := cli.NewApp()
	app.Name = "printev"
	app.Description = "Go Env Printer"
	app.Usage = "Generate env variable from given codes"
	app.Copyright = "Lumochift™ © 2020"
	app.Version = Version
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
}

// Action is a function command executor
func (e *Env) Action(c *cli.Context) error {
	output := c.String("output-file")
	if output == "" {
		output = "env.sample"
	}
	e.verbose = !c.Bool("silent")
	e.writeToFile = c.Bool("write")

	e.opEnv()
	return nil
}

// Flags is a function to return all registered flag
func (e *Env) Flags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "mute, m",
			Usage: "[Optional] Hide preview and log.",
		},
		cli.BoolFlag{
			Name:  "write, w",
			Usage: "[Optional] Write environment variables found.",
		},
		cli.StringFlag{
			Name:  "output-file, o",
			Usage: "[Optional] Output location of generated env files, by default write to env.sample",
		},
	}
}

func (e *Env) opEnv() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	envs := env.FindEnv(dir)

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
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%s created \n", e.output)
}

func (e *Env) showEnv(envs []string) {
	if len(envs) == 0 {
		return
	}
	fmt.Println("List env variable:")
	sort.Strings(envs)
	fmt.Println(strings.Join(envs, "\n"))
}

func (e *Env) writeEnvFile(envs []string, output string) error {
	var buff bytes.Buffer
	sort.Strings(envs)
	buff.WriteString(strings.Join(envs, "=\n") + "=")
	return ioutil.WriteFile(output, buff.Bytes(), 0644)
}

func (e *Env) confirmOverwrite(path string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Are you sure overwrite %s (y/n)?", path)
	confirmation, _ := reader.ReadString('\n')
	return strings.TrimSuffix(strings.ToLower(confirmation), "\n") == "y"
}
