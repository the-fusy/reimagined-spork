.PHONY: go-app

go-app:
	GOOS=js GOARCH=wasm go build -o ./docs/go-app/web/app.wasm ./go-app
	go run ./go-app



