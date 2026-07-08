package httpapi

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/zibok/lbc-fizzbuzz-api/internal/fizzbuzz"
)

type API struct {
	config             Config
	logger             *slog.Logger
	statisticsRecorder StatisticsRecorder
}

type healthResponse struct {
	Status string `json:"status"`
}

type fizzbuzzResponse struct {
	Limit  int      `json:"limit"`
	Values []string `json:"values"`
}

type statisticsRequestResponse struct {
	Limit        int    `json:"limit"`
	FirstModulo  int    `json:"firstModulo"`
	SecondModulo int    `json:"secondModulo"`
	FirstWord    string `json:"firstWord"`
	SecondWord   string `json:"secondWord"`
}

type statisticsResponse struct {
	Request *statisticsRequestResponse `json:"request"`
	Hits    int                        `json:"hits"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func (api API) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{Status: "ok"})
}

func (api API) fizzbuzz(w http.ResponseWriter, r *http.Request) {
	config := fizzbuzz.Config{
		Limit:        100,
		FirstModulo:  3,
		SecondModulo: 5,
		FirstWord:    "Fizz",
		SecondWord:   "Buzz",
	}

	if limitQueryParam := r.URL.Query().Get("limit"); limitQueryParam != "" {
		parsed, err := strconv.Atoi(limitQueryParam)
		if err != nil || parsed < 1 || parsed > api.config.MaxLimit {
			writeJSON(w, http.StatusBadRequest, errorResponse{Error: fmt.Sprintf("limit must be an integer between 1 and %d", api.config.MaxLimit)})
			return
		}
		config.Limit = parsed
	}

	if firstModuloQueryParam := r.URL.Query().Get("firstModulo"); firstModuloQueryParam != "" {
		parsed, err := strconv.Atoi(firstModuloQueryParam)
		if err != nil || parsed < 1 || parsed > 10000 {
			writeJSON(w, http.StatusBadRequest, errorResponse{Error: "firstModulo must be an integer between 1 and 10000"})
			return
		}
		config.FirstModulo = parsed
	}

	if secondModuloQueryParam := r.URL.Query().Get("secondModulo"); secondModuloQueryParam != "" {
		parsed, err := strconv.Atoi(secondModuloQueryParam)
		if err != nil || parsed < 1 || parsed > 10000 {
			writeJSON(w, http.StatusBadRequest, errorResponse{Error: "secondModulo must be an integer between 1 and 10000"})
			return
		}
		config.SecondModulo = parsed
	}

	if firstWordQueryParam := r.URL.Query().Get("firstWord"); firstWordQueryParam != "" {
		config.FirstWord = firstWordQueryParam
	}

	if secondWordQueryParam := r.URL.Query().Get("secondWord"); secondWordQueryParam != "" {
		config.SecondWord = secondWordQueryParam
	}

	api.statisticsRecorder.Record(config)

	writeJSON(w, http.StatusOK, fizzbuzzResponse{
		Limit:  config.Limit,
		Values: fizzbuzz.Generate(config),
	})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func (api API) statistics(w http.ResponseWriter, r *http.Request) {
	config, hits, found := api.statisticsRecorder.MostFrequent()
	if !found {
		writeJSON(w, http.StatusOK, statisticsResponse{Hits: 0})
		return
	}

	writeJSON(w, http.StatusOK, statisticsResponse{
		Request: &statisticsRequestResponse{
			Limit:        config.Limit,
			FirstModulo:  config.FirstModulo,
			SecondModulo: config.SecondModulo,
			FirstWord:    config.FirstWord,
			SecondWord:   config.SecondWord,
		},
		Hits: hits,
	})
}
