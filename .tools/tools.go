//go:build tools
// +build tools

package tools

// Managing tools with Go Modules:
//

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser"
)

//go:generate go build -v -o .bin/ github.com/golangci/golangci-lint/cmd/golangci-lint

//go:generate go build -v -o .bin/ github.com/goreleaser/goreleaser
