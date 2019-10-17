package main

import (
	"fmt"
	"os"
	"os/exec"

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
				compile(c.Args().First())
				if checkImportFunction("contract.wasm") {
					fmt.Println("compile completed!")
				} else {
					deleteFile("contract.wasm") // ? remove file .wasm
				}
			},
		},
		{
			Name:    "c",
			Aliases: []string{"c"},
			Usage:   "compile c language file",
			Action: func(c *cli.Context) {
				compile(c.Args().First())
				if checkImportFunction("contract.wasm") {
					fmt.Println("compile completed!")
				} else {
					deleteFile("contract.wasm") // ? remove file .wasm
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
func compile(file string) {
	cmd := exec.Command("/opt/wasi-sdk/bin/clang++", file, "-o", "contract.wasm", "--target=wasm32-wasi", "-Wl,--no-entry,--export=main", "--sysroot=/opt/wasi-sdk/share/wasi-sysroot")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
}
func deleteFile(file string) {
	fmt.Println(file)
	cmd := exec.Command("rm", file)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
}
func checkImportFunction(file string) bool {
	bytes, _ := wasm.ReadBytes(file)

	compiled, err := wasm.Compile(bytes)
	if err != nil {
		panic(err)
	}
	importFunction := compiled.Imports
	for _, fn := range importFunction {
		if fn.Namespace != AllowImportWasi {
			if fn.Namespace == "env" {
				if !checkFunction(fn.Name) {
					return false
				}
			}
		} else {
			fmt.Println("warning env: " + fn.Namespace + " ,function " + fn.Name + " not support!")
		}
	}
	fmt.Println("check done!")
	return true
}
func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
