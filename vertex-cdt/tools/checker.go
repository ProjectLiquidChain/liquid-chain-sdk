package tools

import (
	"fmt"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

var AllowFunctionEnv = []string{"vs_value_set", "vs_value_get", "vs_value_size_get"}
var AllowImportWasi = "wasi_unstable"

func checkFunction(fun string) bool {
	for _, cfun := range AllowFunctionEnv {
		if cfun == fun {
			return true
		}
	}
	return false
}

func CheckImportFunction(file string) bool {
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
