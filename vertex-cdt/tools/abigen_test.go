package tools

import "testing"

func TestCovertRustType(t *testing.T) {
	f32 := convertRustType("f32")
	if f32 != "float32" {
		t.Errorf("type was incorrect, got: %s, want: %s.", f32, "float32")
	}
	f64 := convertRustType("f64")
	if f64 != "float64" {
		t.Errorf("type was incorrect, got: %s, want: %s.", f64, "float64")
	}
	i8 := convertRustType("i8")
	if i8 != "int8" {
		t.Errorf("type was incorrect, got: %s, want: %s.", i8, "int8")
	}
	u8 := convertRustType("u8")
	if u8 != "uint8" {
		t.Errorf("type was incorrect, got: %s, want: %s.", u8, "uint8")
	}
	i16 := convertRustType("i16")
	if i16 != "int16" {
		t.Errorf("type was incorrect, got: %s, want: %s.", i16, "int16")
	}
	u16 := convertRustType("u16")
	if u16 != "uint16" {
		t.Errorf("type was incorrect, got: %s, want: %s.", u16, "uint16")
	}
	i32 := convertRustType("i32")
	if i32 != "int32" {
		t.Errorf("type was incorrect, got: %s, want: %s.", i32, "int32")
	}
	u32 := convertRustType("u32")
	if u32 != "uint32" {
		t.Errorf("type was incorrect, got: %s, want: %s.", u32, "uint32")
	}
	i64 := convertRustType("i64")
	if i64 != "int64" {
		t.Errorf("type was incorrect, got: %s, want: %s.", i64, "int64")
	}
	u64 := convertRustType("u64")
	if u64 != "uint64" {
		t.Errorf("type was incorrect, got: %s, want: %s.", u64, "uint64")
	}
	address := convertRustType("address")
	if address != "address" {
		t.Errorf("type was incorrect, got: %s, want: %s.", address, "address")
	}
}

func TestCovertType(t *testing.T) {
	f32 := convertType("float")
	if f32 != "float32" {
		t.Errorf("type was incorrect, got: %s, want: %s.", f32, "float32")
	}
	f64 := convertType("double")
	if f64 != "float64" {
		t.Errorf("type was incorrect, got: %s, want: %s.", f64, "float64")
	}
	i8 := convertType("signed-char")
	if i8 != "int8" {
		t.Errorf("type was incorrect, got: %s, want: %s.", i8, "int8")
	}
	u8 := convertType("unsigned-char")
	if u8 != "uint8" {
		t.Errorf("type was incorrect, got: %s, want: %s.", u8, "uint8")
	}
	i16 := convertType("short")
	if i16 != "int16" {
		t.Errorf("type was incorrect, got: %s, want: %s.", i16, "int16")
	}
	u16 := convertType("unsigned-short")
	if u16 != "uint16" {
		t.Errorf("type was incorrect, got: %s, want: %s.", u16, "uint16")
	}
	i32 := convertType("int")
	if i32 != "int32" {
		t.Errorf("type was incorrect, got: %s, want: %s.", i32, "int32")
	}
	u32 := convertType("unsigned-int")
	if u32 != "uint32" {
		t.Errorf("type was incorrect, got: %s, want: %s.", u32, "uint32")
	}
	i64 := convertType("long-long")
	if i64 != "int64" {
		t.Errorf("type was incorrect, got: %s, want: %s.", i64, "int64")
	}
	u64 := convertType("unsigned-long-long")
	if u64 != "uint64" {
		t.Errorf("type was incorrect, got: %s, want: %s.", u64, "uint64")
	}
	address := convertType("address")
	if address != "address" {
		t.Errorf("type was incorrect, got: %s, want: %s.", address, "address")
	}
}

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
	vertex_function := Function{
		Name: "add",
		Parameters: []Parameter{
			Parameter{
				IsArray: false,
				Type:    "int32",
			},
			Parameter{
				IsArray: true,
				Type:    "float32",
			},
			Parameter{
				IsArray: false,
				Type:    "address",
			},
			Parameter{
				IsArray: false,
				Type:    "uint32",
			},
		},
	}
	if function.Name != vertex_function.Name {
		t.Errorf("function was incorrect name, got: %s, want: %s.", function.Name, vertex_function.Name)
	}
	if function.Parameters[0].Type != vertex_function.Parameters[0].Type {
		t.Errorf("function was incorrect parameter type index 0 , got: %s, want: %s.", function.Parameters[0].Type, vertex_function.Parameters[0].Type)
	}
	if function.Parameters[0].IsArray {
		t.Errorf("function was incorrect parameter array index 0, got: %t, want: %t.", function.Parameters[0].IsArray, false)
	}
	if function.Parameters[1].Type != vertex_function.Parameters[1].Type {
		t.Errorf("function was incorrect parameter type index 1 , got: %s, want: %s.", function.Parameters[1].Type, vertex_function.Parameters[1].Type)
	}
	if !function.Parameters[1].IsArray {
		t.Errorf("function was incorrect parameter array index 1, got: %t, want: %t.", function.Parameters[1].IsArray, true)
	}
	if function.Parameters[2].Type != vertex_function.Parameters[2].Type {
		t.Errorf("function was incorrect parameter type index 2 , got: %s, want: %s.", function.Parameters[2].Type, vertex_function.Parameters[2].Type)
	}
	if function.Parameters[2].IsArray {
		t.Errorf("function was incorrect parameter array index 2, got: %t, want: %t.", function.Parameters[2].IsArray, false)
	}
	if function.Parameters[3].Type != vertex_function.Parameters[3].Type {
		t.Errorf("function was incorrect parameter type index 3 , got: %s, want: %s.", function.Parameters[3].Type, vertex_function.Parameters[3].Type)
	}
	if function.Parameters[3].IsArray {
		t.Errorf("function was incorrect parameter array index 2, got: %t, want: %t.", function.Parameters[3].IsArray, false)
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
	event := parseEvent("Transfer", params, "line 9")
	vertex_event := Event{
		Name: "Transfer",
		Parameters: []EventParam{
			EventParam{
				Name: "from",
				Type: "address",
			},
			EventParam{
				Name: "to",
				Type: "address",
			},
			EventParam{
				Name: "amount",
				Type: "uint64",
			},
		},
	}
	if event.Name != vertex_event.Name {
		t.Errorf("event was incorrect name, got: %s, want: %s.", event.Name, vertex_event.Name)
	}
	if event.Parameters[0].Name != vertex_event.Parameters[0].Name {
		t.Errorf("event was incorrect parameter type name, got: %s, want: %s.", event.Parameters[0].Name, vertex_event.Parameters[0].Name)
	}
	if event.Parameters[0].Type != vertex_event.Parameters[0].Type {
		t.Errorf("event was incorrect parameter type, got: %s, want: %s.", event.Parameters[0].Type, vertex_event.Parameters[0].Type)
	}
	if event.Parameters[1].Name != vertex_event.Parameters[1].Name {
		t.Errorf("event was incorrect parameter type name, got: %s, want: %s.", event.Parameters[1].Name, vertex_event.Parameters[1].Name)
	}
	if event.Parameters[1].Type != vertex_event.Parameters[1].Type {
		t.Errorf("event was incorrect parameter type, got: %s, want: %s.", event.Parameters[1].Type, vertex_event.Parameters[1].Type)
	}
	if event.Parameters[2].Name != vertex_event.Parameters[2].Name {
		t.Errorf("event was incorrect parameter type name, got: %s, want: %s.", event.Parameters[2].Name, vertex_event.Parameters[2].Name)
	}
	if event.Parameters[2].Type != vertex_event.Parameters[2].Type {
		t.Errorf("event was incorrect parameter type, got: %s, want: %s.", event.Parameters[2].Type, vertex_event.Parameters[2].Type)
	}
}
