package main

import (
	"fmt"
)

type Character struct {
	Name       string
	Image_path string
}

func main() {
	var manifest = Manifest_Json{}
	err := manifest.FromFile("/mnt/c/GIT_REPOS/Visual_Novels/Practice_Export/Organized_Export/manifest.json")
	if err != nil {
		fmt.Println("error loading JSON:", err)
		return
	}

	var object_json = Raw_Object_Json{}
	err = object_json.FromFile(manifest.ObjectMap()["Character_Exports"])
	if err != nil {
		fmt.Println("error extracting characters JSON:", err)
		return
	}

	asset_objects, character_objects, err := object_json.ExtractCharacterPackage()
	if err != nil {
		fmt.Println("error extracting characters JSON:", err)
		return
	}

	localization_json, err := LocalizationJsonFromFile(manifest.LocalizationMap()["Character_Exports"])
	if err != nil {
		fmt.Println("error extracting characters JSON:", err)
		return
	}

	character_list := make([]Character, 0)

	heirarchy_json := Heirarchy_json{}
	err = heirarchy_json.FromFile("/mnt/c/GIT_REPOS/Visual_Novels/Practice_Export/Organized_Export/hierarchy.json")
	if err != nil {
		fmt.Println("error extracting characters JSON:", err)
		return
	}

	for _, char := range character_objects {
		asset_technical_name := heirarchy_json.IdToTechnicalNameMap()[char.Properties.PreviewImage.Asset]
		character_list = append(character_list, Character{
			Name:       localization_json[char.Properties.DisplayName].Text,
			Image_path: asset_objects[asset_technical_name].AssetRef,
		})
	}

	for _, item := range character_list {
		fmt.Println(item)
	}

}
