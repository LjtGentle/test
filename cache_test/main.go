package main

import (
	"fmt"
	"github.com/coocood/freecache"
	"github.com/jellydator/ttlcache/v2"
	"time"
)

// ttl缓存，只要有读取的时候就会被刷新ttl的时间 有获取缓存，缓存就不过期
func test1() {
	cache := ttlcache.NewCache()

	cache.SetWithTTL("testKey", 1234, 5*time.Second)
	time.Sleep(6 * time.Second)

	for {

		time.Sleep(1 * time.Second)
		v, t, err2 := cache.GetWithTTL("testKey")
		if err2 != nil {
			fmt.Println("err2=", err2)
		} else {
			fmt.Println("剩余时间=", t, "值=", v)
		}
		if val, err := cache.Get("testKey"); err == nil {
			fmt.Println(val)
		} else {
			fmt.Printf("err=%+v\n", err)
		}

	}

}

// 不管是否有人访问，时间到就刷新
func test2() {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	key := []byte("key")
	err := cache.Set(key, []byte("value"), 5)
	if err != nil {
		fmt.Println("set error=", err)
		return
	}
	for {
		time.Sleep(1 * time.Second)
		got, err := cache.Get(key)
		if err != nil {
			fmt.Println("get err=", err)
			return
		} else {
			fmt.Printf("%s\n", got)
		}
	}
}

func main() {
	test2()
}
