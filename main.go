package main

import (
	"fmt"
	_ "net/http/pprof"
	"time"

	filewalk "github.com/boring-darwin/Filewalk"
)

// func main() {
// 	// Set up the pipeline.
// 	narr := make([]int, 0)
// 	for i := 1; i < 1000; i++ {
// 		narr = append(narr, i)
// 	}
// 	fmt.Printf("Length of array %d\n", len(narr))
// 	c := pipelines.Gen(narr...)
// 	out := pipelines.Sq(c)

// 	// Consume the output.
// 	// fmt.Println(<-out) // 4
// 	// fmt.Println(<-out) // 9

// 	for val := range out {
// 		fmt.Println(val)
// 	}

// }

// Since sq has the same type for its inbound and outbound channels,
// we can compose it any number of times. We can also rewrite main as a range loop, like the other stages:
// func main() {
// 	// Set up the pipeline and consume the output.
// 	for n := range pipelines.Sq(pipelines.Sq(pipelines.Gen(2, 3))) {
// 		fmt.Println(n) // 16 then 81
// 	}
// }

// func main() {
// 	in := pipelines.Gen(2, 3)

// 	// Distribute the sq work across two goroutines that both read from in.
// 	c1 := pipelines.Sq(in)
// 	c2 := pipelines.Sq(in)

// 	// Consume the merged output from c1 and c2.
// 	for n := range faninfanout.Merge(c1, c2) {
// 		fmt.Println(n) // 4 then 9, or 9 then 4
// 	}
// }

// func main() {
// 	// Set up the pipeline.
// 	c := pipelines.Gen(2, 3)
// 	out := pipelines.Sq(c)

// 	// Consume the output.
// 	fmt.Println(<-out) // 4
// 	// return
// 	// fmt.Println(<-out) // 9
// 	log.Println(http.ListenAndServe("localhost:6060", nil))
// }

// func main() {
// 	in := pipelines.Gen(2, 3)

// 	// Consume the first value from output.
// 	done := make(chan struct{})
// 	defer close(done)

// 	// Distribute the sq work across two goroutines that both read from in.
// 	c1 := pipelines.Sq(done, in)
// 	c2 := pipelines.Sq(done, in)

// 	// Consume the merged output from c1 and c2.
// 	out := faninfanout.Merge(done, c1, c2)
// 	fmt.Println(<-out) // 4 or 9
// 	fmt.Println(<-out)
// 	// Tell the remaining senders we're leaving.

// 	log.Println(http.ListenAndServe("localhost:6060", nil))
// }

//------------- File Walking Example --------------------

func main() {

	now := time.Now()
	data, err := filewalk.MD5AllWithoutParalleism("/Users/nitin/workspace/")
	if err != nil {
		fmt.Printf("Something went wrong in MD5AllWithoutParalleism %v\n", err.Error())
	}
	totalTime := time.Since(now)
	fmt.Printf("Result is having length: %d Total Time for MD5AllWithoutParalleism %v\n", len(data), totalTime)

	now = time.Now()
	data, err = filewalk.MD5All("/Users/nitin/workspace/")
	if err != nil {
		fmt.Printf("Something went in wrong MD5All %v\n", err.Error())
	}
	totalTime = time.Since(now)
	fmt.Printf("Result is having length: %d Total Time for MD5All %v\n", len(data), totalTime)
}
