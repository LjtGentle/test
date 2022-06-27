package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	fmt.Println("come in main")
	job()

	fmt.Println("out of main")

}
func job () {
	cron := cron.New()
	cron.AddFunc("@every 1m", func() {
		fmt.Println("@every 1m time=",time.Now())
	})
	cron.AddFunc("g", func() {
		fmt.Println("分 time=",time.Now())
	})
	cron.AddFunc("0 22,23 * * *", func() {
		fmt.Println("22 时 23时 time=",time.Now())
	})
	cron.Run()
}