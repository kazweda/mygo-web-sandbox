.PHONY: build generate clean

generate:
	templ generate

build: generate
	go run ./cmd/build

clean:
	rm -rf dist
