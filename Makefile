VERSION ?= $$(git describe --tags)

setup:
	go get github.com/jteeuwen/go-bindata/...

build:
	go-bindata templates/
	go build -ldflags "-X main.Version=$(VERSION)" -o out/license

release: build
	tar czvf license.tar.gz out
