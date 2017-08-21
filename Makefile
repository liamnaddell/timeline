DEP=~/go/bin/dep
GO=go
FINAL=/usr/local/bin

all:
	$(DEP) ensure
	$(GO) build

install:
	mv timeline $(FINAL)
clean:
	rm timeline

uninstall:
	rm $(FINAL)/timeline
