package main

import (
	"sync"
)

type fun func() bool

// RunShortCircuit 多个函数并发执行，短路返回
// 比 sync.WaitGroup 延迟要低
func RunShortCircuit(fs ...fun) bool {
	quit := make(chan bool)
	done := make(chan bool)

	go func() {
		g := sync.WaitGroup{}

		rs := make(chan bool, len(fs))

		for i := range fs {
			g.Add(1)
			f := fs[i]
			go func() {
				defer g.Done()
				v := f()
				if !v {
					quit <- v
				}
				rs <- v
			}()
		}

		g.Wait()

		close(rs)

		for v := range rs {
			if !v {
				done <- false
				return
			}
		}
		done <- true
	}()

	select {
	case <-quit:
		return false
	case v := <-done:
		return v
	}
}
