VERSION ?= $$(git describe --tags)

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o out/license

release: build
	tar czvf license.tar.gz out
