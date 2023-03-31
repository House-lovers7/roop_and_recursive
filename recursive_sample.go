package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func globPaths(targetPath string, maxDepth int, suffix string, depth int) []string {
	var matchedPaths []string

	if depth == maxDepth {
		fmt.Printf("Reached maximum depth of %d at path %s\n", maxDepth, targetPath)
	} else {
		fileInfo, err := os.Stat(targetPath)
		if err != nil {
			panic(err)
		}

		if fileInfo.IsDir() {
			depth++

			fileInfos, err := ioutil.ReadDir(targetPath)
			if err != nil {
				panic(err)
			}

			for _, fileInfo := range fileInfos {
				subPath := filepath.Join(targetPath, fileInfo.Name())
				subPathMatchedPaths := globPaths(subPath, maxDepth, suffix, depth)
				matchedPaths = append(matchedPaths, subPathMatchedPaths...)
			}
		} else if fileInfo.Mode().IsRegular() {
			if filepath.Ext(targetPath) == suffix {
				matchedPaths = append(matchedPaths, targetPath)
			}
		}
	}

	return matchedPaths
}

func main() {
	result := globPaths(".", 2, ".json", 0)
	for _, item := range result {
		fmt.Println(filepath.Join(filepath.Dir(item), fmt.Sprintf("%s.json5", filepath.Base(item[:len(item)-len(filepath.Ext(item))]))))
	}
}
