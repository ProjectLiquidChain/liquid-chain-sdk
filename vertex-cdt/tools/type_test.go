package tools

import "testing"

func TestValidateType(t *testing.T) {
	type_128 := validateType("uint128")
	if type_128 {
		t.Errorf("type was incorrect, got: %t, want: %t.", true, false)
	}
	type_64 := validateType("uint64")
	if !type_64 {
		t.Errorf("type was incorrect, got: %t, want: %t.", false, true)
	}
	type_address := validateType("address")
	if !type_address {
		t.Errorf("type was incorrect, got: %t, want: %t.", false, true)
	}
}
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
