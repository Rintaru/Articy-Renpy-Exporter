package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Rintaru/Articy-Renpy-Exporter/internal"
)

type Character struct {
	Name       string
	Image_path string
}

func main() {
	top_level_path := "/mnt/c/GIT_REPOS/Visual_Novels/Practice_Export/Organized_Export/"
	package_map, err := internal.ExtractPackageMap(top_level_path, "manifest.json")
	if err != nil {
		return
	}

	data, err := os.ReadFile(top_level_path + "hierarchy.json")
	if err != nil {
		fmt.Println("error loading JSON:", err)
		return
	}

	var heirarchy_data internal.Heirarchy_json
	if err := json.Unmarshal(data, &heirarchy_data); err != nil {
		fmt.Println("error parsing JSON:", err)
		return
	}
	_, _, err = internal.ExtractCharacterPackages(package_map)
	if err != nil {
		fmt.Println("error extracting characters JSON:", err)
		return
	}

}
