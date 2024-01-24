.PHONY: build
build:
	@echo Building executable binary
	mkdir -p build
	go build -o ./build .
