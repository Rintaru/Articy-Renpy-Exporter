package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Manifest struct {
	Packages []struct {
		Name  string `json:"Name"`
		Files struct {
			Objects struct {
				FileName string `json:"FileName"`
			} `json:"Objects"`
		} `json:"Files"`
	} `json:"Packages"`
}

func ExtractPackageMap(filename string) (map[string]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, pkg := range manifest.Packages {
		result[pkg.Name] = pkg.Files.Objects.FileName
	}
	return result, nil
}

func main() {
	package_map, err := ExtractPackageMap("/mnt/c/GIT_REPOS/Visual_Novels/Practice_Export/Organized_Export/manifest.json")
	if err != nil {
		return
	}
	for key, value := range package_map {
		fmt.Println(key, value)
	}

}
