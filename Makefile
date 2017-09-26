.PHONY: test install

test:
	go test -cover
	go vet

install:
	go install
