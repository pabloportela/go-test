package main

import (
	"fmt"
	"time"
)

const max_calls_per_interval = 100
const call_interval_nonosecs = 1000000000

func callApi(ch chan int, i int) {
	_ = <-ch
	doCallApi(i)
}

func doCallApi(i int) {
	fmt.Println(i)
}

func dispatchCalls(ch chan int) {
	i := 0
	v := make([]int64, max_calls_per_interval)

	for {
		current := time.Now().UnixNano()
		if current-v[i%max_calls_per_interval] >= call_interval_nonosecs {
			ch <- 1
			v[i%max_calls_per_interval] = current
			i++
		}
	}
}

func main() {
	ch := make(chan int)

	go dispatchCalls(ch)

	// queue up calls
	for i := 0; i < 102; i++ {
		go callApi(ch, i)
	}

	// live for ten seconds
	time.Sleep(time.Second * 10)
}
