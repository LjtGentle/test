package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"math/rand"
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

func main() {
	test3()

}
