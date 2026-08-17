package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMainDeleteComment(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("delete.go"))
	if err != nil {
		t.Fatalf("read delete.go: %v", err)
	}
	content := string(data)
	if !strings.Contains(content, "// mainDelete is the entry point for delete command.") {
		t.Errorf("expected comment identifying mainDelete as delete command")
	}
	if strings.Contains(content, "// mainDelete is the entry point for get command.") {
		t.Errorf("found outdated comment identifying mainDelete as get command")
	}
}
