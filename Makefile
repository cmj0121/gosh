.PHONY: all clean help

SRC=$(wildcard *.go) $(wildcard */*.go)
BIN=$(subst .go,,$(wildcard cmd/*.go))

all: $(BIN)	# build all

clean:	# clean-up the temporary file and environment
	rm -f $(BIN)

help:	# show this message
	@printf "Usage: make [OPTION]\n"
	@printf "\n"
	@perl -nle 'print $$& if m{^[\w-]+:.*?#.*$$}' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?#"} {printf "    %-18s %s\n", $$1, $$2}'

linter: $(SRC)
	gofmt -s -w $^
	go test -cover -failfast -timeout 2s ./...

$(BIN): linter

%: %.go
	go build -ldflags="-s -w" -o $@ $<
