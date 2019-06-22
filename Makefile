all: deps testall

deps: tidy vend

test:
	go test -v -tags quick ./tests/...

testall:
	go test -v ./tests/...

tidy:
	go mod tidy -v

vend:
	go mod vendor -v