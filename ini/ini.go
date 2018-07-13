package ini

import (
	"bytes"
	"os"
	"regexp"
)

type INI map[string]map[string]string

type token byte

const (
	global token = iota
)

func Parse(data []byte) *INI {
	ini := &INI{}

	for _, line := range bytes.Split(data, []byte("\r\n")) {
		line = bytes.Trim(line, " \t")

		if len(line) < 2 {
			continue
		}
	}

	rx := regexp.MustCompile("^(?:([^=]+)=([^;#]+)|([([^\\]]+)]))")

	for _, line := range rx.FindAllSubmatch(data, -1) {
		os.Setenv(string(line[1]), string(line[2]))
	}

	return ini
}
