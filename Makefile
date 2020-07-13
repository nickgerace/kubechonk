MAKEPATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
NAME:=kubectl-chonk
TAG:=v0.1.1

build: fmt-and-vet test
	cd $(MAKEPATH); go build cmd/$(NAME).go

fmt-and-vet:
	cd $(MAKEPATH); go fmt ./...
	cd $(MAKEPATH); go vet ./...

test:
	cd $(MAKEPATH); go test ./...

tidy:
	cd $(MAKEPATH); go mod tidy

test-release:
	cd $(MAKEPATH); goreleaser --snapshot --skip-publish --rm-dist

release: test-release
	cd $(MAKEPATH); git tag -a $(TAG) -m "Release $(TAG)"
	cd $(MAKEPATH); git push origin $(TAG)

install:
	kubectl krew install --manifest=$(MAKEPATH)/.krew.yaml -v=4

uninstall:
	-kubectl krew uninstall chonk

upgrade: uninstall install
