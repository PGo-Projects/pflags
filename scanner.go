package pflags

import (
	"bufio"
	"io"
)

type scanner struct {
	originalScanner *bufio.Scanner
	lineNum         int

	peeked        bool
	linePeeked    string
	lineNumPeeked int
}

func NewScanner(reader io.Reader) *scanner {
	return &scanner{
		originalScanner: bufio.NewScanner(reader),
		lineNum:         0,

		peeked:        false,
		linePeeked:    "",
		lineNumPeeked: 0,
	}
}

func (s *scanner) Scan() bool {
	return s.originalScanner.Scan()
}

func (s *scanner) Text() (string, int) {
	var line string
	var lineNum int
	if s.peeked {
		line = s.linePeeked
		lineNum = s.lineNumPeeked

		s.lineNum = lineNum
	} else {
		line = s.originalScanner.Text()
		lineNum = s.lineNum + 1

		s.lineNum += 1
	}
	return line, lineNum
}

func (s *scanner) Peek() string {
	s.peeked = true
	s.linePeeked = s.originalScanner.Text()
	s.lineNumPeeked = s.lineNum + 1

	return s.linePeeked
}
