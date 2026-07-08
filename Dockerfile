FROM golang:1.26-alpine AS build

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/lbc-fizzbuzz-api ./cmd/api

FROM scratch

COPY --from=build /out/lbc-fizzbuzz-api /lbc-fizzbuzz-api

USER 65532:65532
EXPOSE 8080

ENTRYPOINT ["/lbc-fizzbuzz-api"]
