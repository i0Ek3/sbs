package sbs

import (
	"testing"

	a "github.com/i0Ek3/asrt"
)

const str = "sbs"

func TestS2B(t *testing.T) {
	got := S2B(str)
	want := []byte(str)

	a.Asrt(t, got, want)
}

func TestB2S(t *testing.T) {
	byt := []byte(str)
	got := B2S(byt)
	want := string(byt)

	a.Asrt(t, got, want)
}
