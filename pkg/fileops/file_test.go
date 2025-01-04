package fileops

import (
	"bytes"
	"crypto/sha256"
	"io"
	"os"
	"testing"
)

func TestIsTextFile(t *testing.T) {
	srcFile, _ := os.CreateTemp("", "valid_text_file_src_*.txt")
	defer os.Remove(srcFile.Name())

	srcContent := []byte("This is a sample content for testing.")
	srcFile.Write(srcContent)
	srcFile.Close()

	t.Run("Valid Text File", func(t *testing.T) {
		got, _ := IsTextFile(srcFile.Name())
		if got != true {
			t.Errorf("Expected %v, got %v", true, got)
		}
	})

	srcFile, _ = os.CreateTemp("", "invalid_text_file_src_*.txt")
	defer os.Remove(srcFile.Name())

	srcFile.Write([]byte{0x00, 0xFF, 0xFE, 0xFD})
	srcFile.Close()

	t.Run("Invalid Text File", func(t *testing.T) {
		got, _ := IsTextFile(srcFile.Name())
		if got != false {
			t.Errorf("Expected %v, got %v", false, got)
		}
	})
	t.Run("Invalid Path returns error", func(t *testing.T) {
		_, err := IsTextFile("i_do_not_exist.txt")
		if err == nil {
			t.Errorf("Expected an error, got %v", err)
		}
	})
}

func TestCopy(t *testing.T) {
	srcFile, _ := os.CreateTemp("", "valid_text_file_src_*.txt")
	defer os.Remove(srcFile.Name())

	srcContent := []byte("This is a sample content for testing.")
	srcFile.Write(srcContent)
	srcFile.Close()

	t.Run("Copy file", func(t *testing.T) {
		destFile, _ := os.CreateTemp("", "valid_text_file_dest_*.txt")
		defer os.Remove(destFile.Name())

		err := Copy(srcFile.Name(), destFile.Name())
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		srcFileReopen, _ := os.Open(srcFile.Name())
		defer srcFileReopen.Close()

		destFileReopen, _ := os.Open(destFile.Name())
		defer destFileReopen.Close()

		srcHash := sha256.New()
		destHash := sha256.New()

		io.Copy(srcHash, srcFileReopen)
		io.Copy(destHash, destFileReopen)

		if !bytes.Equal(srcHash.Sum(nil), destHash.Sum(nil)) {
			t.Errorf("Original and destination files are not the same")
		}
	})
	t.Run("Source file does not exist", func(t *testing.T) {
		err := Copy("non_existent_file.txt", "destination.txt")
		if err == nil {
			t.Errorf("Expected an error, got %v", err)
		}
	})
	t.Run("Destination file cannot be created", func(t *testing.T) {
		err := Copy(srcFile.Name(), "/invalid_path/destination.txt")
		if err == nil {
			t.Errorf("Expected an error, got %v", err)
		}
	})
}
