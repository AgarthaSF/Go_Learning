package main

import (
	"fmt"
	"sync"
	"time"
)

func watchDog(stopCh chan bool, name string) {
	//开启for select循环，一直后台监控
	for {
		select {
		case <-stopCh:
			fmt.Println("received instruction")
			return
		default:
			fmt.Println(name, "monitoring...")
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	stopCh := make(chan bool)

	go func() {
		defer wg.Done()
		watchDog(stopCh, "watchDog")
	}()

	time.Sleep(5 * time.Second)
	stopCh <- true
	wg.Wait()

}
