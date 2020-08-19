package tools

import "strings"

const (
	Uint8   string = "uint8"
	Int8    string = "int8"
	Uint16  string = "uint16"
	Int16   string = "int16"
	Uint32  string = "uint32"
	Int32   string = "int32"
	Uint64  string = "uint64"
	Int64   string = "int64"
	Float32 string = "float32"
	Float64 string = "float64"
	Address string = "address"
	Event   string = "Event"
	Array   string = "array"
	Pointer string = "pointer"
	LpArray string = "lparray"
)

/*
	function checkAllowType : check vertex VM type.
	params:
		- atype: c or c++ type
*/
func validateType(Type string) bool {
	if strings.Contains(Type, "[]") {
		Type = Type[:len(Type)-2]
	}
	switch Type {
	case Int8, Uint8, Int16, Uint16, Int32, Uint32,
		Int64, Uint64, Float32, Float64, Address, LpArray:
		return true
	default:
		return false
	}
}

// convert from c,c++ type to assembly vertex type
func convertType(Type string) string {
	switch Type {
	case "float":
		return Float32
	case "double":
		return Float64
	case "signed-char":
		return Int8
	case "char":
		return Int8
	case "unsigned-char":
		return Uint8
	case "short":
		return Int16
	case "unsigned-short":
		return Uint16
	case "int":
		return Int32
	case "unsigned-int":
		return Uint32
	case "unsigned-long":
		return Uint32
	case "long-long":
		return Int64
	case "unsigned-long-long":
		return Uint64
	default:
		return Type
	}
}
func convertRustType(Type string) string {
	switch Type {
	case "f32":
		return Float32
	case "f64":
		return Float64
	case "i8":
		return Int8
	case "u8":
		return Uint8
	case "i16":
		return Int16
	case "u16":
		return Uint16
	case "i32":
		return Int32
	case "u32":
		return Uint32
	case "i64":
		return Int64
	case "u64":
		return Uint64
	default:
		return Type
	}
}
