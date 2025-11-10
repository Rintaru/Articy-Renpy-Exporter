package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func get_packages(manifest_file_path string) map[string]string {
	data, err := os.ReadFile(manifest_file_path)
	if err != nil {
		fmt.Println("error reading file:", err)
	}

	var top_level_dictionary map[string]any
	err = json.Unmarshal(data, &top_level_dictionary)
	if err != nil {
		fmt.Println("error parsing JSON:", err)
	}

	for _, package_dict := range top_level_dictionary["Packages"].([]any) {
		package_dict[]
	}

}
func main() {
	data, err := os.ReadFile("/mnt/c/GIT_REPOS/Visual_Novels/Practice_Export/Organized_Export/package_0100000000000D66_objects.json")
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}
	var parsed_data map[string]any
	if err := json.Unmarshal(data, &parsed_data); err != nil {
		fmt.Println("error parsing JSON:", err)
		return
	}

	for _, item := range parsed_data["Objects"].([]any) {
		parsed_item := item.(map[string]any)

		fmt.Println(parsed_item["Type"].(string))

	}
}
