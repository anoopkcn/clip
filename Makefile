GO  		   ?= go
GOOS    	   ?= $(word 1, $(subst /, " ", $(word 4, $(shell go version))))

MAKEFILE       := $(realpath $(lastword $(MAKEFILE_LIST)))
ROOT_DIR       := $(shell dirname $(MAKEFILE))
SOURCES        := $(wildcard *.go src/*.go src/*/*.go) $(MAKEFILE)

REVISION       := $(shell git log -n 1 --pretty=format:%h -- $(SOURCES))
BUILD_FLAGS    := -a -ldflags "-X main.revision=$(REVISION)" #-w -extldflags=$(LDFLAGS)"

BINARY32       := clip-$(GOOS)_386
BINARY64       := clip-$(GOOS)_amd64
BINARYARM5     := clip-$(GOOS)_arm5
BINARYARM6     := clip-$(GOOS)_arm6
BINARYARM7     := clip-$(GOOS)_arm7
BINARYARM8     := clip-$(GOOS)_arm8
BINARYPPC64LE  := clip-$(GOOS)_ppc64le
VERSION        := $(shell awk -F= '/version =/ {print $$2}' src/constants.go | tr -d "\" ")
RELEASE32      := clip-$(VERSION)-$(GOOS)_386
RELEASE64      := clip-$(VERSION)-$(GOOS)_amd64
RELEASEARM5    := clip-$(VERSION)-$(GOOS)_arm5
RELEASEARM6    := clip-$(VERSION)-$(GOOS)_arm6
RELEASEARM7    := clip-$(VERSION)-$(GOOS)_arm7
RELEASEARM8    := clip-$(VERSION)-$(GOOS)_arm8
RELEASEPPC64LE := clip-$(VERSION)-$(GOOS)_ppc64le


# https://en.wikipedia.org/wiki/Uname
UNAME_M := $(shell uname -m)
ifeq ($(UNAME_M),x86_64)
	BINARY := $(BINARY64)
else ifeq ($(UNAME_M),amd64)
	BINARY := $(BINARY64)
else ifeq ($(UNAME_M),i686)
	BINARY := $(BINARY32)
else ifeq ($(UNAME_M),i386)
	BINARY := $(BINARY32)
else ifeq ($(UNAME_M),armv5l)
	BINARY := $(BINARYARM5)
else ifeq ($(UNAME_M),armv6l)
	BINARY := $(BINARYARM6)
else ifeq ($(UNAME_M),armv7l)
	BINARY := $(BINARYARM7)
else ifeq ($(UNAME_M),armv8l)
	BINARY := $(BINARYARM8)
else ifeq ($(UNAME_M),ppc64le)
	BINARY := $(BINARYPPC64LE)
else
$(error "Build on $(UNAME_M) is not supported, yet.")
endif

all: target/$(BINARY)

target:
	mkdir -p $@

install: bin/clip

clean:
	$(RM) -r target

target/$(BINARY32): $(SOURCES)
	GOARCH=386 $(GO) build $(BUILD_FLAGS) -o $@

target/$(BINARY64): $(SOURCES)
	GOARCH=amd64 $(GO) build $(BUILD_FLAGS) -o $@

# https://github.com/golang/go/wiki/GoArm
target/$(BINARYARM5): $(SOURCES)
	GOARCH=arm GOARM=5 $(GO) build $(BUILD_FLAGS) -o $@

target/$(BINARYARM6): $(SOURCES)
	GOARCH=arm GOARM=6 $(GO) build $(BUILD_FLAGS) -o $@

target/$(BINARYARM7): $(SOURCES)
	GOARCH=arm GOARM=7 $(GO) build $(BUILD_FLAGS) -o $@

target/$(BINARYARM8): $(SOURCES)
	GOARCH=arm64 $(GO) build $(BUILD_FLAGS) -o $@

target/$(BINARYPPC64LE): $(SOURCES)
	GOARCH=ppc64le $(GO) build $(BUILD_FLAGS) -o $@

bin/clip: target/$(BINARY) | bin
	cp -f target/$(BINARY) bin/clip
