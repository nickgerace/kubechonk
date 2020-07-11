MAKEPATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
NAME:=kubectl-chonk
DEST:=/usr/local/bin

build: fmt-and-vet test
	cd $(MAKEPATH); go build cmd/$(NAME).go

fmt-and-vet:
	cd $(MAKEPATH); go fmt ./...
	cd $(MAKEPATH); go vet ./...

test:
	cd $(MAKEPATH); go test ./...

tidy:
	cd $(MAKEPATH); go mod tidy

install: uninstall
	mv $(NAME) $(DEST)/$(NAME)

uninstall:
	-rm $(DEST)/$(NAME)
