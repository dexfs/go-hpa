package main

import (
	"testing"
)

func TestRaiz(t *testing.T) {
	result := squareroot()
	if result != "Code.education Rocks!" {
		t.Error("Opa! o resultado deveria ser Code.education Rocks2!", result)
	}
}