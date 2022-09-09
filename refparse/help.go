package refparse

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/iancoleman/strcase"
	"github.com/mitchellh/go-wordwrap"
	"github.com/mplewis/figyr/types"
	"golang.org/x/exp/slices"
)

type ArgHelp struct {
	FieldDef
	Description string
}

func (a ArgHelp) String() string {
	return fmt.Sprintf("%s: %s", a.Name, a.Description)
}

func printHelpAndExitIfRequested(po types.ParserOptions, defs []FieldDef) {
	if !slices.Contains(os.Args, "--help") {
		return
	}

	fmt.Fprintln(os.Stderr, wordwrap.WrapString(po.Description, po.LineLength))
	fmt.Fprintln(os.Stderr)

	fmt.Fprintln(os.Stderr, "Options:")
	tw := tabwriter.NewWriter(os.Stderr, 8, 0, 3, ' ', 0)
	for _, def := range defs {
		fmt.Fprintf(tw, "    --%s\t%s\t%s\n", strcase.ToKebab(def.Name), def.Constraint(), def.Description)
	}
	tw.Flush()
	os.Exit(0)
}
