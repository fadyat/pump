BIN_TOOLS = ./.tools/.bin


tools:
	@cd ./.tools && go mod tidy && go mod verify && go generate --tags=tools -x

lint:
	@$(BIN_TOOLS)/golangci-lint run --config .golangci.yml ./...

pre-release:
	@$(BIN_TOOLS)/goreleaser --snapshot --skip=publish --clean

.PHONY: tools lint pre-release
