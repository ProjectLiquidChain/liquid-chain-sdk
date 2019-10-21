package tools

import (
	"fmt"
	"os/exec"
	"strings"
)

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
	fmt.Print(string(out))
}
