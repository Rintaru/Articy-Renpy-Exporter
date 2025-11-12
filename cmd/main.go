package main

import (
	"fmt"

	"github.com/Rintaru/Articy-Renpy-Exporter/internal"
)

type Character struct {
	Name       string
	Image_path string
}

func main() {
	var manifest = internal.Manifest_json{}
	err := manifest.From_file("/mnt/c/GIT_REPOS/Visual_Novels/Practice_Export/Organized_Export/manifest.json")
	if err != nil {
		fmt.Println("error loading JSON:", err)
		return
	}

	var object_json = internal.Raw_Object_Json{}
	err = object_json.From_file(manifest.ObjectMap()["Character_Exports"])
	if err != nil {
		fmt.Println("error extracting characters JSON:", err)
		return
	}

	_, _, err = internal.ExtractCharacterPackages(&object_json)
	if err != nil {
		fmt.Println("error extracting characters JSON:", err)
		return
	}
}
