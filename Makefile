all: build

.PHONY: build
build:
	go build cmd/jackup/jackup.go

# Set GITHUB_TOKEN personal access token and create release git tag
.PHONY: release
release:
	go get -u github.com/goreleaser/goreleaser
	goreleaser --rm-dist
