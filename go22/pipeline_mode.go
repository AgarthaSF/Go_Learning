package main

import (
	"fmt"
	"sync"
)

func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- fmt.Sprint("component", i)
		}
	}()
	return out
}

func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "build(" + c + ")"
		}
	}()
	return out
}

func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "pack(" + c + ")"
		}
	}()
	return out
}

func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	// 将channel数据发送至out中
	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}

	wg.Add(len(ins))

	// 扇入
	for _, cs := range ins {
		go p(cs)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// result := make(chan string)
	// go func() {
	// 	// 模拟网络访问
	// 	time.Sleep(8 * time.Second)
	// 	result <- "Server result"
	// }()

	// select {
	// case v := <-result:
	// 	fmt.Println(v)

	// case <-time.After(5 * time.Second):
	// 	fmt.Println("connection timeout")
	// }

	// coms := buy(10)
	// phones := build(coms)
	// packs := pack(phones)

	// for p := range packs {
	// 	fmt.Println(p)
	// }

	coms := buy(100)

	phones1 := build(coms)
	phones2 := build(coms)
	phones3 := build(coms)

	phones := merge(phones1, phones2, phones3)
	packs := pack(phones)

	for p := range packs {
		fmt.Println(p)
	}

}
