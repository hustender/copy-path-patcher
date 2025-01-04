package fileops

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func ReplacePath(filePath, oldPath string, newPath string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for reading: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(file)
	for scanner.Scan() {
		updatedLine := strings.ReplaceAll(scanner.Text(), oldPath, newPath)
		_, err := writer.WriteString(updatedLine + "\n")
		if err != nil {
			return fmt.Errorf("failed to write updated line: %v", err)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush content file: %v", err)
	}
	return nil
}

func Copy(path string, destinationPath string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	_, err = io.Copy(destinationFile, file)
	if err != nil {
		return err
	}
	sourceInfo, err := destinationFile.Stat()
	if err != nil {
		return fmt.Errorf("could not get source file info: %v", err)
	}
	err = os.Chmod(path, sourceInfo.Mode())
	if err != nil {
		return fmt.Errorf("could not set file permissions: %v", err)
	}
	return nil
}

func IsTextFile(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 4096)
	for {
		n, err := reader.Read(buf)
		if err != nil && err.Error() != "EOF" {
			return false, err
		}
		if n == 0 {
			break
		}
		if !utf8.Valid(buf[:n]) {
			return false, nil
		}
	}
	return true, nil
}
