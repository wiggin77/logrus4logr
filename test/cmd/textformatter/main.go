package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

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

var (
	queueFullCount       uint32
	targetQueueFullCount uint32
)

func handleLoggerError(err error) {
	panic(err)
}

func main() {
	// create writer target to stdout using Logrus TextFormatter
	var t logr.Target
	filter := &logr.StdFilter{Lvl: logr.Warn, Stacktrace: logr.Error}
	formatter := logrus4logr.FAdapter{Fmtr: &logrus.TextFormatter{}}
	t = target.NewWriterTarget(filter, formatter, os.Stdout, 1000)
	lgr.AddTarget(t)

	wg := sync.WaitGroup{}
	var id int32
	var filterCount int32
	var logCount int32

	runner := func(loops int) {
		defer wg.Done()
		tid := atomic.AddInt32(&id, 1)
		logger := lgr.NewLogger().WithFields(logr.Fields{"id": tid})

		for i := 1; i <= loops; i++ {
			atomic.AddInt32(&filterCount, 2)
			logger.Debug("XXX")
			logger.Trace("XXX")

			lc := atomic.AddInt32(&logCount, 1)
			logger.Warnf("count:%d -- random data: %s", lc, test.StringRnd(10))
			time.Sleep(1 * time.Millisecond)
		}
	}

	start := time.Now()

	for i := 0; i < GOROUTINES; i++ {
		wg.Add(1)
		go runner(LOOPS)
	}
	wg.Wait()

	end := time.Now()
	lgr.NewLogger().Errorf("TextFormatter test ending at %v", end)

	err := lgr.Shutdown()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(atomic.LoadInt32(&logCount), " log entries output.")
	fmt.Println(atomic.LoadInt32(&filterCount), " log entries filtered.")
	fmt.Println(end.Sub(start).String())
}
