package util

type T33File struct {
	Tag     string `json:"tag"`
	Port    string `json:"port"`
	Service struct {
		Path string `json:"path"`
		File string `json:"file"`
	} `json:"service"`
}
