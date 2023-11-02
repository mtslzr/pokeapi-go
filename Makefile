GOCMD=go
GOMOD=${GOCMD} mod
GOTEST=${GOCMD} test
PACKAGE=github.com/mtslzr/pokeapi-go
SCHEMA_DIR=../api-data/data/schema/

CODECOVFLAGS=-coverprofile=coverage.txt -covermode=atomic -coverpkg=${PACKAGE}

all: structs deps test

deps: tidy

test:
	${GOTEST} -v -race ${CODECOVFLAGS} ./...

test-client:
	${GOTEST} -v ./.

tidy:
	${GOMOD} tidy -v

structs: structs/generated.go

structs/generated.go: jsonschemagen schema
	find schema -name '*.json' | xargs jsonschemagen --rootdir=$(PWD) -n structs -u id -u url > structs/generated.go

schema:
	cp -r ${SCHEMA_DIR} schema/

jsonschemagen:
	${GOCMD} install github.com/RyoJerryYu/go-jsonschema/cmd/jsonschemagen@v0.1.1
