package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing counter 3 times", func(t *testing.T) {
		counter := Counter{}
		counter.Increment()
		counter.Increment()
		counter.Increment()

		assertCounter(t, &counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCounted := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCounted)

		for i := 0; i < wantedCounted; i++ {
			go func() {
				counter.Increment()
				wg.Done()
			}()
		}

		wg.Wait()
		assertCounter(t, &counter, wantedCounted)
	})

}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
