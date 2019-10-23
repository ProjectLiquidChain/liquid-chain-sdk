package tools

import (
	"log"
	"os/exec"
	"strings"
)

type Compile struct {
	File     string
	Language string
}

func (c *Compile) Clang(option string) string {
	names := strings.Split(c.File, "/")
	last := names[len(names)-1]
	var nameFile, tool string
	if c.Language == "c++" {
		tool = "/opt/wasi-sdk/bin/clang++"
		nameFile = last[:len(last)-4]
	} else if c.Language == "c" {
		tool = "/opt/wasi-sdk/bin/clang"
		nameFile = last[:len(last)-2]
	}
	wasmFile := nameFile + ".wasm"
	cmd := exec.Command(tool, c.File, "-o", wasmFile, "--target=wasm32-wasi", "-Wl,--no-entry,--export=main", "--sysroot="+option)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatal(err)
	}
	return wasmFile
}
func (c *Compile) Rust() {
	cmd := exec.Command("cargo", "build", "--manifest-path", c.File+"/Cargo.toml", "--target", "wasm32-wasi")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))
}
