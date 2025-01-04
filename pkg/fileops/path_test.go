package fileops

import "testing"

func TestIsValidPath(t *testing.T) {
	t.Run("Empty path returns error", func(t *testing.T) {
		got, err := IsValidPath("")
		if got != false || err == nil {
			t.Errorf("Expected false and error, got %v and %v", got, err)
		}
	})
	t.Run("Mixed path separators return error", func(t *testing.T) {
		got, err := IsValidPath("folder\\subfolder/file")
		if got != false || err == nil {
			t.Errorf("Expected false and error, got %v and %v", got, err)
		}
	})
	t.Run("No valid path separator returns error", func(t *testing.T) {
		got, err := IsValidPath("invalidpath")
		if got != false || err == nil {
			t.Errorf("Expected false and error, got %v and %v", got, err)
		}
	})
	t.Run("Valid Unix path returns true", func(t *testing.T) {
		got, err := IsValidPath("/valid/unix/path")
		if got != true || err != nil {
			t.Errorf("Expected true and no error, got %v and %v", got, err)
		}
	})
	t.Run("Valid Windows path returns true", func(t *testing.T) {
		got, err := IsValidPath("C:\\valid\\windows\\path")
		if got != true || err != nil {
			t.Errorf("Expected true and no error, got %v and %v", got, err)
		}
	})
}
