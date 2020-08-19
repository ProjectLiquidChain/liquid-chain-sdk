package tools

import "testing"

func TestCheckAllowFucntion(t *testing.T) {
	var AllowFunctionEnv = []string{"chain_print_bytes", "chain_event_emit", "chain_get_caller", "chain_get_creator",
		"chain_invoke", "chain_get_owner", "chain_method_bind", "chain_arg_size_get", "chain_arg_size_set", "get_mean",
		"sum_of_squares", "sqroot", "get_average", "address_xor", "chain_block_height", "chain_block_time", "chain_storage_set",
		"chain_storage_size_get", "chain_storage_get"}
	checker := checkAllowFunction("test1", AllowFunctionEnv)
	if checker {
		t.Errorf("function was not allow, got: %t, want: %t.", checker, false)
	}
	checker = checkAllowFunction("torage_size_get", AllowFunctionEnv)
	if checker {
		t.Errorf("function was allow, got: %t, want: %t.", checker, false)
	}
	checker = checkAllowFunction("chain_storage_size_get", AllowFunctionEnv)
	if !checker {
		t.Errorf("function was allow, got: %t, want: %t.", checker, true)
	}
	checker = checkAllowFunction("chain_storage_get", AllowFunctionEnv)
	if !checker {
		t.Errorf("function was allow, got: %t, want: %t.", checker, true)
	}
	checker = checkAllowFunction("chain_storage_set", AllowFunctionEnv)
	if !checker {
		t.Errorf("function was allow, got: %t, want: %t.", checker, true)
	}
	checker = checkAllowFunction("chain_get_caller", AllowFunctionEnv)
	if !checker {
		t.Errorf("function was allow, got: %t, want: %t.", checker, true)
	}
	checker = checkAllowFunction("chain_get_creator", AllowFunctionEnv)
	if !checker {
		t.Errorf("function was allow, got: %t, want: %t.", checker, true)
	}
	checker = checkAllowFunction("chain_get_owner", AllowFunctionEnv)
	if !checker {
		t.Errorf("function was allow, got: %t, want: %t.", checker, true)
	}
	checker = checkAllowFunction("chain_event_emit", AllowFunctionEnv)
	if !checker {
		t.Errorf("function was allow, got: %t, want: %t.", checker, true)
	}
}

func TestCheckAllowEvent(t *testing.T) {
	checker := checkEvent("Test", []string{"Transfer", "Mint"})
	if checker {
		t.Errorf("Event was not allow, got: %t, want: %t.", checker, false)
	}
	checker = checkEvent("Transfer", []string{"Transfer", "Mint"})
	if !checker {
		t.Errorf("Event was allow, got: %t, want: %t.", checker, true)
	}
}
func TestGetImportFunctions(t *testing.T) {
	import_function := getImportFunction("./tests/vertex_contract.wasm")
	if !(import_function[0] == "chain_get_caller") {
		t.Errorf("funtion was not found, got: %t, want: %t.", false, true)
	}
	if !(import_function[1] == "chain_get_creator") {
		t.Errorf("function was not found, got: %t, want: %t.", false, true)
	}
	if !(import_function[2] == "chain_storage_set") {
		t.Errorf("function was not found, got: %t, want: %t.", false, true)
	}
	if !(import_function[3] == "chain_storage_size_get") {
		t.Errorf("function was not found, got: %t, want: %t.", false, true)
	}
	if !(import_function[4] == "chain_storage_get") {
		t.Errorf("funtion was not found, got: %t, want: %t.", false, true)
	}
	if !(import_function[5] == "Mint") {
		t.Errorf("function was not found, got: %t, want: %t.", false, true)
	}
	if !(import_function[6] == "Transfer") {
		t.Errorf("funtion was not found, got: %t, want: %t.", false, true)
	}
}

func TestGetExportFunctions(t *testing.T) {
	export_function := getExportFunction("./tests/vertex_contract.wasm")
	if !(export_function[3] == "set_owner") {
		t.Errorf("funtion was not found, got: %t, want: %t.", false, true)
	}
	if !(export_function[4] == "pause") {
		t.Errorf("function was not found, got: %t, want: %t.", false, true)
	}
	if !(export_function[5] == "unpause") {
		t.Errorf("function was not found, got: %t, want: %t.", false, true)
	}
	if !(export_function[6] == "set_owner_to_creator") {
		t.Errorf("function was not found, got: %t, want: %t.", false, true)
	}
	if !(export_function[7] == "change_balance") {
		t.Errorf("funtion was not found, got: %t, want: %t.", false, true)
	}
	if !(export_function[8] == "mint") {
		t.Errorf("function was not found, got: %t, want: %t.", false, true)
	}
	if !(export_function[9] == "get_balance") {
		t.Errorf("funtion was not found, got: %t, want: %t.", false, true)
	}
	if !(export_function[10] == "transfer") {
		t.Errorf("funtion was not found, got: %t, want: %t.", false, true)
	}
}

func TestCheckImportFunctions(t *testing.T) {
	check := CheckImportFunction("./tests/vertex_contract.wasm", []string{"Mint", "Transfer"})
	if !check {
		t.Errorf("funtion was not found, got: %t, want: %t.", false, true)
	}
	check = CheckImportFunction("./tests/vertex_contract.wasm", []string{"Mint"})
	if !check {
		t.Errorf("funtion was not found, got: %t, want: %t.", false, true)
	}
}
