GOCMD=go
GOMOD=${GOCMD} mod
GOTEST=${GOCMD} test
PACKAGE=github.com/mtslzr/pokeapi-go

CODECOVFLAGS=-coverprofile=coverage.txt -covermode=atomic -coverpkg=${PACKAGE}

all: deps test

deps: tidy vend

test:
	${GOTEST} -v -race ${CODECOVFLAGS} ./...

test-client:
	${GOTEST} -v ./.

tidy:
	${GOMOD} tidy -v

vend:
	${GOMOD} vendor -v