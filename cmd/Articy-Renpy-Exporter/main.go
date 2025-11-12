package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type _Manifest_json struct {
	Packages []struct {
		Name  string `json:"Name"`
		Files struct {
			Objects struct {
				FileName string `json:"FileName"`
			} `json:"Objects"`
		} `json:"Files"`
	} `json:"Packages"`
}

type Heirarchy_json struct {
	Id            string            `json:"Id"`
	TechnicalName string            `json:"TechnicalName"`
	Type          string            `json:"Type"`
	Children      *[]Heirarchy_json `Json:"Children,omitempty"`
}

type Character struct {
	Name       string
	Image_path string
}

type Character_package_json struct {
	Properties struct {
		DisplayName  string `json:"DisplayName"`
		PreviewImage struct {
			Asset string `json:"Asset"`
		} `json:"PreviewImage"`
	} `json:"Properties"`
}
type Image_asset_package_json struct {
	AssetRef   string `json:"AssetRef"`
	Properties struct {
		TechnicalName string `json:"TechnicalName"`
		Id            string `json:"Id"`
	} `json:"Properties"`
}

// extract manifest.json and map package name to the corresponding file path
func ExtractPackageMap(top_level_path string, filename string) (map[string]string, error) {

	data, err := os.ReadFile(top_level_path + filename)
	if err != nil {
		return nil, err
	}

	var manifest _Manifest_json
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, pkg := range manifest.Packages {
		result[pkg.Name] = top_level_path + pkg.Files.Objects.FileName
	}
	return result, nil
}

// extract characters and image assets into a respective list
func ExtractCharacterPackages(package_manifest map[string]string) ([]Image_asset_package_json, []Character_package_json, error) {
	data, err := os.ReadFile(package_manifest["Character_Exports"])
	if err != nil {
		return []Image_asset_package_json{}, []Character_package_json{}, err
	}

	var raw struct {
		Objects []json.RawMessage `json:"Objects"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		fmt.Println("error parsing JSON:", err)
		return []Image_asset_package_json{}, []Character_package_json{}, err
	}

	var type_only struct {
		Type string `json:"Type"`
	}

	var asset_packages []Image_asset_package_json
	var character_packages []Character_package_json

	for _, raw_item := range raw.Objects {
		if err := json.Unmarshal(raw_item, &type_only); err != nil {
			fmt.Println("error parsing JSON:", err)
			return []Image_asset_package_json{}, []Character_package_json{}, err
		}

		switch type_only.Type {
		case "Entity":
			var temp Character_package_json
			json.Unmarshal(raw_item, &temp)
			character_packages = append(character_packages, temp)
		case "Asset":
			var temp Image_asset_package_json
			json.Unmarshal(raw_item, &temp)
			asset_packages = append(asset_packages, temp)
		}
	}

	return asset_packages, character_packages, err

}

// map Object IDs to Object TechnicalNames
func (heirarchy Heirarchy_json) IdToTechnicalName() map[string]string {
	output := make(map[string]string, 0)
	var queue []Heirarchy_json

	output[heirarchy.Id] = heirarchy.TechnicalName

	if heirarchy.Children == nil {
		return output
	}
	queue = append(queue, *heirarchy.Children...)
	for len(queue) > 0 {
		h := queue[0]
		queue = queue[1:]
		output[h.Id] = h.TechnicalName

		if h.Children != nil {
			queue = append(queue, *h.Children...)
		}

	}

	return output

}
func main() {
	top_level_path := "/mnt/c/GIT_REPOS/Visual_Novels/Practice_Export/Organized_Export/"
	package_map, err := ExtractPackageMap(top_level_path, "manifest.json")
	if err != nil {
		return
	}

	data, err := os.ReadFile(top_level_path + "hierarchy.json")
	if err != nil {
		fmt.Println("error loading JSON:", err)
		return
	}

	var heirarchy_data Heirarchy_json
	if err := json.Unmarshal(data, &heirarchy_data); err != nil {
		fmt.Println("error parsing JSON:", err)
		return
	}
	_, _, err = ExtractCharacterPackages(package_map)
	if err != nil {
		fmt.Println("error extracting characters JSON:", err)
		return
	}
}
