.PHONY: clean install

clean:
	go clean -i github.com/kengibson1111/go-stdlib/...

install:
	go install ./...
