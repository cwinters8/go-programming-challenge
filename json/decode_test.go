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

func TestObjectDecode(t *testing.T) {
	got, err := json.Decode([]string{
		`cars/Ford="F-350"`,
		`cars/Dodge="Ram 2500"`,
	})
	if err != nil {
		t.Error("failed to decode slice of strings: " + err.Error())
		return
	}
	want := `{"cars":{"Dodge":"Ram 2500","Ford":"F-350"}}`
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
		return
	}
}

func TestComplexDecode(t *testing.T) {
	// got, err := json.Decode([]string{`cars[0]="Acura"`, `cars[1]="Ford"`, `cars[2]="Chevrolet"`})
	got, err := json.Decode([]string{
		`cars/Ford="F-350"`,
		`cars/Dodge="Ram 2500"`,
		`animals/horses[0]="Percy"`,
		`animals/horses[1]="Quinn"`,
	})
	if err != nil {
		t.Error("failed to decode slice of strings: " + err.Error())
		return
	}
	want := `{"animals":{"horses":["Percy","Quinn"]},"cars":{"Dodge":"Ram 2500","Ford":"F-350"}}`
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
		return
	}
}
