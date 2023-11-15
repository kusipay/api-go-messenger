package output

import (
	"fmt"
	"strings"
)

type Logger interface {
	Verbose(tag string, message ...string)
	Debug(tag string, message ...string)
	Info(tag string, message ...string)
	Warn(tag string, message ...string)
	Error(tag string, message ...string)
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) Verbose(tag string, message ...string) {
	log("[Verbose]", tag, message...)
}

func (l *logger) Debug(tag string, message ...string) {
	log("[Debug]", tag, message...)
}

func (l *logger) Info(tag string, message ...string) {
	log("[Info]", tag, message...)
}

func (l *logger) Warn(tag string, message ...string) {
	log("[Warn]", tag, message...)
}

func (l *logger) Error(tag string, message ...string) {
	log("[Error]", tag, message...)
}

func log(level, tag string, vars ...string) {
	results := []any{level, tag, "\r"}
	for _, v := range vars {
		results = append(results, strings.ReplaceAll(v, "\n", "\r"))
	}

	fmt.Println(results...)
}
