package tools

import (
	"encoding/json"
	"io/ioutil"
	"log"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

type AllowFunction struct {
	Functions []string `json:"functions"`
}

var AllowImportWasi = "wasi_unstable"

func checkFunction(function string) bool {
	jsonFile, _ := ioutil.ReadFile("./env/functions.json")
	data := AllowFunction{}
	_ = json.Unmarshal([]byte(jsonFile), &data)
	for _, f := range data.Functions {
		if f == function {
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
