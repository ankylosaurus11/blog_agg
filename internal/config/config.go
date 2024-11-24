package gatorconfig

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	fileContent, err := os.ReadFile("/home/anky/workspace/github.com/ankylosaurus11/blog_agg/gatorconfig.json")
	if err != nil {
		fmt.Println(err)
	}

	var gatorconfig Config

	err = json.Unmarshal(fileContent, &gatorconfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(gatorconfig)
	return gatorconfig
}
