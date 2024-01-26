package app

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/android-lewis/note-taker/mocks"
)

func TestStripNewLines(t *testing.T) {
	testString := " # Testing 123 \n## Testing \n### Testing"
	isMatching, err := regexp.MatchString("\n", testString)

	if err != nil || !isMatching {
		t.Errorf("test string not formatted correctly: %s", testString)
	}

	formatted := stripNewLines(testString)
	isMatching, err = regexp.MatchString("\n", formatted)
	if err != nil || isMatching {
		t.Errorf("string not formatted correctly: %s", formatted)
	}
}

func TestGetFirstNLines(t *testing.T) {
	testString := "This is a test string"
	shortString := "This is"
	length := 7

	short := firstN(testString, length)
	if len(short) != length {
		t.Errorf("string incorrect length: %s", short)
	}

	if short != shortString {
		t.Errorf("incorrect string segment returned: %s", short)
	}

	fmt.Println(short)
}

func TestCleanName(t *testing.T) {
	fileContent := "This is a test of file content"
	length := 6
	info := &mocks.MockFileInfo{}

	fileInfo := formatFileInfo(info, []byte(fileContent), length)

	if fileInfo.preview != "This i" || fileInfo.noteId != "1704281311080479" ||
		fileInfo.noteName != "testnote" {
		t.Errorf("file formatting incorrect: %s", fileInfo)
	}

	_, err := time.Parse(time.RFC822, fileInfo.formattedDate)
	if err != nil {
		t.Errorf("time format is incorrect: %s", fileInfo.formattedDate)
	}

}

var fileNameTests = []struct {
	name string
	in   string
	out  string
}{
	{"empty", "", "new_note"},
	{"standard", "note 1", "note1"},
}

func TestGenerateFileName(t *testing.T) {
	for _, test := range fileNameTests {
		t.Run(test.name, func(t *testing.T) {
			// Breakup returned string into key, name and extension
			// Ensure matches returned key, name and extension

			fmt.Println(test)
		})
	}
}
