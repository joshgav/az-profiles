PACKAGE   = github.com/joshgav/az-profiles/go/latest
BASE      = $(GOPATH)/src/$(PACKAGE)
GO = go

run: dep
	go run $(BASE)/main.go

build: dep
	mkdir $(BASE)/out
	go build -o $(BASE)/out/auth-tester .

dep: $(BASE)
	go get -u github.com/golang/dep
	cd $(BASE) && dep ensure

.PHONY: run build dep
