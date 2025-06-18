package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	startPath := "../qkonnectSoftphone"
	count := 0
	traverseDir(startPath, &count)
	fmt.Println("Total matched files:", count)
}

func traverseDir(path string, count *int) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", path, err)
		return
	}

	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())

		if entry.IsDir() {
			traverseDir(fullPath, count) // recurse with same counter
		} else {
			content, err := os.ReadFile(fullPath)
			if err != nil {
				fmt.Println("Error reading file:", fullPath)
				continue
			}

			if strings.Contains(string(content), "sessionManager") {
				fmt.Println("Match found in:", fullPath)
				(*count)++
			}
		}
	}
}
