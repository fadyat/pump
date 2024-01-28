package internal

import (
	"fmt"
	"github.com/fadyat/pump/pkg"
	"os"
	"os/exec"
	"strings"
)

type Editor interface {
	Edit(values []string) ([]string, error)
}

type editor struct {
	editorName string
	units      []string
}

func NewEditor(units ...string) Editor {
	var editorName = os.Getenv("EDITOR")
	if editorName == "" {
		editorName = "vim"
	}

	return &editor{
		units:      units,
		editorName: editorName,
	}
}

func (e *editor) Edit(values []string) ([]string, error) {
	var content = e.getRawContent(values)
	if err := pkg.WriteFile(pkg.GetTempFile(), content); err != nil {
		return nil, fmt.Errorf("failed to write file: %w", err)
	}

	if err := e.open(pkg.GetTempFile()); err != nil {
		return nil, fmt.Errorf("failed to open editor: %w", err)
	}

	rawContent, err := os.ReadFile(pkg.GetTempFile())
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return e.parseRawContent(rawContent), nil
}

func (e *editor) getRawContent(values []string) []byte {
	var content = make([]string, len(e.units))
	for idx, unit := range e.units {
		content[idx] = fmt.Sprintf("### %s\n%s\n", unit, values[idx])
	}

	return []byte(strings.Join(content, "\n"))
}

func (e *editor) parseRawContent(rawContent []byte) []string {
	lines := strings.Split(string(rawContent), "\n")
	unitToValues := make(map[string]string)

	var (
		currentUnit  string
		currentValue strings.Builder
	)

	for _, line := range lines {
		if strings.HasPrefix(line, "### ") {
			if currentUnit != "" {
				unitToValues[currentUnit] = currentValue.String()
			}

			currentUnit = strings.TrimPrefix(line, "### ")
			currentValue.Reset()
			continue
		}

		currentValue.WriteString(line)
		currentValue.WriteString("\n")
	}

	unitToValues[currentUnit] = currentValue.String()

	var values = make([]string, len(e.units))
	for idx, unit := range e.units {
		values[idx] = strings.Trim(unitToValues[unit], "\n")
	}

	return values
}

func (e *editor) open(path string) error {
	cmd := exec.Command(e.editorName, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
