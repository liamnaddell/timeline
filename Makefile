DEP=~/go/bin/dep
GO=go
FINAL=/usr/local/bin
VERSION=0.0.1-alpha
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

all:
	$(DEP) ensure
	$(GO) build $(LDFLAGS)

install:
	mv timeline $(FINAL)
clean:
	rm timeline
	rm -r vendor/

uninstall:
	rm $(FINAL)/timeline
