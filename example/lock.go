package main

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/jasonlvhit/gocron"
)

// Run a Redis instance with Docker: docker run --rm -tid -p 6379:6379 redis:alpine

func lockedTask(name string) { _ = "STUB: not implemented"; return }

// locker implementation with Redis
type locker struct {
	cache *redis.Client
}

func (s *locker) Lock(key string) (success bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *locker) Unlock(key string) error { _ = "STUB: not implemented"; return nil }

// Run the example in different terminals,
// passing a different name parameter to each
func main() {
	// Get a locker
	l := &locker{
		redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		}),
	}

	// Make locker available for the cron jobs
	gocron.SetLocker(l)

	arg := "Some Name"
	args := os.Args[1:]
	if len(args) > 0 {
		arg = args[0]
	}

	gocron.Every(1).Second().Lock().Do(lockedTask, arg)
	<-gocron.Start()
}
