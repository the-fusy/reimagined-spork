.PHONY: build

build:
	GOOS=js GOARCH=wasm go build -o ./docs/web/app.wasm ./cmd/weather
	go run ./cmd/weather



