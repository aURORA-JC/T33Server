package util

import (
	"encoding/json"
	"io/ioutil"
)

func ParseT33(p string) T33File {
	var f T33File

	// return default config
	if p == "" {
		return T33File{
			Tag:  "T33Server Example service",
			Port: "3330",
			Service: struct {
				Path string `json:"path"`
				File string `json:"file"`
			}(struct {
				Path string
				File string
			}{Path: "/", File: "index.html"}),
		}
	}

	// read customer config
	file, err := ioutil.ReadFile(p)
	if err != nil {
		return T33File{Port: "-1"}
	}

	err = json.Unmarshal(file, &f)
	if err != nil {
		return T33File{Port: "-1"}
	}

	return f
}
