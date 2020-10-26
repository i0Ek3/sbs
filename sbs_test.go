package sbs

import (
	"reflect"
	"testing"
)

const str = "sbs"

func TestS2B(t *testing.T) {
	got := S2B(str)
	want := []byte(str)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func TestB2S(t *testing.T) {
	byt := []byte(str)
	got := B2S(byt)
	want := string(byt)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
