package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	got := Generate(Config{
		Limit:        15,
		FirstModulo:  3,
		SecondModulo: 5,
		FirstWord:    "Fizz",
		SecondWord:   "Buzz",
	})
	want := []string{
		"1", "2", "Fizz", "4", "Buzz",
		"Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Generate() = %#v, want %#v", got, want)
	}
}

func TestGenerateWithCustomConfig(t *testing.T) {
	got := Generate(Config{
		Limit:        10,
		FirstModulo:  2,
		SecondModulo: 3,
		FirstWord:    "Foo",
		SecondWord:   "Bar",
	})
	want := []string{
		"1", "Foo", "Bar", "Foo", "5",
		"FooBar", "7", "Foo", "Bar", "Foo",
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Generate() = %#v, want %#v", got, want)
	}
}

func TestGenerateWithZeroLimit(t *testing.T) {
	got := Generate(Config{
		Limit:        0,
		FirstModulo:  3,
		SecondModulo: 5,
		FirstWord:    "Fizz",
		SecondWord:   "Buzz",
	})
	if len(got) != 0 {
		t.Fatalf("Generate() returned %d values, want none", len(got))
	}
}
