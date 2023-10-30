package utils

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Operation func(context.Context) error

func GracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]Operation) <-chan struct{} {
	// used to block goroutine until all goroutines finished
	wait := make(chan struct{})

	go func() {
		s := make(chan os.Signal, 1)

		defer close(wait)

		// add syscall that want to be notified with
		signal.Notify(s, syscall.SIGTERM, syscall.SIGINT)
		<-s

		defer close(s)

		log.Println("shutting down application")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit\n", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key

			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %v", innerKey)

				if err := innerOp(ctx); err != nil {
					log.Printf("cleaning up %s failed: %s", innerKey, err.Error())
					return
				}
			}()
		}

		wg.Wait()
	}()

	return wait
}
