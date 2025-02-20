package test

import (
	"github.com/EmYiQing/go-sqlmap/api"
	"github.com/EmYiQing/go-sqlmap/input"
	"testing"
)

func TestLess1Union1(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-1/?id=1",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"U"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess1Union2(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-1/?id=2",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"U"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess2Union1(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-2/?id=1",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"U"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess2Union2(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-2/?id=2",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"U"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess3Union1(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-3/?id=1",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"U"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess3Union2(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-3/?id=2",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"U"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess4Union1(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-4/?id=1",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"U"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess4Union2(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-4/?id=2",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"U"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}
