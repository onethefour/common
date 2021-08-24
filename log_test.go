package main

import (
	"testing"

	"github.com/onethefour/common/log"
)

func Test_log(t *testing.T) {
	log.InitLogger("common", true, "info", "text", "./tmplogs/info", "./tmplogs/err")
	log.Debug("log.Debug")
	log.Info("log.info")
	log.Error("log.err")
}
