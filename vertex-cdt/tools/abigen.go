package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

type Parameter struct {
	IsArray bool   `json:"is_array"`
	Type    string `json:"type"`
}
type Function struct {
	Name       string      `json:"name"`
	Parameters []Parameter `json:"parameters"`
}
type ABI struct {
	Version   string     `json:"version"`
	Functions []Function `json:"functions"`
}
type Ctype struct {
	Tag  string `json:"tag"`
	Type Type   `json:"type"`
}
type Cparam struct {
	Tag  string `json:"tag"`
	Type Ctype  `json:"type"`
}
type CFunction struct {
	Name       string   `json:"name"`
	Parameters []Cparam `json:"parameters"`
}
type Type struct {
	Tag  string `json:"tag"`
	Type string `json:"type"`
}

func ABIgen(file string, language string) {
	names := strings.Split(file, "/")
	last := names[len(names)-1]
	var nameFile string
	if language == "c++" {
		nameFile = last[:len(last)-4]
	} else if language == "c" {
		nameFile = last[:len(last)-2]
	}
	jsonFile := nameFile + ".json"
	cmd := exec.Command("c2ffi", "-o", jsonFile, file)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	parse(jsonFile)
	fmt.Print(string(out))
}
func parse(file string) {
	jsonFile, _ := ioutil.ReadFile(file)
	data := []CFunction{}
	_ = json.Unmarshal([]byte(jsonFile), &data)
	result := ABI{}
	result.Version = "1.0"
	functions := []Function{}
	for i := 0; i < len(data); i++ {
		params := []Parameter{}
		fmt.Println(data[i])
		for j := 0; j < len(data[i].Parameters); j++ {
			param := Parameter{false, data[i].Parameters[j].Type.Tag[1:]}
			if data[i].Parameters[j].Type.Tag[1:] == "array" {
				param.IsArray = true
				param.Type = data[i].Parameters[j].Type.Type.Tag[1:]
			}
			params = append(params, param)
		}
		function := Function{data[i].Name, params}
		functions = append(functions, function)
	}
	result.Functions = functions
	fmt.Println(result)
	resultJson, _ := json.Marshal(result)
	err := ioutil.WriteFile(file, resultJson, 0644)
	fmt.Printf("%+v", err)
}
