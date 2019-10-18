package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

var AllowFunctionEnv = []string{"vs_value_set", "vs_value_get", "vs_value_size_get"}
var AllowImportWasi = "wasi_unstable"
var app = cli.NewApp()

func info() {
	app.Name = "smart contract development CLI"
	app.Usage = "vertex-cdt [language option] compile [file]"
	app.Version = "0.0.1"
	app.Author = "vertex team"
}
func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "c++",
			Aliases: []string{"c++"},
			Usage:   "compile c++ language file",
			Action: func(c *cli.Context) {
				result := compileCplus(c.Args().First())
				if checkImportFunction(result) {
					fmt.Println("compile completed!")
				} else {
					deleteFile(result) // ? remove file .wasm
				}
			},
		},
		{
			Name:    "c",
			Aliases: []string{"c"},
			Usage:   "compile c language file",
			Action: func(c *cli.Context) {
				result := compileC(c.Args().First())
				if checkImportFunction(result) {
					fmt.Println("compile completed!")
				} else {
					deleteFile(result) // ? remove file .wasm
				}
			},
		},
		{
			Name:    "rust",
			Aliases: []string{"r"},
			Usage:   "compile rust language file",
			Action: func(c *cli.Context) {
				file := strings.Split(c.Args().First(), "/")
				if file[len(file)-1] == "" {
					file[len(file)-1] = file[len(file)-2]
				}
				compileRust(c.Args().First())
				if checkImportFunction(c.Args().First() + "/target/wasm32-wasi/debug/" + file[len(file)-1] + ".wasm") {
					fmt.Println("compile completed!")
				} else {
					deleteFolder(c.Args().First() + "/target") // ? remove file .wasm
				}
			},
		},
	}
}
func checkFunction(fun string) bool {
	for _, cfun := range AllowFunctionEnv {
		if cfun == fun {
			return true
		}
	}
	return false
}
func compileCplus(file string) string {
	name := strings.Split(file, "/")
	last := name[len(name)-1]
	wasmFile := last[:len(last)-4] + ".wasm"
	cmd := exec.Command("/opt/wasi-sdk/bin/clang++", file, "-o", wasmFile, "--target=wasm32-wasi", "-Wl,--no-entry,--export=main", "--sysroot=/opt/wasi-sdk/share/wasi-sysroot")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
	return wasmFile
}
func compileC(file string) string {
	name := strings.Split(file, "/")
	last := name[len(name)-1]
	wasmFile := last[:len(last)-2] + ".wasm"
	cmd := exec.Command("/opt/wasi-sdk/bin/clang", file, "-o", wasmFile, "--target=wasm32-wasi", "-Wl,--no-entry,--export=main", "--sysroot=/opt/wasi-sdk/share/wasi-sysroot")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
	return wasmFile
}
func compileRust(folder string) {
	cmd := exec.Command("cargo", "build", "--manifest-path", folder+"/Cargo.toml", "--target", "wasm32-wasi")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
}
func deleteFile(file string) {
	cmd := exec.Command("rm", file)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
}
func deleteFolder(folder string) {
	cmd := exec.Command("rm", "-rf", folder)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
}
func checkImportFunction(file string) bool {
	bytes, _ := wasm.ReadBytes(file)
	var check = true
	compiled, err := wasm.Compile(bytes)
	if err != nil {
		panic(err)
	}
	importFunction := compiled.Imports
	for _, fn := range importFunction {
		if fn.Namespace != AllowImportWasi {
			if fn.Namespace == "env" {
				if !checkFunction(fn.Name) {
					fmt.Println("error: function " + fn.Name + " not support!")
					check = false
				}
			}
		} else {
			fmt.Println("warning env: " + fn.Namespace + " ,function " + fn.Name + " not support!")
		}
	}
	fmt.Println("check done!")
	return check
}
func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
