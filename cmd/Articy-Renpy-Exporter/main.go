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

type Heirarchy struct {
	Id            string       `json:"Id"`
	TechnicalName string       `json:"TechnicalName"`
	Type          string       `json:"Type"`
	Children      *[]Heirarchy `Json:"Children,omitempty"`
}

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

func IdToTechnicalName(heirarchy_data *Heirarchy) (map[string]string, error) {
	output := make(map[string]string, 0)
	var queue []Heirarchy

	output[heirarchy_data.Id] = heirarchy_data.TechnicalName

	if heirarchy_data.Children == nil {
		return output, nil
	}
	queue = append(queue, *heirarchy_data.Children...)
	for len(queue) > 0 {
		h := queue[0]
		queue = queue[1:]
		output[h.Id] = h.TechnicalName

		if h.Children != nil {
			queue = append(queue, *h.Children...)
		}

	}

	return output, nil

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

	data, err := os.ReadFile(top_level_path + "hierarchy.json")
	if err != nil {
		fmt.Println("error loading JSON:", err)
		return
	}

	var heirarchy_data Heirarchy
	if err := json.Unmarshal(data, &heirarchy_data); err != nil {
		fmt.Println("error parsing JSON:", err)
		return
	}

	id_map, _ := IdToTechnicalName(&heirarchy_data)

	fmt.Println(len(id_map))

	for key, value := range id_map {
		fmt.Println(key, value)
	}

}
