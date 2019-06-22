all: deps test

deps: tidy vend

test:
	go test -v ./...

tidy:
	go mod tidy -v

vend:
	go mod vendor -v