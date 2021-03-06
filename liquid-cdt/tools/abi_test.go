package tools

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"testing"
)

func TestParseFunction(t *testing.T) {
	params := []Cparam{
		Cparam{
			Tag:  "parameter",
			Name: "x",
			Type: Ctype{
				Tag: ":int",
				Type: Type{
					Tag:  "",
					Type: "",
				},
			},
		},
		Cparam{
			Tag:  "parameter",
			Name: "y",
			Type: Ctype{
				Tag: ":array",
				Type: Type{
					Tag:  ":float",
					Type: "",
				},
			},
		},
		Cparam{
			Tag:  "parameter",
			Name: "z",
			Type: Ctype{
				Tag: "address",
				Type: Type{
					Tag:  "",
					Type: "",
				},
			},
		},
		Cparam{
			Tag:  "parameter",
			Name: "t",
			Type: Ctype{
				Tag: "uint32_t",
				Type: Type{
					Tag:  "",
					Type: "",
				},
			},
		},
	}
	function := parseFunction("add", params, "line 19")
	liquid_function := Function{
		Name: "add",
		Parameters: []Parameter{
			Parameter{
				Name: "x",
				Type: "int32",
			},
			Parameter{
				Name: "y",
				Type: "float32[]",
			},
			Parameter{
				Name: "z",
				Type: "address",
			},
			Parameter{
				Name: "t",
				Type: "uint32",
			},
		},
	}
	if function.Name != liquid_function.Name {
		t.Errorf("function was incorrect name, got: %s, want: %s.", function.Name, liquid_function.Name)
	}
	if function.Parameters[0].Type != liquid_function.Parameters[0].Type {
		t.Errorf("function was incorrect parameter type index 0 , got: %s, want: %s.", function.Parameters[0].Type, liquid_function.Parameters[0].Type)
	}

	if function.Parameters[1].Type != liquid_function.Parameters[1].Type {
		t.Errorf("function was incorrect parameter type index 1 , got: %s, want: %s.", function.Parameters[1].Type, liquid_function.Parameters[1].Type)
	}

	if function.Parameters[2].Type != liquid_function.Parameters[2].Type {
		t.Errorf("function was incorrect parameter type index 2 , got: %s, want: %s.", function.Parameters[2].Type, liquid_function.Parameters[2].Type)
	}

	if function.Parameters[3].Type != liquid_function.Parameters[3].Type {
		t.Errorf("function was incorrect parameter type index 3 , got: %s, want: %s.", function.Parameters[3].Type, liquid_function.Parameters[3].Type)
	}

}
func TestParseEvent(t *testing.T) {
	params := []Cparam{
		Cparam{
			Tag:  "parameter",
			Name: "from",
			Type: Ctype{
				Tag: "address",
				Type: Type{
					Tag:  "",
					Type: "",
				},
			},
		},
		Cparam{
			Tag:  "parameter",
			Name: "to",
			Type: Ctype{
				Tag: "address",
				Type: Type{
					Tag:  "",
					Type: "",
				},
			},
		},
		Cparam{
			Tag:  "parameter",
			Name: "amount",
			Type: Ctype{
				Tag: "uint64_t",
				Type: Type{
					Tag:  "",
					Type: "",
				},
			},
		},
	}
	event := parseFunction("Transfer", params, "line 9")
	liquid_event := Function{
		Name: "Transfer",
		Parameters: []Parameter{
			Parameter{
				Name: "from",
				Type: "address",
			},
			Parameter{
				Name: "to",
				Type: "address",
			},
			Parameter{
				Name: "amount",
				Type: "uint64",
			},
		},
	}
	if event.Name != liquid_event.Name {
		t.Errorf("event was incorrect name, got: %s, want: %s.", event.Name, liquid_event.Name)
	}
	if event.Parameters[0].Name != liquid_event.Parameters[0].Name {
		t.Errorf("event was incorrect parameter type name, got: %s, want: %s.", event.Parameters[0].Name, liquid_event.Parameters[0].Name)
	}
	if event.Parameters[0].Type != liquid_event.Parameters[0].Type {
		t.Errorf("event was incorrect parameter type, got: %s, want: %s.", event.Parameters[0].Type, liquid_event.Parameters[0].Type)
	}
	if event.Parameters[1].Name != liquid_event.Parameters[1].Name {
		t.Errorf("event was incorrect parameter type name, got: %s, want: %s.", event.Parameters[1].Name, liquid_event.Parameters[1].Name)
	}
	if event.Parameters[1].Type != liquid_event.Parameters[1].Type {
		t.Errorf("event was incorrect parameter type, got: %s, want: %s.", event.Parameters[1].Type, liquid_event.Parameters[1].Type)
	}
	if event.Parameters[2].Name != liquid_event.Parameters[2].Name {
		t.Errorf("event was incorrect parameter type name, got: %s, want: %s.", event.Parameters[2].Name, liquid_event.Parameters[2].Name)
	}
	if event.Parameters[2].Type != liquid_event.Parameters[2].Type {
		t.Errorf("event was incorrect parameter type, got: %s, want: %s.", event.Parameters[2].Type, liquid_event.Parameters[2].Type)
	}
}

func TestToken(t *testing.T) {
	token_t := token("  add")
	if token_t != "add" {
		t.Errorf("token was incorrect, got: %s, want: %s.", token_t, "add")
	}
}

func TestParseRustEvent(t *testing.T) {
	event := parseRustFunction("fn Transfer(from: address, to: address, amount: u64) -> Event;")
	liquid_event := Function{
		Name: "Transfer",
		Parameters: []Parameter{
			Parameter{
				Name: "from",
				Type: "address",
			},
			Parameter{
				Name: "to",
				Type: "address",
			},
			Parameter{
				Name: "amount",
				Type: "uint64",
			},
		},
	}
	if event.Name != liquid_event.Name {
		t.Errorf("event was incorrect name, got: %s, want: %s.", event.Name, liquid_event.Name)
	}
	if event.Parameters[0].Name != liquid_event.Parameters[0].Name {
		t.Errorf("event was incorrect parameter type name, got: %s, want: %s.", event.Parameters[0].Name, liquid_event.Parameters[0].Name)
	}
	if event.Parameters[0].Type != liquid_event.Parameters[0].Type {
		t.Errorf("event was incorrect parameter type, got: %s, want: %s.", event.Parameters[0].Type, liquid_event.Parameters[0].Type)
	}
	if event.Parameters[1].Name != liquid_event.Parameters[1].Name {
		t.Errorf("event was incorrect parameter type name, got: %s, want: %s.", event.Parameters[1].Name, liquid_event.Parameters[1].Name)
	}
	if event.Parameters[1].Type != liquid_event.Parameters[1].Type {
		t.Errorf("event was incorrect parameter type, got: %s, want: %s.", event.Parameters[1].Type, liquid_event.Parameters[1].Type)
	}
	if event.Parameters[2].Name != liquid_event.Parameters[2].Name {
		t.Errorf("event was incorrect parameter type name, got: %s, want: %s.", event.Parameters[2].Name, liquid_event.Parameters[2].Name)
	}
	if event.Parameters[2].Type != liquid_event.Parameters[2].Type {
		t.Errorf("event was incorrect parameter type, got: %s, want: %s.", event.Parameters[2].Type, liquid_event.Parameters[2].Type)
	}
	event = parseRustFunction("fn ArrayTest(from: &[u8]) -> Event;")
	if event.Name != "ArrayTest" {
		t.Errorf("event was incorrect name, got: %s, want: %s.", event.Name, "ArrayTest")
	}
	if event.Parameters[0].Name != "from" {
		t.Errorf("event was incorrect parameter type name, got: %s, want: %s.", event.Parameters[0].Name, "from")
	}
	if event.Parameters[0].Type != "uint8[]" {
		t.Errorf("event was incorrect parameter type, got: %s, want: %s.", event.Parameters[0].Type, "uint8[]")
	}

}

func TestParseRustFunction(t *testing.T) {
	function := parseRustFunction("fn add(x: i32, y: &[f32], z: address, t:u32) { ")
	liquid_function := Function{
		Name: "add",
		Parameters: []Parameter{
			Parameter{
				Type: "int32",
			},
			Parameter{
				Type: "float32[]",
			},
			Parameter{
				Type: "address",
			},
			Parameter{
				Type: "uint32",
			},
		},
	}
	if function.Name != liquid_function.Name {
		t.Errorf("function was incorrect name, got: %s, want: %s.", function.Name, liquid_function.Name)
	}
	if function.Parameters[0].Type != liquid_function.Parameters[0].Type {
		t.Errorf("function was incorrect parameter type index 0 , got: %s, want: %s.", function.Parameters[0].Type, liquid_function.Parameters[0].Type)
	}

	if function.Parameters[1].Type != liquid_function.Parameters[1].Type {
		t.Errorf("function was incorrect parameter type index 1 , got: %s, want: %s.", function.Parameters[1].Type, liquid_function.Parameters[1].Type)
	}

	if function.Parameters[2].Type != liquid_function.Parameters[2].Type {
		t.Errorf("function was incorrect parameter type index 2 , got: %s, want: %s.", function.Parameters[2].Type, liquid_function.Parameters[2].Type)
	}

	if function.Parameters[3].Type != liquid_function.Parameters[3].Type {
		t.Errorf("function was incorrect parameter type index 3 , got: %s, want: %s.", function.Parameters[3].Type, liquid_function.Parameters[3].Type)
	}

}

func TestParse(t *testing.T) {
	jsonFile, _ := ioutil.ReadFile("./tests/contract-abi.json")
	data := []CFunction{}
	_ = json.Unmarshal([]byte(jsonFile), &data)
	abi := parse("./tests/contract-abi.json", []string{"add"}, "./tests/contract.wasm")
	if len(abi) > 0 {
		t.Errorf("parse fail")
	}
	resultJson, _ := json.Marshal(data)
	err := ioutil.WriteFile("./tests/contract-abi.json", resultJson, 0644)
	if err != nil {
		log.Println(err)
	}
}

func TestABIRust(t *testing.T) {
	jsonFile, events := ABIRust("./tests/add/src/lib.rs", "add", "./tests/add/", "./tests/add/add.wasm")
	if jsonFile != "./tests/add/add-abi.json" {
		t.Errorf("parse fail")
	}
	if events[0] != "Add" {
		t.Errorf("parse event fail")
	}
}
func TestABIC(t *testing.T) {
	jsonFile, events := ABIC("./tests/contract.c", "contract", "add", "./tests/add/add.wasm")
	if jsonFile != "contract-abi.json" {
		t.Errorf("parse fail")
	}
	if events[0] != "Add" {
		t.Errorf("parse event fail")
	}
	cmd := exec.Command("rm", "-rf", jsonFile)
	_, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("file not found")
	}
}
