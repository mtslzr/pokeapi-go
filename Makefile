GOCMD=go
GOMOD=${GOCMD} mod
GOTEST=${GOCMD} test

CODECOVFLAGS=-coverprofile=coverage.txt -covermode=atomic -coverpkg=./...

all: deps test

deps: tidy vend

test:
	${GOTEST} -v -race ${CODECOVFLAGS} ./tests/...

tidy:
	${GOMOD} tidy -v

vend:
	${GOMOD} vendor -v