package main

import (
	"testing"
)

func TestIsURL(t *testing.T) {
	urls := []struct {
		url    string
		result bool
	}{
		{"hello", false},
		{"/opt", false},
		{"kas-converter_test.go", false},
		{"./kas-converter_test.go", false},
		{"github.com", false},
		{"github.com/chprat", false},
		{"www.github.com", false},
		{"www.github.com/chprat", false},
		{"https://www.github.com", true},
		{"https://www.github.com/chprat", true},
	}

	for _, element := range urls {
		ret := isURL(element.url)
		if ret != element.result {
			t.Fatalf(`isURL("%v") = %v, want %v`, element.url, ret, element.result)
		}
	}
}
