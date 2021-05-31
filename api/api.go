package api

import (
	"github.com/EmYiQing/go-sqlmap/input"
	"github.com/EmYiQing/go-sqlmap/log"
	"github.com/EmYiQing/go-sqlmap/start"
	"github.com/EmYiQing/go-sqlmap/util"
	"os"
)

type Scanner struct {
	Options input.Input
}

func NewScanner(options input.Input) *Scanner {
	return &Scanner{
		Options: options,
	}
}
func (w *Scanner) Run() {
	log.Info("start scan using api...")
	target := util.CheckUrl(w.Options.Url)
	log.Info("target is " + target)
	if !start.DetectAlive(target) {
		os.Exit(-1)
	}
	if start.DetectSafeDogWaf(target) {
		os.Exit(-1)
	}
	start.NewSimpleStarter(target, w.Options)
}

func (w *Scanner) Stop() {
	log.Info("stop scan using api...")
	os.Exit(0)
}
