package fileutils

import (
	"path/filepath"
	"testing"
	"time"
)

func TestBuildPath(t *testing.T) {
	pb := DateBasedPathBuilder{}

	basePath := "/mnt/ssd"
	testDate := time.Date(2024, time.April, 3, 0, 0, 0, 0, time.UTC)
	filename := "photo.jpg"

	expectedPath := filepath.Join(basePath, "2024/04/03", filename)
	resultPath := pb.BuildPath(basePath, testDate, filename)

	if resultPath != expectedPath {
		t.Errorf("Expected path: %s, but got: %s", expectedPath, resultPath)
	}
}

func TestBuildPath_DifferentBasePath(t *testing.T) {
	pb := DateBasedPathBuilder{}

	basePath := "/Users/elliot/Pictures"
	testDate := time.Date(2023, time.December, 25, 0, 0, 0, 0, time.UTC)
	filename := "xmas_photo.png"

	expectedPath := filepath.Join(basePath, "2023/12/25", filename)
	resultPath := pb.BuildPath(basePath, testDate, filename)

	if resultPath != expectedPath {
		t.Errorf("Expected path: %s, but got: %s", expectedPath, resultPath)
	}
}

func TestBuildPath_EmptyFilename(t *testing.T) {
	pb := DateBasedPathBuilder{}

	basePath := "/data"
	testDate := time.Date(2025, time.June, 15, 0, 0, 0, 0, time.UTC)
	filename := ""

	expectedPath := filepath.Join(basePath, "2025/06/15", "")
	resultPath := pb.BuildPath(basePath, testDate, filename)

	if resultPath != expectedPath {
		t.Errorf("Expected path: %s, but got: %s", expectedPath, resultPath)
	}
}
