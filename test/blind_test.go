package test

import (
	"github.com/EmYiQing/go-sqlmap/api"
	"github.com/EmYiQing/go-sqlmap/input"
	"testing"
)

func TestLess8Bool1(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-8/?id=1",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"B"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess8Bool2(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-8/?id=2",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"B"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess9Bool1(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-9/?id=1",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"B"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess9Bool2(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-9/?id=2",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"B"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess10Bool1(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-10/?id=1",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"B"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess10Bool2(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-10/?id=2",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"B"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}
