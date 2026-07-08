package httpapi

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	server := httptest.NewServer(NewRouter(Config{}, slog.New(slog.NewTextHandler(io.Discard, nil))))
	defer server.Close()

	res, err := http.Get(server.URL + "/healthz")
	if err != nil {
		t.Fatalf("GET /healthz failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusOK)
	}
}

func TestFizzBuzz(t *testing.T) {
	server := httptest.NewServer(NewRouter(Config{}, slog.New(slog.NewTextHandler(io.Discard, nil))))
	defer server.Close()

	res, err := http.Get(server.URL + "/v1/fizzbuzz?limit=5")
	if err != nil {
		t.Fatalf("GET /v1/fizzbuzz failed: %v", err)
	}
	defer res.Body.Close()

	var body fizzbuzzResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if body.Limit != 5 {
		t.Fatalf("limit = %d, want 5", body.Limit)
	}

	want := []string{"1", "2", "Fizz", "4", "Buzz"}
	for i := range want {
		if body.Values[i] != want[i] {
			t.Fatalf("values[%d] = %q, want %q", i, body.Values[i], want[i])
		}
	}
}

func TestFizzBuzzWithCustomWords(t *testing.T) {
	server := httptest.NewServer(NewRouter(Config{}, slog.New(slog.NewTextHandler(io.Discard, nil))))
	defer server.Close()

	res, err := http.Get(server.URL + "/v1/fizzbuzz?limit=6&firstModulo=2&secondModulo=3&firstWord=Foo&secondWord=Bar")
	if err != nil {
		t.Fatalf("GET /v1/fizzbuzz failed: %v", err)
	}
	defer res.Body.Close()

	var body fizzbuzzResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	want := []string{"1", "Foo", "Bar", "Foo", "5", "FooBar"}
	for i := range want {
		if body.Values[i] != want[i] {
			t.Fatalf("values[%d] = %q, want %q", i, body.Values[i], want[i])
		}
	}
}

func TestFizzBuzzUsesConfiguredMaxLimit(t *testing.T) {
	server := httptest.NewServer(NewRouter(Config{MaxLimit: 3}, slog.New(slog.NewTextHandler(io.Discard, nil))))
	defer server.Close()

	res, err := http.Get(server.URL + "/v1/fizzbuzz?limit=4")
	if err != nil {
		t.Fatalf("GET /v1/fizzbuzz failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusBadRequest)
	}

	var body errorResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if body.Error != "limit must be an integer between 1 and 3" {
		t.Fatalf("error = %q, want configured max limit message", body.Error)
	}
}

func TestFizzBuzzRejectsInvalidLimit(t *testing.T) {
	server := httptest.NewServer(NewRouter(Config{}, slog.New(slog.NewTextHandler(io.Discard, nil))))
	defer server.Close()

	res, err := http.Get(server.URL + "/v1/fizzbuzz?limit=nope")
	if err != nil {
		t.Fatalf("GET /v1/fizzbuzz failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusBadRequest)
	}
}

func TestMetricsExposeRouteAndStatusLabels(t *testing.T) {
	server := httptest.NewServer(NewRouter(Config{}, slog.New(slog.NewTextHandler(io.Discard, nil))))
	defer server.Close()

	res, err := http.Get(server.URL + "/v1/fizzbuzz?limit=nope")
	if err != nil {
		t.Fatalf("GET /v1/fizzbuzz failed: %v", err)
	}
	res.Body.Close()

	res, err = http.Get(server.URL + "/metrics")
	if err != nil {
		t.Fatalf("GET /metrics failed: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read metrics response: %v", err)
	}

	metric := `http_request_duration_seconds_count{route="GET /v1/fizzbuzz",status_code="400"}`
	if !strings.Contains(string(body), metric) {
		t.Fatalf("metrics response does not contain %q", metric)
	}
}
