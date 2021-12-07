package cmd_test

import (
	"bytes"
	"repo/cmd"
	"fmt"
	"io"
	"strings"
	"testing"
)

func (m *promptCapture) ReadAll() string {
	r := strings.ReplaceAll(m.String(), "\033[2K\r", "\n") // clear line
	r = strings.ReplaceAll(r, "\033[1A", "")               // move up
	r = strings.ReplaceAll(r, "\033[1B", "")               // move down
	r = strings.ReplaceAll(r, "\033[?25h", "")             // show cursor
	r = strings.ReplaceAll(r, "\033[?25l", "")             // hide cursor
	return r
}

func testRunClone(t *testing.T) {
	mock := &reposMock{}

	var b bytes.Buffer
	m := promptCapture{&b}

	cmd.CloneRepo(mock, io.NopCloser(strings.NewReader("bet\n")), &m)
	instances := m.ReadAll()
	fmt.Printf("%v\n", instances)

}
