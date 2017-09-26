.PHONY: test build release

test:
	go test -cover
	go vet

build:
	go install -ldflags "-X main.version=`date +"%Y%m%d"`+`git rev-parse --short HEAD`"

release:
	mkdir -p binaries/linux/amd64
	mkdir -p binaries/darwin/amd64
	mkdir -p binaries/windows/amd64
	GOARCH=amd64 GOOS=linux go build -ldflags "-X main.version=`date +"%Y%m%d"`+`git rev-parse --short HEAD`" -o binaries/linux/amd64/makedesc
	GOARCH=amd64 GOOS=darwin go build -ldflags "-X main.version=`date +"%Y%m%d"`+`git rev-parse --short HEAD`" -o binaries/darwin/amd64/makedesc
	GOARCH=amd64 GOOS=windows go build -ldflags "-X main.version=`date +"%Y%m%d"`+`git rev-parse --short HEAD`" -o binaries/windows/amd64/makedesc.exe
	git branch -D binaries 2>/dev/null | true
	git branch -D binaries-draft 2>/dev/null | true
	git checkout -b binaries-draft
	git add -f binaries
	git commit -m "Release to branch binaries"
	git subtree split --prefix binaries -b binaries
	git push --force origin binaries:binaries
	git checkout -
