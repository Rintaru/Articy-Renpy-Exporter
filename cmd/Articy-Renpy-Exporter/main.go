package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func main() {
	data, err := os.ReadFile("/mnt/c/GIT_REPOS/Visual_Novels/Practice_Export/articy_export/package_0100000000000110_objects.json")
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}
	var cfg map[string]interface{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		fmt.Println("error parsing JSON:", err)
		return
	}

	fmt.Println(reflect.TypeOf(cfg["Objects"]))
	usersSlice, _ := cfg["Objects"].([]interface{})
	fmt.Println(usersSlice[0])
}
