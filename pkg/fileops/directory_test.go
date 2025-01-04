package fileops

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetSubFiles(t *testing.T) {
	t.Run("Valid directory returns files", func(t *testing.T) {
		dir := t.TempDir()
		file1, _ := os.CreateTemp(dir, "file1_*.txt")
		file2, _ := os.CreateTemp(dir, "file2_*.txt")
		defer os.Remove(file1.Name())
		defer os.Remove(file2.Name())

		files, err := GetSubFiles(dir)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(files) != 2 {
			t.Errorf("Expected 2 files, got %d", len(files))
		}
	})
	t.Run("Empty directory returns no files", func(t *testing.T) {
		dir := t.TempDir()

		files, err := GetSubFiles(dir)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(files) != 0 {
			t.Errorf("Expected 0 files, got %d", len(files))
		}
	})
	t.Run("Non-existent directory returns error", func(t *testing.T) {
		_, err := GetSubFiles("/non_existent_directory")
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
	t.Run("Directory with subdirectories returns only files", func(t *testing.T) {
		dir := t.TempDir()
		subDir := filepath.Join(dir, "subdir")
		os.Mkdir(subDir, 0755)
		file1, _ := os.CreateTemp(dir, "file1_*.txt")
		file2, _ := os.CreateTemp(subDir, "file2_*.txt")
		defer os.Remove(file1.Name())
		defer os.Remove(file2.Name())

		files, err := GetSubFiles(dir)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(files) != 2 {
			t.Errorf("Expected 2 files, got %d", len(files))
		}
	})
}
