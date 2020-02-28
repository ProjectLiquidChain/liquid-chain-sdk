package tools

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
	PlArray string = "plarray"
)

/*
	function checkAllowType : check vertex VM type.
	params:
		- atype: c or c++ type
*/
func validateType(Type string) bool {
	switch Type {
	case Int8, Uint8, Int16, Uint16, Int32, Uint32, Int64, Uint64, Float32, Float64, Address, PlArray:
		return true
	default:
		return false
	}
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
