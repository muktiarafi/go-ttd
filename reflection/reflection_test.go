package reflection

import (
	"reflect"
	"testing"
)

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
		{
			"Slices",
			[]Profile{
				{40, "Semarang"},
				{22, "Jakarta"},
			},
			[]string{"Semarang", "Jakarta"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "Surabaya"},
				{22, "Malang"},
			},
			[]string{"Surabaya", "Malang"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with cannels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{24, "Semarang"}
			aChannel <- Profile{31, "Cirebon"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Semarang", "Cirebon"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
