package tools

import "testing"

func TestCheckAllowFucntion(t *testing.T) {
	checker := checkFunction("test1")
	if checker {
		t.Errorf("function was not allow, got: %t, want: %t.", checker, false)
	}
	checker = checkFunction("chain_event_emit")
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
