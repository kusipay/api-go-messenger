package output

import (
	"fmt"
	"strings"
)

type Logger interface {
	Verbose(tag, message string)
	Debug(tag, message string)
	Info(tag, message string)
	Warn(tag, message string)
	Error(tag, message string)
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) Verbose(tag, message string) {
	log("[Verbose]", tag, message)
}

func (l *logger) Debug(tag, message string) {
	log("[Debug]", tag, message)
}

func (l *logger) Info(tag, message string) {
	log("[Info]", tag, message)
}

func (l *logger) Warn(tag, message string) {
	log("[Warn]", tag, message)
}

func (l *logger) Error(tag, message string) {
	log("[Error]", tag, message)
}

func log(vars ...string) {
	results := []any{}
	for _, v := range vars {
		results = append(results, strings.ReplaceAll(v, "\n", "\r"))
	}

	fmt.Println(results...)
}
