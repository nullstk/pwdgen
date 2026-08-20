package core

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Config holds the options for a run.
type Config struct {
	Verbose bool
	DryRun bool
	Output string
}

// Result summarizes one run.
type Result struct {
	Files int `json:"files"`
	Elapsed time.Duration `json:"elapsed_ms"`
}

// Run executes the pwdgen workload.
func Run(cfg Config) error {
	start := time.Now()

	if !cfg.DryRun {
 if err := os.MkdirAll(cfg.Output, 0o755); err != nil {
 return fmt.Errorf("create output: %w", err)
 }
	}

	for i := 0; i < 4; i++ {
 name := filepath.Join(cfg.Output, fmt.Sprintf("result-%d.txt", i))
 content := fmt.Sprintf("PwdGen\nitem %d\n", i)
 if cfg.DryRun {
 if cfg.Verbose {
 fmt.Printf("would write %s\n", name)
 }
 continue
 }
 if err := os.WriteFile(name, []byte(content), 0o644); err != nil {
 return fmt.Errorf("write %s: %w", name, err)
 }
	}

	if cfg.Verbose {
 fmt.Printf("wrote files in %s\n", time.Since(start))
	}
	return nil
}