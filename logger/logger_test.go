package logger

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkFprintf(b *testing.B) {
	buf := &bytes.Buffer{}

	for n := 0; n < b.N; n++ {
		buf.Reset()
		fmt.Fprintf(buf, "%s %f %d %d %d %d %d %d", "string", 2.22, 3, 4, 5, 6, 7, 8)
	}
}

func BenchmarkFprint(b *testing.B) {
	buf := &bytes.Buffer{}

	for n := 0; n < b.N; n++ {
		buf.Reset()
		fmt.Fprint(buf, "string", 2.22, 3, 4, 5, 6, 7, 8)
	}
}

func BenchmarkLog(b *testing.B) {
	buf := &bytes.Buffer{}
	l := WithWriter(LogLevelNone, buf)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		l.Write("string", 2.22, 3, 4, 5, 6, 7, 8)
	}
}

func BenchmarkLogger(b *testing.B) {
	buf := &bytes.Buffer{}
	l := log.New(buf, "none", log.Ldate|log.Ltime)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		l.Print("string", 2.22, 3, 4, 5, 6, 7, 8)
	}
}

func BenchmarkJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strings.Join([]string{"abc", "def"}, "/")
	}
}
func BenchmarkFormatJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Sprint("abc", "/", "def")
	}
}

func BenchmarkFormatInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strconv.FormatInt(int64(32), 10)
	}
}
func BenchmarkSprintInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Sprint(32)
	}
}
