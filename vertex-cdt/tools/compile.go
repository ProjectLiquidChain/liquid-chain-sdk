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
		tool = "/usr/local/opt/wasi-sdk/bin/clang++"
		nameFile = last[:len(last)-4]
	} else if c.Language == "c" {
		tool = "/usr/local/opt/wasi-sdk/bin/clang"
		nameFile = last[:len(last)-2]
	}
	wasmFile := nameFile + ".wasm"
	function := strings.Split(option, ",")
	var op = ""
	if option != "" {
		for _, cfun := range function {
			op += ",--export=" + cfun
		}
	}
	cmd := exec.Command(tool, c.File, "-o", wasmFile, "-nostartfiles", "--target=wasm32-wasi", "-Wl,--no-entry,--allow-undefined,--demangle", "-Wl"+op, "--sysroot=/usr/local/opt/wasi-sdk/share/wasi-sysroot")
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
