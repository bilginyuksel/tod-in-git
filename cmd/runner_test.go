package cmd

import "testing"

func TestRunPsl_RunPslCommand(t *testing.T) {
	_, err := RunPsl("pwd")
	if err != nil {
		t.Error()
	}
}

func TestRunPsl_RunLinuxCommand(t *testing.T) {
	_, err := RunPsl("ls -l")
	t.Logf("%v", err)
	if err == nil {
		t.Error()
	}
}

func TestRunLin_RunPslCommand(t *testing.T) {
	_, err := RunLx("dir")
	if err == nil {
		t.Error()
	}
}
