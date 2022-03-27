package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)
	go func() {
		go fmt.Println("AgarthaSF")
		ch <- "goroutine complete"
	}()

	fmt.Println("I'm main goroutine")
	v := <-ch
	fmt.Println("The value received from chan is", v)

	cacheCh := make(chan int, 5)
	cacheCh <- 2
	cacheCh <- 3

	fmt.Println("cacheCh容量为:", cap(cacheCh), ",元素个数为：", len(cacheCh))

	firstCh := make(chan string)
	secondCh := make(chan string)
	thirdCh := make(chan string)

	// 同时开启三个goroutine下载

	go func() {
		firstCh <- downloadFile("firstCh")
	}()

	go func() {
		secondCh <- downloadFile("secondCh")
	}()

	go func() {
		thirdCh <- downloadFile("thirdCh")
	}()

	//开始select多路复用，哪个channel能获取到值，
	//就说明哪个最先下载好，就用哪个。
	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-thirdCh:
		fmt.Println(filePath)
	}

}

func downloadFile(chanName string) string {
	// 模拟下载文件
	time.Sleep(time.Second)
	return chanName + ":filePath"
}
