package osutils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ListFilesInDirectory(path string, depth int, ext ...string) ([]string, error) {
	if depth < 0 {
		return nil, fmt.Errorf("depth cannot be negative")
	}
	var files []string
	err := filepath.Walk(path, func(path string, item os.FileInfo, err error) error {
		if err != nil {
			if os.IsPermission(err) {
				log.Printf("Permission denied: %v", err)
				return filepath.SkipDir
			} else {
				log.Printf("Error accessing the file %v: %v", path, err)
				return nil
			}
		}
		if item.IsDir() {
			return nil
		}
		d := getDepth(path)
		if d > depth {
			return filepath.SkipDir
		}
		if len(ext) > 0 {
			for _, e := range ext {
				if filepath.Ext(path) == e {
					files = append(files, path)
					break
				}
			}
		} else {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Printf("Error walking the path %v: %v", path, err)
		return nil, fmt.Errorf("error walking the path %v: %w", path, err)
	}
	return files, nil
}

func getDepth(path string) int {
	depth := 0
	fmt.Println("Path:", path)
	if path != "." {
		depth = len(strings.Split(path, string(filepath.Separator)))
	}
	return depth
}
