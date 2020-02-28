package tools

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const VERSION = 1

// define type of ABI
type Parameter struct {
	IsArray bool   `json:"is_array"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}
type Function struct {
	Name       string      `json:"name"`
	Parameters []Parameter `json:"parameters"`
}
type ABI struct {
	Version   int        `json:"version"`
	Events    []Function `json:"events"`
	Functions []Function `json:"functions"`
}

/*
	function ABI create json file
	params:
		- file: name of file c or c++
		- language: c or c++
		- option: export functions name
		- wasmfile: name of file .wasm
*/
func ABIC(file string, nameFile string, option string, wasmfile string) (string, []string) {
	jsonFile := c2ffi(file, nameFile)
	exportFunction := strings.Split(option, ",")
	event_names := parse(jsonFile, exportFunction, wasmfile)
	return jsonFile, event_names
}
func ABIRust(file string, nameFile string, path string, wasmfile string) (string, []string) {
	rustfile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer rustfile.Close()
	result := ABI{}
	result.Version = VERSION
	functions := []Function{}
	events := []Function{}
	event_names := []string{}
	import_func := getImportFunction(wasmfile)
	export_func := getExportFunction(wasmfile)
	scanner := bufio.NewScanner(rustfile)
	var funcDecl string
	var block bool
	var block_comment bool
	for scanner.Scan() {
		code := scanner.Text()
		if strings.Contains(code, "//") {
			code = strings.Split(code, "//")[0]
		}
		if strings.Contains(code, "/*") {
			code = strings.Split(code, "/*")[0]
			block_comment = true
		}
		if strings.Contains(code, "*/") && block_comment {
			code = strings.Split(code, "*/")[1]
			block_comment = false
		}
		if strings.Contains(code, "fn") {
			funcDecl = code
			block = true
		}
		if block {
			funcDecl += code
			if strings.Contains(code, "{") {
				function := parseRustFunction(funcDecl)
				if checkAllowFunction(function.Name, export_func) {
					functions = append(functions, function)
				}
				funcDecl = ""
				block = false
			}
			if strings.Contains(code, ";") && strings.Contains(funcDecl, "Event") {
				event := parseRustFunction(funcDecl)
				if checkAllowFunction(event.Name, import_func) {
					event_names = append(event_names, event.Name)
					events = append(events, event)
				} else {
					log.Println("warning: Event " + event.Name + " is declared but not use!")
				}
				funcDecl = ""
				block = false
			}
		}
	}
	result.Functions = functions
	result.Events = events
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	jsonFile := path + nameFile + "-abi.json"
	resultJson, _ := json.Marshal(result)
	err_json := ioutil.WriteFile(jsonFile, resultJson, 0644)
	if err_json != nil {
		log.Println(err_json)
	}
	return jsonFile, event_names
}
func token(stringToken string) string {
	return strings.ReplaceAll(stringToken, " ", "")
}
func parseRustFunction(declFunction string) Function {
	function_params := []Parameter{}
	name := strings.Split(declFunction, "(")
	function_name := strings.Split(name[0], "fn")
	params := strings.Split(name[1], ")")
	list_params := strings.Split(params[0], ",")
	for _, param := range list_params {
		rust_type := strings.Split(param, ":")
		if len(rust_type) < 2 {
			continue
		}
		var param_rust Parameter
		if strings.Contains(rust_type[1], "[") {
			array := token(rust_type[1])
			array = strings.Replace(array, "&", "", -1)
			array = strings.Replace(array, "[", "", -1)
			array = strings.Replace(array, "]", "", -1)
			array_type := strings.Split(array, ";")
			param_rust = Parameter{true, token(rust_type[0]), convertRustType(array_type[0])}
		} else {
			param_rust = Parameter{false, token(rust_type[0]), convertRustType(token(rust_type[1]))}
		}
		function_params = append(function_params, param_rust)
	}
	return Function{token(function_name[1]), function_params}
}

/*
	function parse: parse from c2ffi functions to vertex abi functions and events
	params:
		- file : name of c2ffi json file
		- exportFunction: list export function
		- wasmfile: name of wasm file
*/
func parse(file string, exportFunction []string, wasmfile string) []string {
	jsonFile, _ := ioutil.ReadFile(file)
	decls := []CFunction{}
	_ = json.Unmarshal([]byte(jsonFile), &decls)
	result := ABI{}
	result.Version = VERSION
	functions := []Function{}
	events := []Function{}
	eventNames := []string{}
	functionName := []string{}
	import_func := getImportFunction(wasmfile)
	for _, decl := range decls {
		if decl.Tag != "function" {
			continue
		}
		if decl.ReturnType.Tag == Event {
			if checkAllowFunction(decl.Name, import_func) {
				eventNames = append(eventNames, decl.Name)
				event := parseFunction(decl.Name, decl.Parameters, decl.Location)
				events = append(events, event)
			} else {
				log.Println("warning: "+decl.Location, "Event "+decl.Name+" is declared but not use!")
			}
			continue
		}
		functionName = append(functionName, decl.Name)
		if !checkAllowFunction(decl.Name, exportFunction) {
			continue
		}

		function := parseFunction(decl.Name, decl.Parameters, decl.Location)
		functions = append(functions, function)
	}
	result.Functions = functions
	result.Events = events
	resultJson, _ := json.Marshal(result)
	err := ioutil.WriteFile(file, resultJson, 0644)
	if err != nil {
		log.Println(err)
	}
	// warning function not found
	for _, fn := range exportFunction {
		if !checkAllowFunction(fn, functionName) {
			log.Println("warning: ", "export function "+fn+" not found!")
		}
	}
	return eventNames
}

// parse to vertex function
func parseFunction(name string, cparams []Cparam, location string) Function {
	params := []Parameter{}
	for _, cparam := range cparams {
		param := Parameter{false, cparam.Name, cparam.Type.Tag}
		if cparam.Type.Tag[1:] == Array || cparam.Type.Tag[1:] == Pointer {
			param.IsArray = true
			param.Type = cparam.Type.Type.Tag
			if string(cparam.Type.Type.Tag[0]) == ":" {
				param.Type = convertType(param.Type[1:])
			} else {
				param.Type = cparam.Type.Type.Tag[:len(param.Type)-2]
			}
		} else if string(cparam.Type.Tag[0]) == ":" {
			param.IsArray = false
			param.Type = convertType(cparam.Type.Tag[1:])
		} else {
			param.IsArray = false
			if cparam.Type.Tag == Address {
				param.Type = Address
			} else if cparam.Type.Tag == PlArray {
				param.Type = PlArray
			} else {
				param.Type = cparam.Type.Tag[:len(param.Type)-2]
			}
		}
		if !validateType(param.Type) {
			log.Println(location, "variable "+cparam.Name, "warning: type "+param.Type+" not support!")
		}
		params = append(params, param)
	}
	return Function{name, params}
}
