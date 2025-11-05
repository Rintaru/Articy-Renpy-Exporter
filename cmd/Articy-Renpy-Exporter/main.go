package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("/mnt/c/GIT_REPOS/Visual_Novels/Practice_Export/articy_export/package_0100000000000110_objects.json")
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
