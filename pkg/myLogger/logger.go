package myLogger

import (
	"fmt"
	"io"
	"sync"
)

const (
	Info     = "info"
	Err      = "error"
	Critical = "critical"
	Route    = "route"
)

type Logger struct {
	outs []io.Writer
	mu   sync.Mutex
}

func NewLogger(outs []io.Writer) *Logger {
	return &Logger{
		outs: outs,
		mu:   sync.Mutex{},
	}
}

func (l *Logger) PrintWithErr(v string) {
	for _, w := range l.outs {
		l.write(Err, fmt.Sprintf("MSG: '%s'", v), w)
	}
}

func (l *Logger) PrintWithInf(v string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, w := range l.outs {
		l.write(Info, v, w)
	}
	//var file string
	//var line int
	//var ok bool
	//
	//_, file, line, ok = runtime.Caller(1)
	//if !ok {
	//	file = "!?"
	//	line = 0
	//}
	//h := l.formatHead(file, line)
	//fmt.Println(h)
	return
}

func (l *Logger) PrintWithCrit(v string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, w := range l.outs {
		l.write(Route, v, w)
	}
}

func (l *Logger) PrintWithRoute(v string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, w := range l.outs {
		l.write(Route, v, w)
	}

}

func (l *Logger) formatHead(file string, line int) string {
	return ""
}

func (l *Logger) write(prefix, message string, w io.Writer) {
	v := fmt.Sprintf("| %s | %s \n", prefix, message)
	w.Write([]byte(v))
	return
}
