package main

import (
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/wiggin77/logr"
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
	// create a Local Filesystem Hook
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "./info.log",
		logrus.WarnLevel:  "./warn.log",
		logrus.ErrorLevel: "./error.log",
	}

	lfsHook := lfshook.NewHook(pathMap, &logrus.JSONFormatter{})
	filter := &logr.StdFilter{Lvl: logr.Info, Stacktrace: logr.Error}

	// create adapter wrapping lfshook.
	target := logrus4logr.NewAdapterTarget(filter, nil, lfsHook, 1000)
	lgr.AddTarget(target)

	cfg := test.DoSomeLoggingCfg{
		Lgr:        lgr,
		Goroutines: 10,
		Loops:      50,
		GoodToken:  "woot!",
		BadToken:   "!!!XXX!!!",
		Lvl:        logr.Error,
		Delay:      time.Millisecond * 1,
	}
	test.DoSomeLogging(cfg)
}
