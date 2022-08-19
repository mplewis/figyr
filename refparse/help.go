package refparse

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
)

type ArgHelp struct {
	FieldDef
	Description string
}

func (a ArgHelp) String() string {
	return fmt.Sprintf("%s: %s", a.Name, a.Description)
}

func printHelpAndExitIfRequested(defs []FieldDef) {
	if !slices.Contains(os.Args, "--help") {
		return
	}

	fmt.Println("Options:")
	tw := tabwriter.NewWriter(os.Stdout, 8, 0, 3, ' ', 0)
	for _, def := range defs {
		fmt.Fprintf(tw, "    --%s\t%s\t%s\n", strcase.ToKebab(def.Name), def.Constraint(), def.Description)
	}
	tw.Flush()
	os.Exit(0)
}
