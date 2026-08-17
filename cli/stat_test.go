package cli

import (
	"os"
	"strings"
	"testing"
)

func TestMainStatDocComment(t *testing.T) {
	data, err := os.ReadFile("stat.go")
	if err != nil {
		t.Fatalf("read stat.go: %v", err)
	}
	src := string(data)
	want := "// mainStat is the entry point for stat command."
	if !strings.Contains(src, want) {
		t.Errorf("expected stat.go to contain %q", want)
	}
}
