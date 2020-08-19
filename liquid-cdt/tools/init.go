package tools

import (
	"io/ioutil"
	"log"
	"os/exec"
)

func Create(name string) string {
	cmd := exec.Command("cargo", "new", "--lib", name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	log.Print(string(out))
	config := "./" + name + "/Cargo.toml"
	contents, _ := ioutil.ReadFile(config)
	lib := []byte("[lib] \n crate-type = [\"cdylib\"] \n")
	op := []byte("[profile.release] \n lto = true \n")
	result := append(contents, lib...)
	result = append(result, op...)
	ioutil.WriteFile(config, result, 0644)
	return name
}
