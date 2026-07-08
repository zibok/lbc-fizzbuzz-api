package httpapi

import (
	"testing"

	"github.com/zibok/lbc-fizzbuzz-api/internal/fizzbuzz"
)

func TestInMemoryStatisticsRecorderMostFrequent(t *testing.T) {
	recorder := NewInMemoryStatisticsRecorder()

	first := fizzbuzz.Config{
		Limit:        5,
		FirstModulo:  3,
		SecondModulo: 5,
		FirstWord:    "Fizz",
		SecondWord:   "Buzz",
	}
	second := fizzbuzz.Config{
		Limit:        6,
		FirstModulo:  2,
		SecondModulo: 3,
		FirstWord:    "Foo",
		SecondWord:   "Bar",
	}

	recorder.Record(first)
	recorder.Record(second)
	recorder.Record(second)

	got, hits, found := recorder.MostFrequent()
	if !found {
		t.Fatal("MostFrequent() found = false, want true")
	}

	if got != second {
		t.Fatalf("MostFrequent() config = %#v, want %#v", got, second)
	}

	if hits != 2 {
		t.Fatalf("MostFrequent() hits = %d, want 2", hits)
	}
}

func TestInMemoryStatisticsRecorderMostFrequentWhenEmpty(t *testing.T) {
	recorder := NewInMemoryStatisticsRecorder()

	_, hits, found := recorder.MostFrequent()
	if found {
		t.Fatal("MostFrequent() found = true, want false")
	}

	if hits != 0 {
		t.Fatalf("MostFrequent() hits = %d, want 0", hits)
	}
}
