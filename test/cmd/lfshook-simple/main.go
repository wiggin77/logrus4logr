package main

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/wiggin77/logr"
	"github.com/wiggin77/logrus4logr"
)

func main() {
	var lgr = &logr.Logr{}

	// create a Local File System Hook (LFSHook)
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "./info.log",
		logrus.WarnLevel:  "./warn.log",
		logrus.ErrorLevel: "./error.log",
	}
	lfsHook := lfshook.NewHook(pathMap, &logrus.JSONFormatter{})

	// log severity Info or higher.
	filter := &logr.StdFilter{Lvl: logr.Info}

	// create adapter wrapping lfshook.
	target := logrus4logr.NewAdapterTarget(filter, nil, lfsHook, 1000)
	_ = lgr.AddTarget(target)

	// log stuff!
	logger := lgr.NewLogger().WithField("status", "woot!")

	logger.Info("I'm hooked on Logr")
	logger.WithField("code", 501).Error("Request failed")

	_ = lgr.Shutdown()
}
