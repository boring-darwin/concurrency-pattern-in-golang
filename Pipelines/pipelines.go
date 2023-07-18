package pipelines

import "time"

func Gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func Sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
			time.Sleep(1 * time.Second)
		}
		close(out)
	}()
	return out
}
