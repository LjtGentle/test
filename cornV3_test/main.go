package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	cronExpr := "* * * * * *" //每秒？
	cron := cron.New()
	cron.AddFunc(cronExpr, func() {
		fmt.Println("秒 time=",time.Now())
	})
	cron.AddFunc("* * * * *", func() {
		fmt.Println("分 time=",time.Now())
	})
	cron.AddFunc("0 22,23 * * *", func() {
		fmt.Println("22 时 23时 time=",time.Now())
	})
	cron.Run()
}
