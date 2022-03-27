package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum   int
	mutex sync.RWMutex
)

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}

func readSum() int {
	mutex.RLock()
	defer mutex.RUnlock()
	b := sum
	return b
}

func run() {

	var wg sync.WaitGroup
	// 监控110个携程所以设置为110
	wg.Add(110)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			add(10)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			go fmt.Println(readSum())
		}()
	}

	wg.Wait()
}

func race() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号就位")
			cond.L.Lock()
			cond.Wait() // 等待信号
			fmt.Println(num, "号开始跑")
			cond.L.Unlock()
		}(i)
	}

	//等待所有goroutine都进入wait状
	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("裁判就位，准备信号")
		fmt.Println("比赛开始")
		cond.Broadcast()
	}()

	wg.Wait()

}

func main() {
	//run()
	//race()

	var m sync.Map

	// 写入
	m.Store("test", 18)
	m.Store("ano", 20)

	// 读取，第一个参数为返回值，第二个参数为该数是否存在
	age, ano := m.Load("123")
	fmt.Println(age, ano)

	m.Range(func(key, value interface{}) bool {
		name := key
		age := value
		fmt.Println(name, age)
		return true
	})

	m.Delete("test")
	age, ok := m.Load("test")
	fmt.Println(age, ok)

	// 5. 读取或写入
	m.LoadOrStore("stefno", 100)
	age, _ = m.Load("stefno")
	fmt.Println(age)

}
