// Package printev provides functionality to get all env usage
package printev

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kr/fs"
)

// Lang is type holder for programming language type
type Lang byte

const (
	// Go programming language
	Go Lang = iota
	// Ruby programming language
	Ruby
)

var (
	goEnvRgx       = regexp.MustCompile(`os\.Getenv\(\"(.*?)\"\)`)
	goEnvStructRgx = regexp.MustCompile(`env:\"(.*?)\"`)
	rubyEnvRgx     = regexp.MustCompile(`ENV\[['”](.*?)['”]\]`)
)

// FindEnv from path, currently get env from *.go files
func FindEnv(path string) []string {
	records := []string{}
	walker := fs.Walk(path)
	for walker.Step() {
		if err := walker.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		ext := filepath.Ext(walker.Stat().Name())
		if !strings.Contains(".go .go.template .yml .rb", ext) {
			continue
		}
		f, _ := os.Open(walker.Path())

		reader := bufio.NewReader(f)
		content, _ := ioutil.ReadAll(reader)
		lang := getLang(ext)
		envs := GetEnv(content, lang)
		f.Close()
		if strings.HasPrefix(walker.Stat().Name(), "config.go") {
			structEnvs := GetEnvStruct(content)
			records = append(records, structEnvs...)
		}
		if len(envs) > 0 {
			records = append(records, envs...)
		}
	}
	return unique(records)
}

// get language from ext
func getLang(ext string) Lang {
	switch ext {
	case ".go":
	case ".go.template":
		return Go
	case ".yaml":
	case ".rb":
		return Ruby
	}
	//fallback go
	return Go
}

// check unique string
func unique(s []string) []string {
	keys := make(map[string]bool)
	var uniques []string
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniques = append(uniques, entry)
		}
	}
	return uniques
}

// GetEnv from given []bytes value
func GetEnv(b []byte, lang Lang) (records []string) {
	var results [][][]byte
	switch lang {
	case Go:
		results = goEnvRgx.FindAllSubmatch(b, -1)
	case Ruby:
		results = rubyEnvRgx.FindAllSubmatch(b, -1)
	}

	for _, r := range results {
		if len(r) > 1 {
			records = append(records, string(r[1]))
		}
	}
	return unique(records)
}

// GetEnvStruct fetch env from tag struct
func GetEnvStruct(b []byte) (records []string) {
	results := goEnvStructRgx.FindAllSubmatch(b, -1)
	for _, r := range results {
		if len(r) > 1 {
			cleaned := strings.Split(strings.Trim(string(r[1]), " "), ",")
			records = append(records, cleaned[0])
		}
	}
	return unique(records)
}
