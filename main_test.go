package main

import "testing"

func TestReverse(t *testing.T) {
	testValues := []map[string]string{
		{
			"input":  "123456789",
			"expect": "987654321",
		},
		{
			"input":  "qwerty",
			"expect": "ytrewq",
		},
	}
	for _, test := range testValues {
		if result := Reverse(test["input"]); result != test["expect"] {
			t.Error("Input: ", test["input"], "Expect:", test["expect"], "Result:", result)
		}
	}
}
