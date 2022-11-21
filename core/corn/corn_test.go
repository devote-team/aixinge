package corn

import (
	"fmt"
	"testing"
	"time"
)
import "github.com/go-co-op/gocron"

func TestCorn(t *testing.T) {
	// https://github.com/go-co-op/gocron
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(func() {
		fmt.Println("----cron-----")
	})
	// you can start running the scheduler in two different ways:
	// starts the scheduler asynchronously
	s.StartAsync()
	// starts the scheduler and blocks current execution path
	s.StartBlocking()
}
