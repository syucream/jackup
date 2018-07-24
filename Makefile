all: build

.PHONY: build
build:
	go build cmd/jackup/jackup.go

.PHONY: check
check: build
	./jackup -f examples/create_database.sql > /dev/null
	./jackup -f examples/create_table.sql > /dev/null
	./jackup -f examples/create_index.sql > /dev/null
	./jackup -f examples/composition.sql > /dev/null
	cat examples/create_database.sql | ./jackup > /dev/null
	cat examples/create_table.sql | ./jackup > /dev/null
	cat examples/create_index.sql | ./jackup > /dev/null
	cat examples/composition.sql | ./jackup > /dev/null

# Set GITHUB_TOKEN personal access token and create release git tag
.PHONY: release
release:
	go get -u github.com/goreleaser/goreleaser
	goreleaser --rm-dist
