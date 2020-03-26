package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func Test_orChannel(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-orChannel(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Hour),
		sig(1*time.Second),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
