package main

import (
	"testing"
	"strings"
	"path/filepath"
	"os"
	"os/exec"
	"bytes"
)

func TestGoldenFiles(t *testing.T) {
	testDir := "testdata/"
	files, err := filepath.Glob(filepath.Join(testDir, "*.pd"))
	if err != nil {
		t.Fatalf("failed to read test files: %v", err)
	}

	for _, inputfile := range files {
		name := strings.TrimSuffix(filepath.Base(inputfile), ".pd")
		wantPath := filepath.Join(testDir, name+".stdout.golden")
		binaryfile := filepath.Join(testDir, name+".pdb")

		t.Run(name, func(t *testing.T) {
			// compile piled into .pdb
			cmd := exec.Command("./piled.exe", "compile", inputfile)
			err := cmd.Run()
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &out
			if err != nil {
				t.Fatalf("failed to compile piled executable: %v\n%s", err, out.String())
			}

			// run .pdb file
			cmd = exec.Command("./piled.exe", "run", binaryfile)
			out.Reset()

			err = cmd.Run()
			if err != nil {
				t.Fatalf("failed to run piled executable: %v\n%s", err, out.String())
			}

			got := out.String()

			wantBytes, err := os.ReadFile(wantPath)
			want := string(wantBytes)

			if got != want {
				t.Errorf("output does not match golden file\n")
				t.Errorf("===== Got =====\n")
				t.Errorf("%s\n", got)
				t.Errorf("===== Want =====\n")
				t.Errorf("%s\n", want)
			}
		})
	}
}
