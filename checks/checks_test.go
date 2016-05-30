package checks

import (
	"testing"
)

//tests getCmdOption helper function
func TestGetCmdOption(t *testing.T) {
	t.Log("Creating dummy cmd line")
	cmdLine := []string{"dummy", "--opt=1", "--opt2=2"}
	exists, val := getCmdOption(cmdLine, "opt")
	if exists == false {
		t.Errorf("opt1 should exist, got false instead\n")
	}
	if val != "1" {
		t.Errorf("Expected val equal to 1, got %s instead", val)
	}
}
