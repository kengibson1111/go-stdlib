all: clean install

clean:
	go clean -i github.com/kengibson1111/go-packages/...

install:
	go install ./...

