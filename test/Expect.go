package test

import "testing"

func ExpectEqual(t *testing.T, a interface{}, b interface{}) {
	if a !=b {
		t.Errorf("Expected `%v`;  Received `%v`;", a, b)
	}
}
