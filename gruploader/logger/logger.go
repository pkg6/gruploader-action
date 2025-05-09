package logger

import (
	"fmt"
	"github.com/gookit/color"
	"log"
)

// Println execs Println on writer
func Println(out ...any) {
	log.Println(out...)
}

// Printf execs Printf on writer
func Printf(format string, a ...any) {
	log.Printf(format, a...)
}

// Error error output
func Error(err error) {
	log.Printf("%v\n", color.New(color.BgRed, color.FgWhite).Sprintf("error: %v", err))
}

func Fatalf(err error) {
	log.Fatalf("%v\n", color.New(color.BgRed, color.FgWhite).Sprintf("error: %v", err))
}

// Warning warning message
func Warning(out ...any) {
	log.Printf(color.New(color.Yellow).Sprint(out...))
}

func WarningF(format string, a ...any) {
	Warning(fmt.Sprintf(format, a...))
}

// Success success message
func Success(out ...any) {
	log.Printf(color.New(color.Green).Sprint(out...))
}

func InfoF(format string, a ...any) {
	Info(fmt.Sprintf(format, a...))
}

// Info info message
func Info(out ...any) {
	log.Printf(color.New(color.Cyan).Sprint(out...))
}
