// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package signals

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func LaunchSignalHandler(ctx context.Context, cf context.CancelFunc, wg *sync.WaitGroup) {
	// Make the buffer large enough to hold all signal types we want to catch.
	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func(ctx context.Context, cf context.CancelFunc, wg *sync.WaitGroup, c <-chan os.Signal) {
		defer cf()
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case <-c:
				// Add Leading newline because preceding output might be broken in case of cancel.
				os.Stderr.WriteString("\nsignal received\n")
				return
			}
		}
	}(ctx, cf, wg, c)
}
