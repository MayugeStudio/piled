package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestGoldenFiles(t *testing.T) {
	testDir := "testdata/"
	files, err := filepath.Glob(filepath.Join(testDir, "*.pd"))
	if err != nil {
		t.Fatalf("failed to read test files: %v", err)
	}

	for _, inputfile := range files {
		name := strings.TrimSuffix(filepath.Base(inputfile), ".pd")
		wantPath := filepath.Join(testDir, name+".stdout")
		binaryfile := filepath.Join(testDir, name+".pdb")

		t.Run(name, func(t *testing.T) {
			// compile piled into .pdb
			cmd := exec.Command("./piled.exe", "compile", inputfile)
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &out

			err := cmd.Run()
			if err != nil {
				t.Fatalf("failed to compile piled program: %v\n%s", err, out.String())
			}

			// run .pdb file
			cmd = exec.Command("./piled.exe", "run", binaryfile)
			cmd.Stdout = &out
			cmd.Stderr = &out
			out.Reset()

			err = cmd.Run()
			if err != nil {
				t.Fatalf("failed to run piled executable: %v\n%s", err, out.String())
			}

			got := strings.ReplaceAll(out.String(), "\r\n", "\n")

			wantBytes, err := os.ReadFile(wantPath)
			if err != nil {
				t.Fatalf("failed to read expected-output-file: %v", err)
			}

			want := strings.ReplaceAll(string(wantBytes), "\r\n", "\n")

			if got != want {
				t.Errorf("output does not match golden file\n")
				t.Errorf("===== Got =====\n")
				t.Errorf("%#v\n", got)
				t.Errorf("len: %d\n", len(got))
				t.Errorf("===== Want =====\n")
				t.Errorf("%#v\n", want)
				t.Errorf("len: %d\n", len(want))
			}
		})
	}
}
