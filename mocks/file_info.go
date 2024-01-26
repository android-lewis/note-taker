package mocks

import (
	"os"
	"time"
)

type MockFileInfo struct {
	os.FileInfo
}

func (fi *MockFileInfo) Name() string {
	return "1704281311080479_testnote.md"
}

func (fi *MockFileInfo) ModTime() time.Time {
	return time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
}
