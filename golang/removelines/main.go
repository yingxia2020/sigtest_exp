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
	// Define the directory path
	dir := "C:\\Users\\yingxia\\sandbox\\applications.validation.compliance.io-standards.sigtest-phoenix-code-gen6\\Templates\\PCIe\\6_0\\Base_TX_w_52UIJitterPattn"

	// Get a list of .dat files in the directory
	files, err := filepath.Glob(filepath.Join(dir, "*.dat"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Define the pattern to search for
	pattern := "StandardDevs = 2" //"[NoiseFloor]"  //"BinSize ="
	line := "StandardDevs = 3" //"BinSize = 0.5E-13"
	/*
	constraint := `[CONSTRAINT]
; force the sum of C-2 C-1 C0 and C+1 to be 1
SumOneConstraint = 1`
	*/

	// Loop through each .dat file
	for _, file := range files {
		// Open the file for reading
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("Error opening file %s: %v\n", file, err)
			continue
		}
		defer f.Close()

		// Read the contents of the file
		scanner := bufio.NewScanner(f)
		var lines []string
		var results []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			continue
		}

		// Find and remove the pattern and the 10 lines after it
		for i := 0; i < len(lines); i++ {
			/* remove 10 lines and the pattern line block
			if strings.Contains(lines[i], pattern) {
				lines = append(lines[:i], lines[i+11:]...)
				break
			} */
			// block to replace one line in the dat file
			if strings.Contains(lines[i], pattern) {
				results = append(results, lines[:i]...)
				results = append(results, line)
				results = append(results, lines[i+1:]...)
				break
			}

			/* blocks to remove the pattern and all the lines after it
			if strings.Contains(lines[i], pattern) {
				results = append(results, lines[:i]...)
				break
			}
			 */
		}

		// block to append several lines at the end of the file
		//results = append(results, lines[:]...)
		//results = append(results, constraint)

		// Write the modified content back to the file
		if err := ioutil.WriteFile(file, []byte(strings.Join(results, "\n") + "\n"), 0644); err != nil {
			fmt.Printf("Error writing file %s: %v\n", file, err)
			continue
		}
	}
}