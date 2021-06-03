package test

import (
	"github.com/EmYiQing/go-sqlmap/api"
	"github.com/EmYiQing/go-sqlmap/input"
	"testing"
)

func TestLess5Error1(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-5/?id=1",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"E"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess5Error2(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-5/?id=2",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"E"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess6Error1(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-6/?id=1",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"E"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess6Error2(t *testing.T) {
	opts := input.Input{
		Beta:      false,
		Url:       VmwareUrl + "/Less-6/?id=2",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"E"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess5Error1Beta(t *testing.T) {
	opts := input.Input{
		Beta:      true,
		Url:       VmwareUrl + "/Less-5/?id=1",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"E"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess5Error2Beta(t *testing.T) {
	opts := input.Input{
		Beta:      true,
		Url:       VmwareUrl + "/Less-5/?id=2",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"E"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess6Error1Beta(t *testing.T) {
	opts := input.Input{
		Beta:      true,
		Url:       VmwareUrl + "/Less-6/?id=1",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"E"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}

func TestLess6Error2Beta(t *testing.T) {
	opts := input.Input{
		Beta:      true,
		Url:       VmwareUrl + "/Less-6/?id=2",
		Database:  "security",
		Table:     "users",
		Columns:   []string{"id", "username", "password"},
		Technique: []string{"E"},
	}
	instance := api.NewScanner(opts)
	instance.Run()
}
