package main

func main() {
	ch := make(chan int, 3)
	//<-ch
	for i := 0; i < 3; i++ {
		go func(ch chan int) {
			for {
				println(<-ch)
			}
		}(ch)
	}
	for i := 0; i < 19; i++ {
		ch <- i
	}
	//time.Sleep(1)
}

// go tool compile -S chan-s.go
// GOSSAFUNC=main go build -gcflags="-N -l" chan-s.go
