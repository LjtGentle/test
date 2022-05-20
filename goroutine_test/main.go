package main

func test() {
	// 测试sync.WaitGroup 中途return会如何
	for i:=0;i<10;i++ {
		go func(i int) {

		}(i)
	}
}


func main() {

}