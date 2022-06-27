package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

func main() {
	w := semaphore.NewWeighted(3)
	i := 0
	for {
		i++
		business(w, i)
	}
}

func business(w *semaphore.Weighted, i int) {
	w.Acquire(context.TODO(), 1)

	// 业务逻辑
	go func(i int) {
		defer w.Release(1)
		time.Sleep(2 * time.Second)
		fmt.Printf("i=%d,time=%v\n", i, time.Now())
	}(i)
}

//func main() {
//	param := map[string]interface{}{
//		"status": 0,
//		"amount": 100,
//	}
//	fmt.Println(param)
//}
