//go:build tools
// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser"
	_ "github.com/vektra/mockery/v2"
)

//go:generate go build -v -o .bin/ github.com/golangci/golangci-lint/cmd/golangci-lint

//go:generate go build -v -o .bin/ github.com/goreleaser/goreleaser

//go:generate go build -v -o .bin/ github.com/vektra/mockery/v2
