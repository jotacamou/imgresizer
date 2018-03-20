# Package:   jotacamou/imgresizer
# Build:     docker

# Define at release time
VERSION?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || \
			echo v0)

OWNER?=   jotacamou
PROGRAM?= imgresizer
COMMIT?=  $(shell git rev-list HEAD --max-count=1 --abbrev-commit)
OS?=      linux
ARCH?=    amd64

LDFLAGS=-ldflags "-X main.Service=${PROGRAM} -X main.Version=${VERSION} -X main.Commit=${COMMIT}"

all: dep build

dep:
	dep ensure

build: dep
	CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -x -v ${LDFLAGS} -o ${PROGRAM}

docker: build
	docker build --no-cache -t ${PROGRAM}:${VERSION} .

clean:
	go clean -x -v
	rm -rf vendor/
