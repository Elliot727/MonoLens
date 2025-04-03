package fileutils

import (
	"fmt"
	"time"
)

type DateExtractor interface {
	ExtractDate(filename string) (time.Time, error)
}

type DefaultDateExtractor struct {
	DateFormat string
	DateLength int
}

func NewDefaultDateExtractor() *DefaultDateExtractor {
	return &DefaultDateExtractor{
		DateFormat: "02-01-2006",
		DateLength: 10,
	}
}

func (de *DefaultDateExtractor) ExtractDate(filename string) (time.Time, error) {
	if len(filename) < de.DateLength {
		return time.Time{}, fmt.Errorf("filename too short to extract date")
	}

	dateStr := filename[:de.DateLength]
	return time.Parse(de.DateFormat, dateStr)
}
