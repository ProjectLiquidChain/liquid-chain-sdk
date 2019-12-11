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
	File string
}

func (c *Compile) Clang(option string) (string, string) {
	names := strings.Split(c.File, "/")
	last := names[len(names)-1]
	file := strings.Split(last, ".")
	language := file[len(file)-1]
	var nameFile, tool string
	if language == "cpp" {
		tool = CLANGPLUS
		nameFile = last[:len(last)-4]
	} else if language == "c" {
		tool = CLANG
		nameFile = last[:len(last)-2]
	} else {
		log.Fatal("file not support compile")
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
	return wasmFile, nameFile
}
func (c *Compile) Rust(name string) {
	cmd := exec.Command("cargo", "build", "--manifest-path", c.File+"/Cargo.toml", TARGET, "--release")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatal(err)
	}
	// log.Println(string(out))
	cmd = exec.Command("mv", c.File+"/target/wasm32-wasi/release/"+name, c.File+"/")
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatal(err)
	}
	cmd = exec.Command("rm", "-rf", c.File+"/target", c.File+"/Cargo.lock", c.File+"/.gitignore")
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatal(err)
	}
	log.Println(string(out))
}
