package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type manifest_Package struct {
	Name  string `json:"Name"`
	Files struct {
		Objects struct {
			FileName string `json:"FileName"`
		} `json:"Objects"`
		Texts struct {
			FileName string `json:"FileName"`
		} `json:"Texts"`
	} `json:"Files"`
}

type Manifest_Json struct {
	Packages []manifest_Package `json:"Packages"`
}

func (m *Manifest_Json) FromFile(file_path string) error {
	parent_dir := path.Dir(file_path)
	data, err := os.ReadFile(file_path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	for index := range m.Packages {
		m.Packages[index].Files.Objects.FileName = path.Join(parent_dir, m.Packages[index].Files.Objects.FileName)
		m.Packages[index].Files.Texts.FileName = path.Join(parent_dir, m.Packages[index].Files.Texts.FileName)
	}

	return nil
}

func (m *Manifest_Json) ObjectMap() map[string]string {
	output_map := make(map[string]string, 0)
	for _, pkg := range m.Packages {
		output_map[pkg.Name] = pkg.Files.Objects.FileName
	}
	return output_map
}

func (m *Manifest_Json) LocalizationMap() map[string]string {
	output_map := make(map[string]string, 0)
	for _, pkg := range m.Packages {
		output_map[pkg.Name] = pkg.Files.Texts.FileName
	}
	return output_map
}

type Heirarchy_json struct {
	Id            string            `json:"Id"`
	TechnicalName string            `json:"TechnicalName"`
	Type          string            `json:"Type"`
	Children      *[]Heirarchy_json `Json:"Children,omitempty"`
}

func (m *Heirarchy_json) FromFile(file_path string) error {
	data, err := os.ReadFile(file_path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	return nil
}

type Raw_Object_Json struct {
	Objects []json.RawMessage `json:"Objects"`
}

func (r *Raw_Object_Json) FromFile(file_path string) error {
	data, err := os.ReadFile(file_path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}

	return nil
}

type character_Object struct {
	Properties struct {
		TechnicalName string `json:"TechnicalName"`
		DisplayName   string `json:"DisplayName"`
		PreviewImage  struct {
			Asset string `json:"Asset"`
		} `json:"PreviewImage"`
	} `json:"Properties"`
}
type asset_Object struct {
	AssetRef   string `json:"AssetRef"`
	Properties struct {
		TechnicalName string `json:"TechnicalName"`
		Id            string `json:"Id"`
	} `json:"Properties"`
}

type Asset_Json = map[string]asset_Object
type Character_Json = map[string]character_Object

// extract characters and image assets into their respective list containers
func ExtractCharacterPackages(raw_object *Raw_Object_Json) (Asset_Json, Character_Json, error) {

	var type_only struct {
		Type string `json:"Type"`
	}

	asset_packages := make(Asset_Json, 0)
	character_packages := make(Character_Json, 0)

	for _, raw_item := range raw_object.Objects {
		if err := json.Unmarshal(raw_item, &type_only); err != nil {
			fmt.Println("error parsing JSON:", err)
			return Asset_Json{}, Character_Json{}, err
		}

		switch type_only.Type {
		case "Entity":
			var temp character_Object
			json.Unmarshal(raw_item, &temp)
			character_packages[temp.Properties.TechnicalName] = temp
		case "Asset":
			var temp asset_Object
			json.Unmarshal(raw_item, &temp)
			asset_packages[temp.Properties.TechnicalName] = temp
		}
	}

	return asset_packages, character_packages, nil

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
