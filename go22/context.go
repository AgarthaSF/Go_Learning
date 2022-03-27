package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func watchDog(ctx context.Context, name string) {
	//开启for select循环，一直后台监控
	for {
		select {
		case <-ctx.Done():
			fmt.Println("received instruction")
			return
		default:
			fmt.Println(name, "monitoring...")
		}
		time.Sleep(1 * time.Second)
	}
}

func getUser(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("get User", " goroutine exits")
			return
		default:
			userId := ctx.Value("userId")
			fmt.Println("get User, ID: ", userId)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(4)

	ctx, stop := context.WithCancel(context.Background())
	valCtx := context.WithValue(ctx, "userId", 2)

	go func() {
		defer wg.Done()
		watchDog(ctx, "watchDog1")
	}()

	go func() {
		defer wg.Done()
		watchDog(ctx, "watchDog2")
	}()

	go func() {
		defer wg.Done()
		watchDog(ctx, "watchDog3")
	}()

	go func() {
		defer wg.Done()
		getUser(valCtx)
	}()

	time.Sleep(5 * time.Second)
	stop()
	wg.Wait()

}
