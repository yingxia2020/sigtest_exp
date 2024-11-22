package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os/exec"
	"path/filepath"
)

func main() {
	ext := ".bin"

	// Define the executable and its parameters
	executable := "C:\\Users\\yingxia\\sandbox\\applications.validation.compliance.io-standards.sigtest-phoenix-code-gen6\\SigTestGUI\\bin\\Debug\\SigTest.exe"

	// Specify the root directory
	rootDir := "C:\\Users\\yingxia\\Downloads\\SigTest\\wave_forms\\gen6_sj_cal\\100MHz"

	// Walk through the directory with a depth limit of 4
	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories beyond the 2-level depth
		if d.IsDir() {
			relativePath, _ := filepath.Rel(rootDir, path)
			depth := len(filepath.SplitList(relativePath))
			if depth > 2 {
				return filepath.SkipDir
			}
			return nil
		}

		// Check for files with ".bin" extension
		if filepath.Ext(path) == ext {
			//fmt.Println(path)
			// run this waveform
			params := []string{"PCIe", "6_0", "CEM_SjCal", "/t", "Sj_Cal", "/s", path,
				"/o", rootDir + "result.csv"}

			// Create the command
			cmd := exec.Command(executable, params...)

			// Create a buffer to capture the standard output and standard error
			var outBuffer bytes.Buffer
			var errBuffer bytes.Buffer
			cmd.Stdout = &outBuffer
			cmd.Stderr = &errBuffer

			// Run the command
			fmt.Println("Sj Cal flow test for file: " + path)
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				fmt.Printf("Standard Error: %s\n", errBuffer.String())
				return nil
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

