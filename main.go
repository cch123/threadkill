package threadkill

import (
	"runtime"
	"sync"
)

// KillOne kills a thread
func KillOne() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		runtime.LockOSThread()
		return
	}()

	wg.Wait()
}
