package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wiggin77/logr"
	"github.com/wiggin77/logr/target"
	"github.com/wiggin77/logr/test"

	"github.com/wiggin77/logrus4logr"
)

const (
	// GOROUTINES is the number of goroutines to spawn.
	GOROUTINES = 10
	// LOOPS is the number of loops per goroutine.
	LOOPS = 1000
)

var lgr = &logr.Logr{
	MaxQueueSize:  1000,
	OnLoggerError: handleLoggerError,
}

func handleLoggerError(err error) {
	panic(err)
}

func main() {
	// create a Logrus TextFormatter with whatever settings you prefer.
	logrusFormatter := &logrus.TextFormatter{
		// settings...
	}

	// create writer target to stdout using adapter wrapping the Logrus TextFormatter.
	var t logr.Target
	filter := &logr.StdFilter{Lvl: logr.Info, Stacktrace: logr.Error}
	formatter := &logrus4logr.FAdapter{Fmtr: logrusFormatter}
	t = target.NewWriterTarget(filter, formatter, os.Stdout, 1000)
	lgr.AddTarget(t)

	test.DoSomeLogging(lgr, GOROUTINES, LOOPS, "good", "XXX")
}
