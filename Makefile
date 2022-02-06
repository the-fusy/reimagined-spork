.PHONY: go-app vugu

go-app:
	GOOS=js GOARCH=wasm go build -o ./docs/go-app/web/app.wasm ./go-app
	go run ./go-app

vugu:
	go generate ./vugu
	GOOS=js GOARCH=wasm go build -o ./docs/vugu/app.wasm ./vugu



