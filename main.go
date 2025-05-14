package main

import (
	"fmt"
	"os"
	osutils "wgrep/os_utils"
)

func main() {
	// Example usage of ListFilesInDirectory
	path, _ := os.UserHomeDir()

	files, err := osutils.ListFilesInDirectory(path, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Files found:", files)
}
