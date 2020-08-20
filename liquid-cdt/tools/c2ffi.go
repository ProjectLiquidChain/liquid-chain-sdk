package tools

import (
	"log"
	"os/exec"
)

const SYS_INCLUDE = "/usr/local/opt/wasi-sdk/share/wasi-sysroot/include"

// type of c2ffi output
type Ctype struct {
	Tag  string `json:"tag"`
	Type Type   `json:"type"`
}
type Cparam struct {
	Tag  string `json:"tag"`
	Name string `json:"name"`
	Type Ctype  `json:"type"`
}
type CFunction struct {
	Name       string   `json:"name"`
	Parameters []Cparam `json:"parameters"`
	Location   string   `json:"location"`
	ReturnType Type     `json:"return-type"`
	Tag        string   `json:"tag"`
}
type Type struct {
	Tag  string `json:"tag"`
	Type string `json:"type"`
}

func c2ffi(file string, nameFile string) string {
	jsonFile := nameFile + "-abi.json"
	cmd := exec.Command("c2ffi", "-o", jsonFile, file, "--sys-include", SYS_INCLUDE)
	out, err := cmd.CombinedOutput()
	// log.Println(string(out))
	if err != nil {
		log.Println(string(out))
		log.Fatalln(err)
		return ""
	}
	return jsonFile
}
