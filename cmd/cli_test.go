package cmd

import (
	"bytes"
	"log"
	"os/exec"
	"testing"
)

var suite = map[string][]struct {
	args     []string
	stdin    string
	stdout   string
	stderr   string
	exitCode int
}{
	"echo": {
		{
			args:   []string{"-n", "Hello"},
			stdout: "Hello",
		},
	},
	"rev": {
		{
			stdin:  "Hello",
			stdout: "olleH",
		},
	},
	"suddendeath": {
		{
			stdin: `突然の死`,
			stdout: `＿人人人人人人＿
＞　突然の死　＜
￣ＹＹＹＹＹＹ￣`,
		}, {
			stdin: `突然



ののののののののの死`,
			stdout: `＿人人人人人人人人人人人人＿
＞　突然                　＜
＞　                    　＜
＞　                    　＜
＞　                    　＜
＞　ののののののののの死　＜
￣ＹＹＹＹＹＹＹＹＹＹＹＹ￣`,
		},
	},
}

func init() {
	for name, _ := range suite {
		name := name
		fp, err := exec.LookPath(name)
		if err != nil {
			cmd := exec.Command("go", "install", "-v", "./"+name)
			err = cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
			fp, err = exec.LookPath(name)
		}
		log.Printf("executable file %s", fp)
	}
}

func Test_CLI(t *testing.T) {
	for name, tests := range suite {
		name := name
		tests := tests
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for _, tt := range tests {
				tt := tt
				stdin := bytes.NewBufferString(tt.stdin)
				stdout := bytes.NewBufferString("")
				stderr := bytes.NewBufferString("")
				cmd := exec.Command(name, tt.args...)
				cmd.Stdin = stdin
				cmd.Stdout = stdout
				cmd.Stderr = stderr
				if err := cmd.Start(); err != nil {
					t.Errorf("got error %v", err)
				}
				if err := cmd.Wait(); err != nil {
					t.Errorf("got error %v", err)
				}
				out1 := stdout.String()
				out2 := stderr.String()
				switch {
				case out1 != tt.stdout:
					t.Errorf("want '%v' but got '%v'", tt.stdout, out1)
				case out2 != tt.stderr:
					t.Errorf("want '%v' but got '%v'", tt.stdout, out2)
				}
			}
		})
	}
}
