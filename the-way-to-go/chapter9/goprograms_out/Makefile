export GOPATH=$(shell pwd)
export GOBIN=$(GOPATH)/bin

install:
	go get
	go test
	go install
test:
	go test
clean:
	rm -rf $(GOBIN)/*
