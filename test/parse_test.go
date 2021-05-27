package test

import (
	"fmt"
	"github.com/EmYiQing/go-sqlmap/parse"
	"testing"
)

func TestParse(t *testing.T) {
	test := parse.RequestParse("../http.txt")
	fmt.Println(test.Cookie)
	fmt.Println(test.Data)
	fmt.Println(test.Headers)
	fmt.Println(test.Method)
	fmt.Println(test.Path)
}
