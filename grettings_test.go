package modularizacion

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Isaias"
	// Expresion irregular con librería regexp,busca una considencia exacta con name
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Isaias")
	// Si coincide devolvera valor booleano
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Isaias") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
