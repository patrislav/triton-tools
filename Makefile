REPOSITORY := github.com/patrislav/triton-tools
GOBUILD := go build

.PHONY: all
all: bin/trasm bin/detrasm

bin/trasm:
	$(GOBUILD) -o bin/trasm $(REPOSITORY)/cmd/trasm

bin/detrasm:
	$(GOBUILD) -o bin/detrasm $(REPOSITORY)/cmd/detrasm
