package fileutils

import (
	"fmt"
	"path/filepath"
	"time"
)

type PathBuilder interface {
	BuildPath(basePath string, date time.Time, filename string) string
}

type DateBasedPathBuilder struct{}

func (pb DateBasedPathBuilder) BuildPath(basePath string, date time.Time, filename string) string {
	datePath := filepath.Join(
		basePath,
		fmt.Sprintf("%d/%02d/%02d", date.Year(), date.Month(), date.Day()),
	)
	return filepath.Join(datePath, filename)
}
