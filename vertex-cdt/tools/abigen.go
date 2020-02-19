package tools

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var allowType = []string{"uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64",
	"float32", "float64", "address"}

const VERSION = 1
const SYS_INCLUDE = "/usr/local/opt/wasi-sdk/share/wasi-sysroot/include"

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
type Event struct {
	Name       string      `json:"name"`
	Parameters []Parameter `json:"parameters"`
}
type ABI struct {
	Version   int        `json:"version"`
	Events    []Event    `json:"events"`
	Functions []Function `json:"functions"`
}

// type of c2ffi output
type Ctype struct {
	Tag  string `json:"tag"`
	Type Type   `json:"type"`
}
type Cparam struct {
	Tag  string `json:"tag"`
	Name string `json:"name"`
	Type Ctype  `json:"type"`
}
type CFunction struct {
	Name       string   `json:"name"`
	Parameters []Cparam `json:"parameters"`
	Location   string   `json:"location"`
	ReturnType Type     `json:"return-type"`
	Tag        string   `json:"tag"`
}
type Type struct {
	Tag  string `json:"tag"`
	Type string `json:"type"`
}

/*
	function ABIgen create json file
	params:
		- file: name of file c or c++
		- language: c or c++
		- option: export functions name
		- wasmfile: name of file .wasm
*/
func ABIgen(file string, nameFile string, option string, wasmfile string) (string, []string) {
	jsonFile := nameFile + "-abi.json"
	cmd := exec.Command("c2ffi", "-o", jsonFile, file, "--sys-include", SYS_INCLUDE)
	out, err := cmd.CombinedOutput()
	// log.Println(string(out))
	if err != nil {
		log.Println(string(out))
		log.Fatalln(err)
	}
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
	events := []Event{}
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
				event := parseRustEvent(funcDecl)
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
func parseRustEvent(declEvent string) Event {
	event_params := []Parameter{}
	name := strings.Split(declEvent, "(")
	event_name := strings.Split(name[0], "fn")
	params := strings.Split(name[1], ")")
	list_params := strings.Split(params[0], ",")
	for _, param := range list_params {
		rust_type := strings.Split(param, ":")
		var param_rust Parameter
		if strings.Contains(rust_type[1], "[") {
			param_rust = Parameter{true, token(rust_type[0]), "array"}
			log.Println("error: type array is not support in event parameter !")
		} else {
			param_rust = Parameter{false, token(rust_type[0]), convertRustType(token(rust_type[1]))}
		}
		event_params = append(event_params, param_rust)
	}
	return Event{token(event_name[1]), event_params}
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
	events := []Event{}
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
				event := parseEvent(data[i].Name, data[i].Parameters, data[i].Location)
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

// parse to vertex event
func parseEvent(name string, params []Cparam, location string) Event {
	event_params := []Parameter{}
	for j := 0; j < len(params); j++ {
		param := Parameter{false, params[j].Name, params[j].Type.Tag}
		if params[j].Type.Tag[1:] == "array" || params[j].Type.Tag[1:] == "pointer" {
			param.IsArray = true
			log.Println(location, "variable "+params[j].Name, "warning: type array "+param.Type+" not support in event!")
			param.Type = ""
		} else if string(params[j].Type.Tag[0]) == ":" {
			param.Type = convertType(params[j].Type.Tag[1:])
		} else {
			if params[j].Type.Tag == "address" {
				param.Type = "address"
			} else {
				param.Type = params[j].Type.Tag[:len(param.Type)-2]
			}
		}
		event_params = append(event_params, param)
	}
	return Event{name, event_params}
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
