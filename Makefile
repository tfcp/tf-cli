.PHONY: install clean help

all:
	@cd ./ && go build && cd - &> /dev/null

.PHONY: install
install: all
	@cp ./tf-cli /usr/bin

