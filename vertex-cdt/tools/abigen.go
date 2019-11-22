package tools

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

var allowType = []string{"uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "address"}

// define type of ABI
type Parameter struct {
	IsArray bool   `json:"is_array"`
	Type    string `json:"type"`
}
type Function struct {
	Name       string      `json:"name"`
	Parameters []Parameter `json:"parameters"`
}
type EventParam struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type Event struct {
	Name       string       `json:"name"`
	Parameters []EventParam `json:"parameters"`
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

func ABIgen(file string, language string, option string) (string, []string) {
	names := strings.Split(file, "/")
	last := names[len(names)-1]
	var nameFile string
	if language == "c++" {
		nameFile = last[:len(last)-4]
	} else if language == "c" {
		nameFile = last[:len(last)-2]
	}
	jsonFile := nameFile + "-abi.json"
	cmd := exec.Command("c2ffi", "-o", jsonFile, file, "--sys-include", "/usr/local/opt/wasi-sdk/share/wasi-sysroot/include")
	out, err := cmd.CombinedOutput()
	// log.Println(string(out))
	if err != nil {
		log.Println(string(out))
		log.Fatalln(err)
	}
	exportFunction := strings.Split(option, ",")
	event_names := parse(jsonFile, exportFunction)
	return jsonFile, event_names
}
func checkAllowType(atype string) bool {
	for _, ctype := range allowType {
		if ctype == atype {
			return true
		}
	}
	return false
}
func checkAllowFunction(function string, allowFunction []string) bool {
	for _, fn := range allowFunction {
		if fn == function {
			return true
		}
	}
	return false
}

func parse(file string, exportFunction []string) []string {
	jsonFile, _ := ioutil.ReadFile(file)
	data := []CFunction{}
	_ = json.Unmarshal([]byte(jsonFile), &data)
	result := ABI{}
	result.Version = 1
	functions := []Function{}
	events := []Event{}
	event_names := []string{}
	for i := 0; i < len(data); i++ {
		if data[i].Tag != "function" {
			continue
		}
		if data[i].ReturnType.Tag == "Event" {
			event_names = append(event_names, data[i].Name)
			event := parseEvent(data[i].Name, data[i].Parameters, data[i].Location)
			events = append(events, event)
			continue
		}
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
	return event_names
}

// parse to vertex event
func parseEvent(name string, params []Cparam, location string) Event {
	event_params := []EventParam{}
	for j := 0; j < len(params); j++ {
		param := EventParam{params[j].Name, params[j].Type.Tag}
		if params[j].Type.Tag[1:] == "array" || params[j].Type.Tag[1:] == "pointer" {
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
		param := Parameter{false, params[j].Type.Tag}
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
func convertType(ctype string) string {
	switch ctype {
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
		return ctype
	}
}
