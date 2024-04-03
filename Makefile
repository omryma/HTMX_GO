.PHONY: all

all:
	templ generate
	go build
	go run .
