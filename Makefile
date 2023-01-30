.PHONY: build, clean, run, first, second, third

APP_BIN = build/csvreader.exe

all: build

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) cmd/main.go

first:
	./build/csvreader.exe data/first.csv

second:
	./build/csvreader.exe data/second.csv

third:
	./build/csvreader.exe data/third.csv

clean:
	rm -rf build || true

test:
	go test ./...