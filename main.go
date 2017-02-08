package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/peiqi-caicloud/exec/priorchan"
)

var wg sync.WaitGroup

func produce(elements *priorchan.Elements) {
	for i := 0; i < 10; i++ {
		msg := &priorchan.Element{Priority: i % 3, Content: "dodo"}
		//	fmt.Printf("producer: priority %v\n", msg)
		elements.Push(msg)
	}
}

func consume(id int, elements *priorchan.Elements) {

	time.Sleep(50 * time.Millisecond)
	msg, flag := elements.Pop()
	if flag == false {
	time.Sleep(50 * time.Millisecond)
		fmt.Printf("No data, waiting...\n")
	}
	fmt.Printf("Consumer%d: consume msg %v\n", id, msg)

	wg.Done()
}

func main() {
	elements := priorchan.NewElements()
	go produce(elements)

	wg.Add(10)
	for id := 0; id < 10; id++ {
		go consume(id, elements)
	}
	wg.Wait()
}
