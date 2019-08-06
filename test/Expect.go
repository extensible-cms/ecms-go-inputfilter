package test

import "testing"

func ExpectEqual(t *testing.T, prefix string, a interface{}, b interface{}) {
	if a !=b {
		t.Errorf(prefix + " Expected `%v`;  Received `%v`;" , b, a)
	}
}
