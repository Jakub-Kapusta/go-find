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

	rootCmd = &cobra.Command{
		Version: version,
		Use:     "gofind TODO",
		Short:   "Find files and directories.",
		Long: `A partial GNU findutils replacement implemented ing GO.

	This application is under construction.`,
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			signals.LaunchSignalHandler(ctx, cf, &wg)
		},

		Run: func(cmd *cobra.Command, _ []string) {
			find.Find(
				ctx,
				&find.FinderOptions{
					RootDir:      rootDir,
					IsSearchPath: isSearchPath,
					SearchPath:   searchPath,
				},
				unsafePrint,
				print0)
		},
		PersistentPostRun: func(cmd *cobra.Command, _ []string) {
			cf()
			wg.Wait()
		},
	}
)

func CreateAndExecute() {
	setFlags()

	if err := rootCmd.Execute(); err != nil {
		os.Stderr.WriteString("rootCmd.Execute() failed")
		os.Exit(1)
	}
}
