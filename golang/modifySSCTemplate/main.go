package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Specify the directory containing the .dat files
	dir := "C:\\Sigtest\\data"

	// Get all files with the ".dat" extension in the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".dat") {
			// Full path to the source file
			srcPath := filepath.Join(dir, file.Name())

			// Create a new file name with "_modified.dat" suffix
			modifiedFileName := strings.TrimSuffix(file.Name(), ".dat") + "_SSC.dat"
			dstPath := filepath.Join(dir, modifiedFileName)

			// Copy the source file to the modified file
			err := copyAndModifyFile(srcPath, dstPath)
			if err != nil {
				fmt.Println("Error processing file:", file.Name(), err)
			} else {
				fmt.Println("Successfully processed file:", modifiedFileName)
			}
		}
	}
}

// Function to copy a file and modify the last line
func copyAndModifyFile(srcPath, dstPath string) error {
	// Open source file
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Open destination file
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Read all lines from the source file
	var lines []string
	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Modify the last line (if the file isn't empty)
	//if len(lines) > 0 {
	//	lines[len(lines)-1] = lines[len(lines)-1] + " - modified"
	//}

	// Write all lines to the destination file
	for _, line := range lines {
		if strings.Contains(line, "SSCRemoval = 0") {
			line = strings.Replace(line, "SSCRemoval = 0", "SSCRemoval = 1", 1)
		}
		_, err := dstFile.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
