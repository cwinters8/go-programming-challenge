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
		return
	}
	want := []string{`car="Acura"`}
	msg := fmt.Sprintf("got: %v, want: %v", got, want)
	if len(got) != len(want) {
		t.Error(msg)
		return
	}
	for i, v := range want {
		if v != got[i] {
			t.Error(msg)
			return
		}
	}
}

func TestEncode(t *testing.T) {
	got, err := json.Encode(`
		{
			"age": 35,
			"weight": 165.5,
			"smoker": false,
			"hobbies": ["horses", "dogs", "comics", "programming"],
			"animal_counts": {
				"horses": 4,
				"dogs": 3,
				"cats": 3
			},
			"animals": {
				"horses": ["Percy", "Quinn", "Bravo", "Rhea"],
				"dogs": ["Four", "Hush", "Riley"],
				"cats": ["Fang", "Minerva", "Charlie"]
			}
		}`)
	if err != nil {
		t.Error("encoding failed: " + err.Error())
		return
	}
	want := []string{
		"age=35",
		"weight=165.5",
		"smoker=false",
		`hobbies[0]="horses"`,
		`hobbies[1]="dogs"`,
		`hobbies[2]="comics"`,
		`hobbies[3]="programming"`,
		`animal_counts/horses=4`,
		`animal_counts/dogs=3`,
		`animal_counts/cats=3`,
		`animals/horses[0]="Percy"`,
		`animals/horses[1]="Quinn"`,
		`animals/horses[2]="Bravo"`,
		`animals/horses[3]="Rhea"`,
		`animals/dogs[0]="Four"`,
		`animals/dogs[1]="Hush"`,
		`animals/dogs[2]="Riley"`,
		`animals/cats[0]="Fang"`,
		`animals/cats[1]="Minerva"`,
		`animals/cats[2]="Charlie"`,
	}
	msg := fmt.Sprintf("got: %v, want: %v", got, want)
	if len(got) != len(want) {
		t.Error(msg)
		return
	}

	for _, v := range want {
		found := false
		for _, val := range got {
			if v == val {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("value %s missing. got: %v", v, got)
		}
	}
}

func TestEmptyEncode(t *testing.T) {
	got, err := json.Encode("{}")
	if err != nil {
		t.Error("encoding failed: " + err.Error())
		return
	}
	if len(got) != 0 {
		t.Errorf("slice should be empty. got: %v", got)
		return
	}
}
