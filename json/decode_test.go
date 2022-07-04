package json_test

import (
	"testing"

	"challenge/json"
)

func TestSimpleDecode(t *testing.T) {
	got, err := json.Decode([]string{`car="Acura"`})
	if err != nil {
		t.Error("failed to decode slice of strings: " + err.Error())
		return
	}
	want := `{"car":"Acura"}`
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
		return
	}
}

func TestDecode(t *testing.T) {
	got, err := json.Decode([]string{`cars[0]="Acura"`, `cars[1]="Ford"`, `cars[2]="Chevrolet"`})
	if err != nil {
		t.Error("failed to decode slice of strings: " + err.Error())
		return
	}
	want := `{"cars":["Acura","Ford","Chevrolet"]}`
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
		return
	}
}
