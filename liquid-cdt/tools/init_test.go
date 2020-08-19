package tools

import (
	"os/exec"
	"testing"
)

func TestCreate(t *testing.T) {
	name := Create("vertex")
	if name != "vertex" {
		t.Errorf("create fail")
	}
	cmd := exec.Command("rm", "-rf", name)
	_, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("file not found")
	}
}
