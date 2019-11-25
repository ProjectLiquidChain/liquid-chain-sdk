package tools

import (
	"log"
	"os/exec"
	"strings"
)

const SYS_WASI = "--sysroot=/usr/local/opt/wasi-sdk/share/wasi-sysroot"
const ALLOW_UNDEFINED = "-Wl,--no-entry,--allow-undefined,--demangle"
const TARGET = "--target=wasm32-wasi"
const CLANG = "/usr/local/opt/wasi-sdk/bin/clang"
const CLANGPLUS = "/usr/local/opt/wasi-sdk/bin/clang++"

type Compile struct {
	File     string
	Language string
}

func (c *Compile) Clang(option string) string {
	names := strings.Split(c.File, "/")
	last := names[len(names)-1]
	var nameFile, tool string
	if c.Language == "c++" {
		tool = CLANGPLUS
		nameFile = last[:len(last)-4]
	} else if c.Language == "c" {
		tool = CLANG
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
	cmd := exec.Command(tool, c.File, "-o", wasmFile, "-O3", "-nostartfiles", TARGET, ALLOW_UNDEFINED, "-Wl"+op, SYS_WASI)
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
