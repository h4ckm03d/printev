package printev_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/lumochift/printev"
	"github.com/stretchr/testify/assert"
)

func Test_envtEnv(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		lang   printev.Lang
		result []string
	}{
		{
			name:   "2 result",
			input:  []byte(` asd asd as dasdos.Getenv("PORT")adasd os.Getenv("MYSQL_CONFIG")`),
			lang:   printev.Go,
			result: []string{"PORT", "MYSQL_CONFIG"},
		},
		{
			name:   "duplicate PORT",
			input:  []byte(` asd asd as dasdos.Getenv("PORT")adasd os.Getenv("PORT")`),
			lang:   printev.Go,
			result: []string{"PORT"},
		},
		{
			name: "ruby test",
			input: []byte(`conf[:service] = ENV['PROTOR_SERVICENAME']
  										conf[:host] = ENV['PROTOR_HOST']`),
			lang:   printev.Ruby,
			result: []string{"PROTOR_SERVICENAME", "PROTOR_HOST"},
		},
		{
			name:   "node test",
			input:  []byte(`console.log('The value of PORT is:', process.env.PORT);`),
			lang:   printev.Node,
			result: []string{"PORT"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := printev.GetEnv(tt.input, tt.lang)
			assert.EqualValues(t, result, tt.result)
		})
	}
}

func TestWalkinenvv(t *testing.T) {
	targetDir, err := ioutil.TempDir("", "go")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}

	targetPath := filepath.Join(targetDir, "config.go")
	f, err := os.Create(targetPath)
	if err != nil {
		t.Error("fail create sample service")
	}

	_, err = f.Write([]byte(`package main

	import (
		"os"
		"fmt"
	)
	
	func main() {
		fmt.Println(os.Getenv("MYSQL_HOST"))
	}
	
	`))
	if err != nil {
		t.Log(err)
	}

	defer os.RemoveAll(targetDir)

	expectedEnv := []string{"MYSQL_HOST"}
	environments := printev.FindEnv(targetDir)
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
			gotRecords := printev.GetEnvStruct(tt.args.b)
			if diff := cmp.Diff(tt.wantRecords, gotRecords); diff != "" {
				t.Errorf("GetEnvStruct() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
