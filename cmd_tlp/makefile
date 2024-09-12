.PHONY: echo build

echo:
	@echo "cmd_tlp"

build:
	go mod tidy
	go build -ldflags "-X main.version=0.0.1" -o cmd_tlp main.go