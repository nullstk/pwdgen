package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"pwdgen/internal/core"
)

func main() {
	verbose := flag.Bool("v", false, "verbose output")
	dry := flag.Bool("dry-run", false, "print without writing")
	output := flag.String("o", "./out", "output directory")
	flag.Parse()

	cfg := core.Config{
 Verbose: *verbose,
 DryRun: *dry,
 Output: *output,
	}

	if err := core.Run(cfg); err != nil {
 log.Fatalf("%v", err)
	}

	if cfg.Verbose {
 fmt.Fprintln(os.Stderr, "PwdGen: done")
	}
}