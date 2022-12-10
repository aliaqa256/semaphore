package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"semaphore/queue"
	"syscall"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	q := queue.Queue{}
	ctx := context.Background()
	count := 10
	sem := semaphore.NewWeighted(int64(count))

	for i := 0; i < 2; i++ {

		go func() {
			a := 0
			for {
				sem.Acquire(ctx, 1)
				q.Enqueue(a)
				a++
				time.Sleep(time.Duration(rand.Intn(500)+100) * time.Millisecond)

				// if signal was received, exit
				select {
				case <-sigs:
					os.Exit(0)
				default:
				}
			}

		}()

	}

	for i := 0; i < count; i++ {
		go func(i int) {
			for {
				out := q.Dequeue()
				if out == -1 {
					continue
				}
				time.Sleep(time.Duration(rand.Intn(500)+100) * time.Millisecond)
				sem.Release(1)
				fmt.Println(out, i)

				// if signal was received, exit
				select {
				case <-sigs:
					os.Exit(0)
				default:
				}
			}
		}(i)

	}

	<-sigs
}
