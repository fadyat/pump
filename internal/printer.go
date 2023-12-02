package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type Printer interface {
	Print(values [][]string)
}

type tablePrinter struct {
	headers []string
}

func NewTablePrinter(
	headers ...string,
) Printer {
	return &tablePrinter{
		headers: headers,
	}
}

func (p *tablePrinter) Print(values [][]string) {
	var (
		pattern   = p.getPattern(values)
		header    = fmt.Sprintf(pattern, toAny(p.headers)...)
		tableRows = make([]any, len(values)+2)
	)

	tableRows[0] = header
	tableRows[1] = strings.Repeat("-", len(header))
	for i, value := range values {
		tableRows[i+2] = fmt.Sprintf(pattern, toAny(value)...)
	}

	for _, row := range tableRows {
		fmt.Println(row)
	}
}

func (p *tablePrinter) getPattern(values [][]string) string {
	var patterns = make([]string, len(p.headers))
	for idx, header := range p.headers {
		var columnArgs = make([]string, len(values))
		for i, value := range values {
			columnArgs[i] = value[idx]
		}

		patterns[idx] = p.takeOptimalPattern(header, columnArgs)
	}

	return strings.Join(patterns, " | ")
}

func (p *tablePrinter) takeOptimalPattern(
	header string,
	values []string,
) string {
	var largestNameLength = len(header)
	for _, value := range values {
		largestNameLength = max(largestNameLength, len(value))
	}

	return "%-" + strconv.Itoa(largestNameLength) + "s"
}

func toAny[T any](val []T) []any {
	var result = make([]any, len(val))
	for idx, value := range val {
		result[idx] = value
	}

	return result
}
