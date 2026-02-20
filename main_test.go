package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Helper: build binaries before testing

func buildBinary(t *testing.T, path, output string) {
	cmd := exec.Command("go", "build", "-o", output, path)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build failed: %v\n%s", err, string(out))
	}
}

// push-swap tests

func TestPushSwap_SortsInput(t *testing.T) {
	buildBinary(t, "./app/push-swap", "push-swap-test")

	defer os.Remove("push-swap-test")

	cmd := exec.Command("./push-swap-test", "3", "2", "1")
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("push-swap execution failed: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 0 {
		t.Fatalf("expected operations output")
	}
}

func TestPushSwap_ErrorOnInvalidInput(t *testing.T) {
	buildBinary(t, "./app/push-swap", "push-swap-test")
	defer os.Remove("push-swap-test")

	cmd := exec.Command("./push-swap-test", "1", "a")
	err := cmd.Run()

	if err == nil {
		t.Fatalf("expected error exit code")
	}
}

// checker tests

func TestChecker_OK(t *testing.T) {
	buildBinary(t, "./app/checker", "checker-test")
	defer os.Remove("checker-test")

	cmd := exec.Command("./checker-test", "2", "1")
	cmd.Stdin = strings.NewReader("sa\n")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		t.Fatalf("checker failed: %v", err)
	}

	if strings.TrimSpace(out.String()) != "OK" {
		t.Fatalf("expected OK, got %s", out.String())
	}
}

func TestChecker_KO(t *testing.T) {
	buildBinary(t, "./app/checker", "checker-test")
	defer os.Remove("checker-test")

	cmd := exec.Command("./checker-test", "2", "1")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		t.Fatalf("checker failed: %v", err)
	}

	if strings.TrimSpace(out.String()) != "KO" {
		t.Fatalf("expected KO, got %s", out.String())
	}
}

func TestChecker_InvalidInstruction(t *testing.T) {
	buildBinary(t, "./app/checker", "checker-test")
	defer os.Remove("checker-test")

	cmd := exec.Command("./checker-test", "1", "2")
	cmd.Stdin = strings.NewReader("invalid\n")

	err := cmd.Run()

	if err == nil {
		t.Fatalf("expected error exit for invalid instruction")
	}
}

// Full integration test: push-swap | checker

func TestPushSwapAndChecker_Integration(t *testing.T) {
	buildBinary(t, "./app/push-swap", "push-swap-test")
	buildBinary(t, "./app/checker", "checker-test")

	defer os.Remove("push-swap-test")
	defer os.Remove("checker-test")

	input := "4 3 2 1"

	push := exec.Command("./push-swap-test", input)
	pushOut, err := push.Output()
	if err != nil {
		t.Fatalf("push-swap failed: %v", err)
	}

	check := exec.Command("./checker-test", input)
	check.Stdin = bytes.NewReader(pushOut)

	var result bytes.Buffer
	check.Stdout = &result

	if err := check.Run(); err != nil {
		t.Fatalf("checker failed: %v", err)
	}

	if strings.TrimSpace(result.String()) != "OK" {
		t.Fatalf("integration failed, expected OK, got %s", result.String())
	}
}