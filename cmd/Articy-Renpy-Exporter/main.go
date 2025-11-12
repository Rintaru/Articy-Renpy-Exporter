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

type _Character_package_json struct {
	Objects []struct {
		Type       string `json:"Type"`
		Properties struct {
			DisplayName  string `json:"DisplayName"`
			Image_path   string `json:"Image_path"`
			PreviewImage struct {
				Asset string `json:"Asset"`
			} `json:"PreviewImage"`
		} `json:"Properties"`
	} `json:"Objects"`
}
type _Image_asset_package_json struct {
	Objects []struct {
		Type       string `json:"Type"`
		AssetRef   string `json:"AssetRef"`
		Properties struct {
			TechnicalName string `json:"TechnicalName"`
			Id            string `json:"Id"`
		} `json:"Properties"`
	} `json:"Objects"`
}

type _Package_json struct {
	Objects []struct {
		Type     string `json:"Type"`
		raw_data []byte
	} `json:"Objects"`
}

// func ExtractImageAssets(package_manifest map[string]string) (_Image_asset_package_json, error) {

// }

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

func ExtractCharacterDefinitions(package_manifest map[string]string) (Character, error) {
	data, err := os.ReadFile(package_manifest["Character_Exports"])
	if err != nil {
		return Character{}, err
	}
	var _character _Character_package_json
	if err := json.Unmarshal(data, &_character); err != nil {
		return Character{}, err
	}
	for _, object := range _character.Objects {
		if object.Type != "Asset" {
			continue
		}
		fmt.Println(object.Properties.DisplayName, object.Properties.PreviewImage.Asset)
	}
	// _character.Objects
	return Character{}, err

}

// map Object IDs to Object TechnicalNames
func IdToTechnicalName(heirarchy_data *Heirarchy_json) map[string]string {
	output := make(map[string]string, 0)
	var queue []Heirarchy_json

	output[heirarchy_data.Id] = heirarchy_data.TechnicalName

	if heirarchy_data.Children == nil {
		return output
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
	_, err = ExtractCharacterDefinitions(package_map)
	if err != nil {
		fmt.Println("error extracting characters JSON:", err)
		return
	}
}
