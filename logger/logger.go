package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

type LogLevel byte

const (
	LogLevelNone LogLevel = 0
	LogLevelDebug
	LogLevelInfo
	LogLevelWarning
	LogLevelError
	LogLevelCritical
)

type Logger interface {
	SetVerbosity(LogLevel)
	Write(...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warning(...interface{})
	Error(...interface{})
	Critical(...interface{})
}

type logger struct {
	sync.Mutex
	prefix  string
	lvl     LogLevel
	doWrite func(LogLevel, []byte)
}

func (l *logger) append(b []byte, i interface{}) []byte {
	switch i := i.(type) {
	case bool:
		if i {
			return append(b, []byte("true")...)
		}

		return append(b, []byte("false")...)

	case uint8:
	case uint16:
	case uint32:
	case uint64:
		return strconv.AppendUint(b, uint64(i), 10)

	case int8:
	case int16:
	case int32:
	case int64:
		return strconv.AppendInt(b, int64(i), 10)

	case float32:
		return strconv.AppendFloat(b, float64(i), 'f', -1, 32)

	case float64:
		return strconv.AppendFloat(b, i, 'f', -1, 64)

	default:
		return append(b, []byte(fmt.Sprint(i))...)
	}

	return b
}

func (l *logger) writeLvl(lvl LogLevel, s ...interface{}) {
	if l.lvl < lvl {
		return
	}

	// b := make([]byte, 0, 64)
	b := bytes.NewBuffer(nil)

	// b = time.Now().AppendFormat(b, time.RFC3339)

	for _, item := range s {
		fmt.Fprint(b, ' ', item)
		// b = l.append(b, ' ')
		// b = l.append(b, item)
	}

	l.Lock()
	l.doWrite(lvl, b.Bytes())
	l.Unlock()
}

func (l *logger) Write(s ...interface{}) {
	l.writeLvl(LogLevelNone, s...)
}

func (l *logger) SetVerbosity(lvl LogLevel) {
	l.lvl = lvl
}

func (l *logger) Debug(s ...interface{}) {
	l.writeLvl(LogLevelDebug, s...)
}

func (l *logger) Info(s ...interface{}) {
	l.writeLvl(LogLevelInfo, s...)
}

func (l *logger) Warning(s ...interface{}) {
	l.writeLvl(LogLevelWarning, s...)
}

func (l *logger) Error(s ...interface{}) {
	l.writeLvl(LogLevelError, s...)
}

func (l *logger) Critical(s ...interface{}) {
	l.writeLvl(LogLevelCritical, s...)
}

func StdInErr(lvl LogLevel) Logger {
	return &logger{
		lvl: lvl,
		doWrite: func(lvl LogLevel, data []byte) {
			if lvl < LogLevelWarning {
				os.Stdout.Write(data)
				return
			}

			os.Stderr.Write(data)
		},
	}
}

func StdErr(lvl LogLevel) Logger {
	return &logger{
		lvl: lvl,
		doWrite: func(lvl LogLevel, data []byte) {
			os.Stderr.Write(data)
		},
	}
}

func WithWriter(lvl LogLevel, writer io.Writer) Logger {
	return &logger{
		lvl: lvl,
		doWrite: func(lvl LogLevel, data []byte) {
			writer.Write(data)
		},
	}
}
