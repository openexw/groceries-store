package interview

import (
	"math/rand"
	"sync"
)

func production_and_consume() {
	out := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			out <- rand.Intn(5)
		}
		close(out)
	}()

	//go func() {
	//	defer wg.Done()
	//	for v := range out {
	//		println(v)
	//	}
	//}()
	for {
		select {
		case v := <-out:
			println(v)
		default:
		}
		//time.Sleep(time.Second * 1)
	}
	wg.Wait()


}
