package tools

import (
	"os/exec"
	"testing"
)

func TestCompileClang(t *testing.T) {

	File := "./tests/contract.c"
	option := "add"
	compile := Compile{File}
	wasmFile, nameFile := compile.Clang(option)
	if !(wasmFile == "contract.wasm") {
		t.Errorf("compile fail")
	}
	if !(nameFile == "contract") {
		t.Errorf("compile fail")
	}
	cmd := exec.Command("rm", "-rf", wasmFile)
	_, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("file not found")
	}
}
