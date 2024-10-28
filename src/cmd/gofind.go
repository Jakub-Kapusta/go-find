// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package cmd

import (
	"context"
	"os"
	"sync"

	"github.com/Jakub-Kapusta/go-find/apps/find"
	"github.com/Jakub-Kapusta/go-find/internal/signals"
	"github.com/spf13/cobra"
)

var (
	// WaitGroup for global goroutines.
	wg sync.WaitGroup
	// Global ctx: should only be cancelled by the signal handler.
	ctx, cf = context.WithCancel(context.Background())
)

var rootCmd = &cobra.Command{
	Version: "v0.0.1",
	Use:     "gofind TODO",
	Short:   "Find files and directories.",
	Long: `A partial GNU findutils replacement implemented ing GO.

This application is under construction.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := validateFlags(); err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			os.Exit(1)
		}
		signals.LaunchSignalHandler(ctx, cf, &wg)
	},
	// args[0] is the first actual argument, and not the name of the program.
	// Only arguments not caught by our flag definitions will be present.
	Run: func(cmd *cobra.Command, args []string) {

		find.Find(
			ctx,
			args,
			&find.FinderOptions{
				RootDir:      rootDir,
				IsSearchPath: isSearchPath,
				SearchPath:   searchPath,
			},
			unsafePrint,
			print0)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		cf()
		wg.Wait()
	},
}

func init() {
	setFlags()
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
