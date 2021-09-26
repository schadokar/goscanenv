package goscanenv

import "testing"

func TestDefaultScanEnv(t *testing.T) {
	status, err := ScanEnv()

	if err != nil {
		t.Error("Missing Envs")
	}
	t.Log("Success", status)
}

func TestScanEnv(t *testing.T) {
	opt := EnvOptions{
		EnvFile:    ".env",
		ExampleEnv: ".env.example",
		IgnoreEnv:  ".envignore",
	}

	status, err := ScanEnv(opt)

	if err != nil {
		t.Error("Missing Envs")
	}
	t.Log("Success", status)
}
