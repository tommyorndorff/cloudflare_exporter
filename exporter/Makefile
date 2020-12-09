all: full

fmt:
	gofmt -e -s -w *.go

get:
	GOPATH="$$PWD" go get github.com/prometheus/client_golang/prometheus
	GOPATH="$$PWD" go get github.com/cloudflare/cloudflare-go

build:
	GOPATH="$$PWD" CGO_ENABLED=0 go build .

full: clean fmt get build

fast: fmt build

clean:
	rm -f exporter
	rm -rf src/ pkg/
