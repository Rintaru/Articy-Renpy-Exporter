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

	_, _, err := object_json.ExtractCharacterPackages()
	if err != nil {
		fmt.Println("error extracting characters JSON:", err)
		return
	}

	_, err := LocalizationJsonFromFile(manifest.LocalizationMap()["Character_Exports"])
	if err != nil {
		fmt.Println("error extracting characters JSON:", err)
		return
	}

}
