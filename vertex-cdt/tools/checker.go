package tools

import (
	"log"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

var AllowFunctionEnv = []string{"chain_storage_size_get", "chain_storage_get", "chain_storage_set",
	"chain_print_bytes", "chain_event_emit", "chain_get_caller", "chain_get_creator", "chain_invoke", "chain_get_owner"}

var AllowImportWasi = "wasi_unstable"

func checkFunction(fun string) bool {
	for _, cfun := range AllowFunctionEnv {
		if cfun == fun {
			return true
		}
	}
	return false
}

func checkEvent(event string, events []string) bool {
	for _, e := range events {
		if e == event {
			return true
		}
	}
	return false
}

func CheckImportFunction(file string, event_names []string) bool {
	bytes, _ := wasm.ReadBytes(file)
	var check = true
	compiled, err := wasm.Compile(bytes)
	if err != nil {
		log.Fatalln(err)
	}
	importFunction := compiled.Imports
	for _, fn := range importFunction {
		if fn.Namespace != AllowImportWasi {
			if fn.Namespace == "env" && !checkFunction(fn.Name) && !checkEvent(fn.Name, event_names) {
				log.Println("error: function " + fn.Name + " not support!")
				check = false
			}
		} else {
			log.Println("warning env: " + fn.Namespace + " ,function " + fn.Name + " not support!")
		}
	}
	log.Println("check done!")
	return check
}

func getImportFunction(file string) []string {
	list := []string{}

	bytes, _ := wasm.ReadBytes(file)
	compiled, err := wasm.Compile(bytes)
	if err != nil {
		log.Fatalln(err)
	}
	importFunction := compiled.Imports
	for _, fn := range importFunction {
		if fn.Namespace == "env" {
			list = append(list, fn.Name)
		}
	}
	return list
}
