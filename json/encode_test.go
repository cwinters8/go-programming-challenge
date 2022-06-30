package json_test

import (
	"fmt"
	"testing"

	"challenge/json"
)

func TestSimpleEncode(t *testing.T) {
	got, err := json.Encode(`{"car": "Acura"}`)
	if err != nil {
		t.Error("encoding failed: " + err.Error())
	}
	want := []string{`car="Acura"`}
	msg := fmt.Sprintf("got: %v, want: %v", got, want)
	if len(got) != len(want) {
		t.Error(msg)
	}
	for i, v := range want {
		if v != got[i] {
			t.Error(msg)
		}
	}
}

func TestEmptyEncode(t *testing.T) {
	got, err := json.Encode("{}")
	if err != nil {
		t.Error("encoding failed: " + err.Error())
	}
	if len(got) != 0 {
		t.Errorf("slice should be empty. got: %v", got)
	}
}
