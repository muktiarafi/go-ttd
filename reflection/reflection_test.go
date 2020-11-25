package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Bambank"},
			[]string{"Bambank"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Paijo", "Solo"},
			[]string{"Paijo", "Solo"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Maul", 20},
			[]string{"Maul"},
		},
		{
			"Nested fields",
			Person{
				"Jojo",
				Profile{20, "Cirebon"},
			},
			[]string{"Jojo", "Cirebon"},
		},
		{
			"Ponters to things",
			&Person{
				"Jojo",
				Profile{20, "Cirebon"},
			},
			[]string{"Jojo", "Cirebon"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
