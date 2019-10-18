package tools

import (
	"fmt"
	"os/exec"
	"strings"
)

type Compile struct {
	File     string
	Language string
}

func (c *Compile) Clang() string {
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
	cmd := exec.Command(tool, c.File, "-o", wasmFile, "--target=wasm32-wasi", "-Wl,--no-entry,--export=main", "--sysroot=/opt/wasi-sdk/share/wasi-sysroot")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
	return wasmFile
}
func (c *Compile) Rust() {
	cmd := exec.Command("cargo", "build", "--manifest-path", c.File+"/Cargo.toml", "--target", "wasm32-wasi")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
}
