package internal

import (
	"fmt"
	"github.com/fadyat/pump/pkg"
	"os"
	"os/exec"
	"strings"
)

type Editor interface {
	Edit(values []string) error
}

type editor struct {
	editorName string
	units      []string
}

func NewEditor(units ...string) Editor {
	return &editor{
		units:      units,
		editorName: "nvim",
	}
}

func (e *editor) Edit(values []string) error {
	var content = e.getRawContent(values)
	if err := pkg.WriteFile(pkg.GetTempFile(), content); err != nil {
		return err
	}

	return e.open(pkg.GetTempFile())
}

func (e *editor) getRawContent(values []string) []byte {
	var content = make([]string, len(e.units))
	for idx, unit := range e.units {
		content[idx] = fmt.Sprintf("### %s\n%s\n", unit, values[idx])
	}

	return []byte(strings.Join(content, "\n"))
}

func (e *editor) open(path string) error {
	cmd := exec.Command(e.editorName, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
