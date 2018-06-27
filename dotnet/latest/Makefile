PACKAGE   = joshgav/az-profiles/dotnet/latest
BASE      = $(GOPATH)/src/github.com/$(PACKAGE)
DOTNET = dotnet

run: restore
	dotnet run $(BASE)

build: restore
	mkdir $(BASE)/out
	dotnet build -o $(BASE)/out/profiles-tester $(BASE)

restore:
	dotnet restore $(BASE)

.PHONY: run build restore
