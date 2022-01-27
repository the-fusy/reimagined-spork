.PHONY: build

build:
	GOOS=js GOARCH=wasm go build -o ./web/app.wasm ./cmd/weather



