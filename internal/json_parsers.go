package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type Manifest_json struct {
	Packages []struct {
		Name  string `json:"Name"`
		Files struct {
			Objects struct {
				FileName string `json:"FileName"`
			} `json:"Objects"`
			Texts struct {
				FileName string `json:"FileName"`
			} `json:"Texts"`
		} `json:"Files"`
	} `json:"Packages"`
}

type Heirarchy_json struct {
	Id            string            `json:"Id"`
	TechnicalName string            `json:"TechnicalName"`
	Type          string            `json:"Type"`
	Children      *[]Heirarchy_json `Json:"Children,omitempty"`
}

type character_object struct {
	Properties struct {
		TechnicalName string `json:"TechnicalName"`
		DisplayName   string `json:"DisplayName"`
		PreviewImage  struct {
			Asset string `json:"Asset"`
		} `json:"PreviewImage"`
	} `json:"Properties"`
}
type asset_object struct {
	AssetRef   string `json:"AssetRef"`
	Properties struct {
		TechnicalName string `json:"TechnicalName"`
		Id            string `json:"Id"`
	} `json:"Properties"`
}

type Asset_json = map[string]asset_object
type Character_json = map[string]character_object

func (m Manifest_json) from_file(top_level_path string, filename string) (*Manifest_json, error) {
	data, err := os.ReadFile(top_level_path + filename)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// extract manifest.json and map package name to the corresponding file path
func ExtractPackageMap(top_level_path string, filename string) (map[string]string, error) {

	data, err := os.ReadFile(top_level_path + filename)
	if err != nil {
		return nil, err
	}

	var manifest Manifest_json
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, pkg := range manifest.Packages {
		result[pkg.Name] = top_level_path + pkg.Files.Objects.FileName
	}
	return result, nil
}

// extract characters and image assets into their respective list containers
func ExtractCharacterPackages(package_manifest map[string]string) (Asset_json, Character_json, error) {
	data, err := os.ReadFile(package_manifest["Character_Exports"])
	if err != nil {
		return Asset_json{}, Character_json{}, err
	}

	var raw struct {
		Objects []json.RawMessage `json:"Objects"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		fmt.Println("error parsing JSON:", err)
		return Asset_json{}, Character_json{}, err
	}

	var type_only struct {
		Type string `json:"Type"`
	}

	asset_packages := make(Asset_json, 0)
	character_packages := make(Character_json, 0)

	for _, raw_item := range raw.Objects {
		if err := json.Unmarshal(raw_item, &type_only); err != nil {
			fmt.Println("error parsing JSON:", err)
			return Asset_json{}, Character_json{}, err
		}

		switch type_only.Type {
		case "Entity":
			var temp character_object
			json.Unmarshal(raw_item, &temp)
			character_packages[temp.Properties.TechnicalName] = temp
		case "Asset":
			var temp asset_object
			json.Unmarshal(raw_item, &temp)
			asset_packages[temp.Properties.TechnicalName] = temp

		}
	}

	return asset_packages, character_packages, err

}

// map Object IDs to Object TechnicalNames
func (heirarchy Heirarchy_json) IdToTechnicalNameMap() map[string]string {
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
