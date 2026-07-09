package httpapi

import (
	"sync"

	"github.com/zibok/lbc-fizzbuzz-api/internal/fizzbuzz"
)

type StatisticsRecorder interface {
	// Record increments the hit count for the given fizzbuzz.Config.
	Record(config fizzbuzz.Config)
	// MostFrequent returns the fizzbuzz.Config with the highest hit count, the number of hits, and a boolean indicating if any records exist.
	MostFrequent() (fizzbuzz.Config, int, bool)
}

// Here we define an in-memory implementation of the StatisticsRecorder interface.
// In case of multiple instances of the API running, the data will not be shared across instances.
// Then a persitence layer like a database would be needed to share the statistics across instances
// And another implementation of the StatisticsRecorder interface would be needed to use that persistence layer.
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
