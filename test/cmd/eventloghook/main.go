// +build windows

package main

import (
	"time"
	
	"github.com/Freman/eventloghook"
	"golang.org/x/sys/windows/svc/eventlog"

	"github.com/wiggin77/logr"
	"github.com/wiggin77/logr/test"

	"github.com/wiggin77/logrus4logr"
)

const (
	// GOROUTINES is the number of goroutines to spawn.
	GOROUTINES = 5
	// LOOPS is the number of loops per goroutine.
	LOOPS = 5
)

var lgr = &logr.Logr{
	MaxQueueSize:  1000,
	OnLoggerError: handleLoggerError,
}

func handleLoggerError(err error) {
	panic(err)
}

func main() {
	elog,err := eventlog.Open("Logr Test")
	if err != nil {
		panic(err)
	}
	eventLogHook := eventloghook.NewHook(elog)

	// create writer target to stdout using adapter wrapping the NestedFormatter.
	var t logr.Target
	filter := &logr.StdFilter{Lvl: logr.Info, Stacktrace: logr.Error}

	t = logrus4logr.NewAdapterTarget(filter, nil, eventLogHook, 1000)
	lgr.AddTarget(t)

	cfg := test.DoSomeLoggingCfg{
		Lgr:        lgr,
		Goroutines: GOROUTINES,
		Loops:      LOOPS,
		GoodToken:  "woot!",
		BadToken:   "!!!XXX!!!",
		Lvl:        logr.Error,
		Delay:      time.Millisecond * 1,
	}
	test.DoSomeLogging(cfg)
}
