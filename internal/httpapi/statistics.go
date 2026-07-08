package httpapi

import (
	"sync"

	"github.com/zibok/lbc-fizzbuzz-api/internal/fizzbuzz"
)

type StatisticsRecorder interface {
	Record(config fizzbuzz.Config)
	MostFrequent() (fizzbuzz.Config, int, bool)
}

type InMemoryStatisticsRecorder struct {
	mu   sync.RWMutex
	hits map[fizzbuzz.Config]int
}

func NewInMemoryStatisticsRecorder() *InMemoryStatisticsRecorder {
	return &InMemoryStatisticsRecorder{
		hits: make(map[fizzbuzz.Config]int),
	}
}

func (recorder *InMemoryStatisticsRecorder) Record(config fizzbuzz.Config) {
	recorder.mu.Lock()
	defer recorder.mu.Unlock()

	recorder.hits[config]++
}

func (recorder *InMemoryStatisticsRecorder) MostFrequent() (fizzbuzz.Config, int, bool) {
	recorder.mu.RLock()
	defer recorder.mu.RUnlock()

	var mostFrequent fizzbuzz.Config
	maxHits := 0
	found := false

	for config, hits := range recorder.hits {
		if !found || hits > maxHits {
			mostFrequent = config
			maxHits = hits
			found = true
		}
	}

	return mostFrequent, maxHits, found
}
