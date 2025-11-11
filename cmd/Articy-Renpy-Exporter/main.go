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

// type Package_Objects struct {
// 	Objects []struct {
// 		Type       string `json:"Entity"`
// 		Properties struct {
// 		}
// 	}
// }

type Character struct {
	Name       string
	Image_path string
}

func ExtractPackageMap(top_level_path string, filename string) (map[string]string, error) {

	data, err := os.ReadFile(top_level_path + filename)
	if err != nil {
		return nil, err
	}

	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, pkg := range manifest.Packages {
		result[pkg.Name] = top_level_path + pkg.Files.Objects.FileName
	}
	return result, nil
}

// func ExtractCharacterDefinitions(package_manifest map[string]string) (Character, error) {
// 	data, err := os.ReadFile(package_manifest["Character_Exoorts"])
// 	if err != nil {
// 		return Character{}, err
// 	}
// 	data
// }

func TechincalNametoID(top_level_path string, filename string) (map[string]string, error) {

}
func main() {
	top_level_path := "/mnt/c/GIT_REPOS/Visual_Novels/Practice_Export/Organized_Export/"
	package_map, err := ExtractPackageMap(top_level_path, "manifest.json")
	if err != nil {
		return
	}
	for key, value := range package_map {
		fmt.Println(key, value)
	}
	// ExtractCharacterDefinitions(package_map)

}
