package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/unravel1020/Easy_Setup/internal/catalog"
)

func main() {
	catalogPath := flag.String("catalog", "src/catalog/catalog.json", "path to catalog JSON")
	platform := flag.String("platform", "windows", "target platform")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "usage: easysetup [flags] catalog|plan [stack-or-recipe-id...]")
		os.Exit(2)
	}

	cat, err := catalog.Load(*catalogPath)
	if err != nil {
		fail(err)
	}

	switch flag.Arg(0) {
	case "catalog":
		writeJSON(cat)
	case "plan":
		plan, err := cat.Plan(*platform, flag.Args()[1:])
		if err != nil {
			fail(err)
		}
		writeJSON(plan)
	default:
		fail(fmt.Errorf("unknown command %q", flag.Arg(0)))
	}
}

func writeJSON(value any) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(value); err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
