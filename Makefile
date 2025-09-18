GO_LDFLAGS  := -ldflags "-X github.com/riceriley59/synkra/internal/version.GIT_SHA=$(GIT_SHA) \
						 -X github.com/riceriley59/synkra/internal/version.VERSION=$(VERSION)"
GOFLAGS 	:=

.PHONY: all
all: build

.PHONY: build
build:
	go build $(GOFLAGS) $(GO_LDFLAGS) -o bin/synkra cmd/synkra/main.go
