package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func test() {
	// 测试sync.WaitGroup 中途return会如何
	for i := 0; i < 10; i++ {
		go func(i int) {

		}(i)
	}
}

func test1() {
	// 测试goroutine 传入slice的问题
	w := semaphore.NewWeighted(2)

	for {
		testSlice := randInt()
		w.Acquire(context.TODO(), 1)

		go func(is []int) {
			defer w.Release(1)
			fmt.Printf("is=%+v\n", is)
			time.Sleep(1 * time.Second)
		}(testSlice)
	}

}
func randInt() []int {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(10)
	ts := make([]int, 0, 3)
	ts = append(ts, i)
	ts = append(ts, i)
	ts = append(ts, i)
	fmt.Printf("in randInt ts=%+v\n", ts)
	return ts
	//filter_tag=100,bill_id=222026|filter_tag=100,bill_id=222027|filter_tag=100,bill_id=222042|filter_tag=100,bill_id=222041|filter_tag=100,bill_id=222029|filter_tag=100,bill_id=222028|filter_tag=100,bill_id=222030|filter_tag=100,bill_id=222043|filter_tag=100,bill_id=222044|filter_tag=100,bill_id=222031|filter_tag=100,bill_id=222033|filter_tag=100,bill_id=222045|filter_tag=100,bill_id=222047|filter_tag=100,bill_id=222048|filter_tag=100,bill_id=222032|filter_tag=100,bill_id=222046|filter_tag=100,bill_id=222034|filter_tag=100,bill_id=222049|filter_tag=100,bill_id=222052|filter_tag=100,bill_id=222051|filter_tag=100,bill_id=222035|filter_tag=100,bill_id=222050|filter_tag=100,bill_id=222036|filter_tag=100,bill_id=222054|filter_tag=100,bill_id=222037|filter_tag=100,bill_id=222053|filter_tag=100,bill_id=222056|filter_tag=100,bill_id=222055|filter_tag=100,bill_id=222057|filter_tag=100,bill_id=222039|filter_tag=100,bill_id=222058|filter_tag=100,bill_id=222038|filter_tag=100,bill_id=222042|filter_tag=100,bill_id=222060|filter_tag=100,bill_id=222041|filter_tag=100,bill_id=222059|filter_tag=100,bill_id=222062|filter_tag=100,bill_id=222044|filter_tag=100,bill_id=222061|filter_tag=100,bill_id=222043|filter_tag=100,bill_id=222064|filter_tag=100,bill_id=222063|filter_tag=100,bill_id=222046|filter_tag=100,bill_id=222066|filter_tag=100,bill_id=222045|filter_tag=100,bill_id=222047|filter_tag=100,bill_id=222065|filter_tag=100,bill_id=222048|filter_tag=100,bill_id=222067|filter_tag=100,bill_id=222049|filter_tag=100,bill_id=222050|filter_tag=100,bill_id=222052|filter_tag=100,bill_id=222051|filter_tag=100,bill_id=222054|filter_tag=100,bill_id=222055|filter_tag=100,bill_id=222053|filter_tag=100,bill_id=222058|filter_tag=100,bill_id=222056|filter_tag=100,bill_id=222060|filter_tag=100,bill_id=222057|filter_tag=100,bill_id=222059|filter_tag=100,bill_id=222061|filter_tag=100,bill_id=222062|filter_tag=100,bill_id=222064|filter_tag=100,bill_id=222063|filter_tag=100,bill_id=222066|filter_tag=100,bill_id=222065|filter_tag=100,bill_id=222069|filter_tag=100,bill_id=222067|filter_tag=100,bill_id=222071|filter_tag=100,bill_id=222073|filter_tag=100,bill_id=222070|filter_tag=100,bill_id=222072|filter_tag=100,bill_id=222075|filter_tag=100,bill_id=222077|filter_tag=100,bill_id=222074|filter_tag=100,bill_id=222076|filter_tag=100,bill_id=222079|filter_tag=100,bill_id=222078|filter_tag=100,bill_id=222083|filter_tag=100,bill_id=222080|filter_tag=100,bill_id=222081|filter_tag=100,bill_id=222082|filter_tag=100,bill_id=222084|filter_tag=100,bill_id=222090|filter_tag=100,bill_id=222091|filter_tag=100,bill_id=222093|filter_tag=100,bill_id=222092|filter_tag=100,bill_id=222107|filter_tag=100,bill_id=222106|filter_tag=100,bill_id=222109|filter_tag=100,bill_id=222108|filter_tag=100,bill_id=222111|filter_tag=100,bill_id=222112|filter_tag=100,bill_id=222114|filter_tag=100,bill_id=222113|filter_tag=100,bill_id=222115

}

func test3() {
	func() {
		defer fmt.Println("defer 111")
		defer fmt.Println("defer 222")
	}()

}

// test04 退出多个线程
func test04() {
	done := make(chan struct{}, 1)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			for {
				select {
				case <-done:
					fmt.Printf("goroutine %d return..\n", number)
					done <- struct{}{}
					return
				default:
					fmt.Printf("goroutine %d sleep..\n", number)
					time.Sleep(1 * time.Second)
				}
			}
		}(i)
	}
	time.Sleep(2 * time.Second)
	done <- struct{}{}

	wg.Wait()
	time.Sleep(2 * time.Second)
	fmt.Println("123456")
}

func test05() {
	// 无缓存出现死锁
	done := make(chan struct{}, 0)
	done <- struct{}{}
	_ = <-done
}

func test06() {
	// 关闭channel 达到多协程退出
	done := make(chan struct{}, 1)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			for {
				select {
				case _, ok := <-done:
					if !ok {
						fmt.Printf("退出 %d\n", number)
						return
					}

				default:
					fmt.Printf("goroutine %d sleep..\n", number)
					time.Sleep(1 * time.Second)
				}
			}
		}(i)
	}
	time.Sleep(2 * time.Second)
	close(done)
	fmt.Println("关闭了chan")
	wg.Wait()
	time.Sleep(2 * time.Second)
	fmt.Println("123456")
}
func test07() {
	ctx := context.TODO()
	ctx.Done()

}

// 控制线程的数量
func test08() {
	sem := semaphore.NewWeighted(10)
	max := math.MaxInt64
	for i := 0; i < max; i++ {
		sem.Acquire(context.TODO(), 1)

		go func() {
			defer sem.Release(1)

			fmt.Printf("i=%d,num=%d\n", i, runtime.NumGoroutine())
		}()
	}
}

// 模拟执行业务的 goroutine
func doBusiness(ch chan bool, i int) {
	fmt.Println("go func:", i, "goroutine count:", runtime.NumGoroutine())
	<-ch
}

func test09() {
	max_cnt := math.MaxInt64
	// max_cnt := 10
	fmt.Println(max_cnt)

	ch := make(chan bool, 3)

	for i := 0; i < max_cnt; i++ {
		ch <- true
		go doBusiness(ch, i)
	}
}

func test10() {
	es := make([]error, 0, 10)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			es = append(es, fmt.Errorf("i=%d", num))
		}(i)
	}
	wg.Wait()
	fmt.Printf("es=%+v\n", es)
}

func test11() {
	//c := sync.Cond{}
	//c.Broadcast()
	//c.Wait()
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//wg.Done()
	//
	//av := atomic.Value{}
	//av.CompareAndSwap()
	//
	//p := sync.Pool{}
	//p.Get()
	//t := time.Now()
	//go func() {
	//	fmt.Println("", t.Sub(time.Now()).Nanoseconds())
	//}()
	//time.Sleep(1 * time.Second)
	dtstatdate := time.Now().Format("20060102")
	fmt.Println(dtstatdate)
}

func main() {
	test11()

}
