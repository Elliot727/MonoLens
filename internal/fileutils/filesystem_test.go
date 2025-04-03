package fileutils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadDir(t *testing.T) {
	fs := DefaultFileSystem{}

	tempDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
	}(tempDir)

	tempFile := filepath.Join(tempDir, "testfile.txt")
	err = os.WriteFile(tempFile, []byte("test data"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	entries, err := fs.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("ReadDir failed: %v", err)
	}

	if len(entries) != 1 {
		t.Errorf("Expected 1 file, found %d", len(entries))
	}

	if entries[0].Name() != "testfile.txt" {
		t.Errorf("Expected file 'testfile.txt', found %s", entries[0].Name())
	}
}

func TestMkdirAll(t *testing.T) {
	fs := DefaultFileSystem{}

	tempDir, err := os.MkdirTemp("", "base")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
	}(tempDir)

	newDir := filepath.Join(tempDir, "nested", "dir")

	err = fs.MkdirAll(newDir, 0755)
	if err != nil {
		t.Fatalf("MkdirAll failed: %v", err)
	}

	info, err := os.Stat(newDir)
	if err != nil {
		t.Fatalf("Directory not created: %v", err)
	}

	if !info.IsDir() {
		t.Errorf("Expected directory, found file")
	}
}

func TestRename(t *testing.T) {
	fs := DefaultFileSystem{}

	tempDir, err := os.MkdirTemp("", "rename_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
	}(tempDir)

	oldPath := filepath.Join(tempDir, "oldname.txt")
	newPath := filepath.Join(tempDir, "newname.txt")

	err = os.WriteFile(oldPath, []byte("rename test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	err = fs.Rename(oldPath, newPath)
	if err != nil {
		t.Fatalf("Rename failed: %v", err)
	}

	if _, err := os.Stat(oldPath); !os.IsNotExist(err) {
		t.Errorf("Old file still exists after rename")
	}

	if _, err := os.Stat(newPath); err != nil {
		t.Errorf("New file does not exist after rename")
	}
}
