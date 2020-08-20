package tools

import (
	"log"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

var AllowFunctionEnv = []string{"chain_print_bytes", "chain_event_emit", "chain_get_caller", "chain_get_creator",
	"chain_invoke", "chain_get_owner", "chain_method_bind", "chain_arg_size_get", "chain_arg_size_set", "get_mean",
	"sum_of_squares", "sqroot", "get_average", "address_xor", "chain_block_height", "chain_block_time", "chain_storage_set",
	"chain_storage_size_get", "chain_storage_get", "chain_args_write", "chain_args_hash", "chain_ed25519_verify", "chain_get_contract_address"}

var AllowImportWasi = "wasi_unstable"

func checkAllowFunction(function string, allowFunction []string) bool {
	for _, fn := range allowFunction {
		if fn == function {
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
			if fn.Namespace == "env" && !checkAllowFunction(fn.Name, AllowFunctionEnv) && !checkEvent(fn.Name, event_names) {
				log.Println("warning: function " + fn.Name + " not support!")
				check = true
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

func getExportFunction(file string) []string {
	list := []string{}

	bytes, _ := wasm.ReadBytes(file)
	compiled, err := wasm.Compile(bytes)
	if err != nil {
		log.Fatalln(err)
	}
	exportFunction := compiled.Exports
	for _, fn := range exportFunction {
		list = append(list, fn.Name)
	}
	return list
}
