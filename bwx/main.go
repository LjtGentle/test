package main

import (
	"errors"
	"fmt"
	proxy "git.code.oa.com/iegm-open/go-http-proxy"
	cronV3 "github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

const (
	WebHook      = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"
	Key          = "89889bfc-8f6f-46dc-96d7-7a2c002c024c"
	MarkDownType = "markdown"
	TextType     = "text"
	TimeLayout   = "2006-01-02 15:04:05"
)

var wxProxy *proxy.Proxy

func init() {
	var err error
	wxProxy, err = proxy.New(proxy.Options{
		ConnectStr: fmt.Sprintf(WebHook, Key),
	})
	if err != nil {
		panic(fmt.Sprintf("init wxProxy err=%+v", err))
	}
}

type SandMessage struct {
	MsgType string `json:"msgtype"`
	Text    `json:"text"`
}

type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

//WorkClock 早上提醒打卡
func WorkClock(msgType string, content string) error {
	msg := SandMessage{
		MsgType: msgType,
		Text: Text{
			Content:       content,
			MentionedList: []string{"@all"},
		},
	}
	var res interface{}
	code, err := wxProxy.PostJson("", &msg, &res)
	if err != nil {
		return err
	}
	if code != 200 {
		return errors.New(fmt.Sprintf("sand MorningClock failed code=%d\n", code))
	}
	return nil
}

func MorningClock() {
	// 上班打卡提醒
	msg := "今天又是充满希望的一天，可爱的同事们别忘记先打上班卡了喔,http://daka.woa.com/"
	err := WorkClock(TextType, msg)
	if err != nil {
		log.Printf("call morning clock err=%+v\n", err)
		return
	}
	log.Println("MorningClock worked!")
}

func AfterWorkClock() {
	msg := "辛苦工作了一天，可爱的同事们别忘记先打下班卡了喔,http://daka.woa.com/"
	err := WorkClock(TextType, msg)
	if err != nil {
		log.Printf("call work after clock err=%+v\n", err)
		return
	}
	log.Println("work after clock worked!")
}

func EatDinnerClock() {
	msg := "干饭拉，干饭拉，干饭要积极!"
	err := WorkClock(TextType, msg)
	if err != nil {
		log.Printf("call eat dinner clock err=%+v\n", err)
		return
	}
	log.Println("eat dinner clock worked!")
}

func CronJson() {
	log.Println("开始调度任务" + time.Now().Format(TimeLayout))
	cron := cronV3.New()

	cron.AddFunc("9 * * *", func() {
		log.Println("定时器提醒早上打卡")
		MorningClock()
	})

	cron.AddFunc("18 * * *", func() {
		log.Println("定时器提醒吃晚饭")
		EatDinnerClock()
	})

	cron.AddFunc("19 * * *", func() {
		log.Println("定时器提醒下午下班打卡")
		AfterWorkClock()
	})
	cron.Start()
}
func initLogger() {
	logFile, err := os.OpenFile("./clock.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
func main() {
	initLogger()
	CronJson()
}
