package env_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"go.lumochift.org/printev/env"
)

var sampleCode = `package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println(os.Getenv("MYSQL_HOST"))
}

`

func Test_envtEnv(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		lang   env.Lang
		result []string
	}{
		{
			name:   "2 result",
			input:  []byte(` asd asd as dasdos.Getenv("PORT")adasd os.Getenv("MYSQL_CONFIG")`),
			lang:   env.Go,
			result: []string{"PORT", "MYSQL_CONFIG"},
		},
		{
			name:   "duplicate PORT",
			input:  []byte(` asd asd as dasdos.Getenv("PORT")adasd os.Getenv("PORT")`),
			lang:   env.Go,
			result: []string{"PORT"},
		},
		{
			name: "ruby test",
			input: []byte(`conf[:service] = ENV['PROTOR_SERVICENAME']
  										conf[:host] = ENV['PROTOR_HOST']`),
			lang:   env.Ruby,
			result: []string{"PROTOR_SERVICENAME", "PROTOR_HOST"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := env.GetEnv(tt.input, tt.lang)
			assert.EqualValues(t, result, tt.result)
		})
	}
}

func TestWalkinenvv(t *testing.T) {
	targetDir, err := ioutil.TempDir("", "template")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	targetPath := filepath.Join(targetDir, "main.go")
	f, err := os.Create(targetPath)
	if err != nil {
		t.Error("fail create sample service")
	}
	_, err = f.Write([]byte(sampleCode))
	if err != nil {
		t.Log(err)
	}

	defer os.RemoveAll(targetDir)

	expectedEnv := []string{"MYSQL_HOST"}
	environments := env.FindEnv(targetDir)
	sort.Strings(expectedEnv)
	sort.Strings(environments)
	assert.EqualValues(t, environments, expectedEnv)
}

func TestGetEnvStruct(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := map[string]struct {
		args        args
		wantRecords []string
	}{
		// TODO: Add test cases.
		"mysql_user": {
			args:        args{[]byte(`env:"MYSQL_USERNAME,required"`)},
			wantRecords: []string{"MYSQL_USERNAME"},
		},
		"mysql_charset": {
			args:        args{[]byte(`env:"MYSQL_CHARSET,required"`)},
			wantRecords: []string{"MYSQL_CHARSET"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			gotRecords := env.GetEnvStruct(tt.args.b)
			if diff := cmp.Diff(tt.wantRecords, gotRecords); diff != "" {
				t.Errorf("GetEnvStruct() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
