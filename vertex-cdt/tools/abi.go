package tools

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var allowType = []string{"uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64",
	"float32", "float64", "address", "plarray"}

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
	function checkAllowType : check vertex VM type.
	params:
		- atype: c or c++ type
*/
func checkAllowType(atype string) bool {
	for _, ctype := range allowType {
		if ctype == atype {
			return true
		}
	}
	return false
}

/*
	function checkAllowFunction: check the export functions
	params:
		- function
		- allowFunction: list export functions or events
*/
func checkAllowFunction(function string, allowFunction []string) bool {
	for _, fn := range allowFunction {
		if fn == function {
			return true
		}
	}
	return false
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
	data := []CFunction{}
	_ = json.Unmarshal([]byte(jsonFile), &data)
	result := ABI{}
	result.Version = VERSION
	functions := []Function{}
	events := []Function{}
	event_names := []string{}
	function_name := []string{}
	import_func := getImportFunction(wasmfile)
	for i := 0; i < len(data); i++ {
		if data[i].Tag != "function" {
			continue
		}
		if data[i].ReturnType.Tag == "Event" {
			if checkAllowFunction(data[i].Name, import_func) {
				event_names = append(event_names, data[i].Name)
				event := parseFunction(data[i].Name, data[i].Parameters, data[i].Location)
				events = append(events, event)
			} else {
				log.Println("warning: "+data[i].Location, "Event "+data[i].Name+" is declared but not use!")
			}
			continue
		}
		function_name = append(function_name, data[i].Name)
		if !checkAllowFunction(data[i].Name, exportFunction) {
			continue
		}

		function := parseFunction(data[i].Name, data[i].Parameters, data[i].Location)
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
		if !checkAllowFunction(fn, function_name) {
			log.Println("warning: ", "export function "+fn+" not found!")
		}
	}
	return event_names
}

// parse to vertex function
func parseFunction(name string, params []Cparam, location string) Function {
	function_params := []Parameter{}
	for j := 0; j < len(params); j++ {
		param := Parameter{false, params[j].Name, params[j].Type.Tag}
		if params[j].Type.Tag[1:] == "array" || params[j].Type.Tag[1:] == "pointer" {
			param.IsArray = true
			param.Type = params[j].Type.Type.Tag
			if string(params[j].Type.Type.Tag[0]) == ":" {
				param.Type = convertType(param.Type[1:])
			} else {
				param.Type = params[j].Type.Type.Tag[:len(param.Type)-2]
			}
		} else if string(params[j].Type.Tag[0]) == ":" {
			param.IsArray = false
			param.Type = convertType(params[j].Type.Tag[1:])
		} else {
			param.IsArray = false
			if params[j].Type.Tag == "address" {
				param.Type = "address"
			} else if params[j].Type.Tag == "plarray" {
				param.Type = "plarray"
			} else {
				param.Type = params[j].Type.Tag[:len(param.Type)-2]
			}
		}
		if !checkAllowType(param.Type) {
			log.Println(location, "variable "+params[j].Name, "warning: type "+param.Type+" not support!")
		}
		function_params = append(function_params, param)
	}
	return Function{name, function_params}
}

// convert from c,c++ type to assembly vertex type
func convertType(Type string) string {
	switch Type {
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "signed-char":
		return "int8"
	case "char":
		return "int8"
	case "unsigned-char":
		return "uint8"
	case "short":
		return "int16"
	case "unsigned-short":
		return "uint16"
	case "int":
		return "int32"
	case "unsigned-int":
		return "uint32"
	case "unsigned-long":
		return "uint32"
	case "long-long":
		return "int64"
	case "unsigned-long-long":
		return "uint64"
	default:
		return Type
	}
}
func convertRustType(Type string) string {
	switch Type {
	case "f32":
		return "float32"
	case "f64":
		return "float64"
	case "i8":
		return "int8"
	case "u8":
		return "uint8"
	case "i16":
		return "int16"
	case "u16":
		return "uint16"
	case "i32":
		return "int32"
	case "u32":
		return "uint32"
	case "i64":
		return "int64"
	case "u64":
		return "uint64"
	default:
		return Type
	}
}
