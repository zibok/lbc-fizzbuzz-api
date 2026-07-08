# lbc-fizzbuzz-api

This project is an API exposing routes to generate FizzBuzz suites and monitor the service.

## Configuration

The executable reads a JSON configuration file. By default it looks for `config.json`; if the file is missing, default values are used.

Example configuration:

```json
{
  "addr": ":8080",
  "maxLimit": 10000
}
```

Configuration fields:

- `addr`: TCP listener address and port. Defaults to `:8080`.
- `maxLimit`: Maximum allowed value for the `/v1/fizzbuzz` `limit` query parameter. Defaults to `10000`.

A sample file is available at `config.example.json`.

## Run the server

```sh
go run ./cmd/api -config config.example.json
```

The server listens on the configured `addr` value.

## Endpoints

### `GET /healthz`

Returns the service health status.

Response:

```json
{"status":"ok"}
```

Example:

```sh
curl "http://localhost:8080/healthz"
```

### `GET /v1/fizzbuzz`

Generates a FizzBuzz sequence.

Query parameters:

- `limit`: Number of values to generate. Defaults to `100`. Must be an integer between `1` and the configured `maxLimit`.
- `firstModulo`: Number that produces `Fizz`. Defaults to `3`. Must be an integer between `1` and `10000`.
- `secondModulo`: Number that produces `Buzz`. Defaults to `5`. Must be an integer between `1` and `10000`.
- `firstWord`: Word to use instead of `Fizz`. Defaults to `Fizz`.
- `secondWord`: Word to use instead of `Buzz`. Defaults to `Buzz`.

A value divisible by both modulo values is returned as the concatenation of `firstWord` and `secondWord`.

Response:

```json
{
  "limit": 5,
  "values": ["1", "2", "Fizz", "4", "Buzz"]
}
```

Example:

```sh
curl "http://localhost:8080/v1/fizzbuzz?limit=15&firstModulo=3&secondModulo=5&firstWord=Fizz&secondWord=Buzz"
```

Invalid query parameters return `400 Bad Request`:

```json
{"error":"limit must be an integer between 1 and 10000"}
```

### `GET /metrics`

Exposes Prometheus metrics for the API.

The response includes `http_request_duration_seconds`, a histogram of response times in seconds labelled by:

- `route`: Matched HTTP route, for example `GET /v1/fizzbuzz`.
- `status_code`: HTTP response status code, for example `200` or `400`.

Example:

```sh
curl "http://localhost:8080/metrics"
```

## Run tests

```sh
go test ./...
```
