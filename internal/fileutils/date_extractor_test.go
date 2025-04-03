package fileutils

import (
	"testing"
	"time"
)

func TestExtractDate_ValidFilename(t *testing.T) {
	extractor := NewDefaultDateExtractor()

	filename := "12-04-2024_some_photo.jpg"
	expectedDate, _ := time.Parse("02-01-2006", "12-04-2024")

	date, err := extractor.ExtractDate(filename)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !date.Equal(expectedDate) {
		t.Errorf("Expected date %v, got %v", expectedDate, date)
	}
}

func TestExtractDate_FilenameTooShort(t *testing.T) {
	extractor := NewDefaultDateExtractor()

	filename := "12-04.jpg"
	_, err := extractor.ExtractDate(filename)

	if err == nil {
		t.Fatal("Expected an error for short filename, but got none")
	}
}

func TestExtractDate_InvalidDateFormat(t *testing.T) {
	extractor := NewDefaultDateExtractor()

	filename := "99-99-9999_invalid.jpg"
	_, err := extractor.ExtractDate(filename)

	if err == nil {
		t.Fatal("Expected an error for invalid date format, but got none")
	}
}
