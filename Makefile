.PHONY: build, clean, run, test

APP_BIN = build/app

all: build run

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) cmd/main.go

clean:
	rm -rf build || true

test:
	go test ./...