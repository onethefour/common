package main

import "github.com/onethefour/common/log"

func main() {
	log.InitLogger("common", true, "debug", "text", "./tmplogs/info.", "./tmplogs/err.")
	log.Info("log.info")
	log.Info("log.err")
}
