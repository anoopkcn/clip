GO  		   ?= go
GOOS    	   ?= $(word 1, $(subst /, " ", $(word 4, $(shell go version))))

MAKEFILE       := $(realpath $(lastword $(MAKEFILE_LIST)))
ROOT_DIR       := $(shell dirname $(MAKEFILE))
SOURCES        := $(wildcard *.go src/*.go src/*/*.go) $(MAKEFILE)

# VERSION        := $(shell awk -F= '/version =/ {print $$2}' src/constants.go | tr -d "\" ")
REVISION       := $(shell git log -n 1 --pretty=format:%h -- $(SOURCES))
BUILD_FLAGS    := -a -ldflags "-X main.revision=$(REVISION)" #-w -extldflags=$(LDFLAGS)"

all: bin/clip

bin:
	mkdir -p $@

clean:
	$(RM) -r bin

bin/clip: $(SOURCES)
	$(GO) build $(BUILD_FLAGS) -o $@
