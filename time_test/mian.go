package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	test08()

}

func test08() {
	t := time.Now()
	ss := t.Format("20060102")
	fmt.Println(ss)
}

func test() {
	//当前时间戳
	t := time.Now().Unix()
	fmt.Println("t=", t)
	h := sha256.New()
	//"Qauth-User-Id": 123,
	//"Qauth-User-Openid": "openid",
	//"Qauth-User-Name": "name",
	//"Qauth-User-Domain": "OA",
	//"Qauth-timestamp": 1648000965,
	//"Qauth-seq": "123456",
	//"Qauth-sign": "fed9b567426dec4f33016c048f8e8c5a01a25036fb50b79767582ce8a35d9f0f",
	appid := 21
	appsecret := "bddf99dbd1a06d7a69d6c5da129b9b0e"
	QauthSeq := "123456"
	QauthTimestamp := t
	QauthUserDomain := "OA"
	QauthUserId := 123
	QauthUserOpenid := "openid"
	//{appid}|{appsecret}|{Qauth-seq}|{Qauth-timestamp}|{Qauth-User-Domain}|{Qauth-User-Id}|{Qauth-User-Openid}
	message := fmt.Sprintf("%d|%s|%s|%d|%s|%d|%s", appid, appsecret, QauthSeq, QauthTimestamp, QauthUserDomain, QauthUserId, QauthUserOpenid)

	h.Write([]byte(message))
	bytes := h.Sum(nil)
	hs := hex.EncodeToString(bytes)
	fmt.Println("hs=", hs)

}

func test2() {
	t := time.Now().Unix()
	//1648608103
	fmt.Println(t)
}

func test3() {
	t1 := "2022-01-07 00:00:00"
	t2 := "2022-01-13 00:00:00"
	st1, _ := time.ParseInLocation("2006-01-02 15:04:05", t1, time.Local)
	st2, _ := time.ParseInLocation("2006-01-02 15:04:05", t2, time.Local)
	fmt.Println("st1=", st1)
	fmt.Println("st2=", st2)
	flag := st1.Before(st2)
	fmt.Println("flag=", flag)
	flag = st2.After(st1)
	fmt.Println("flag=", flag)

}

func test07() {
	var ts int64 = 1669132800
	t := time.UnixMicro(ts)
	fmt.Println("t=", t)
	time := time.Unix(ts, 0)
	fmt.Println("time=", time)
}

// 2021-04-08T00:00:00Z 转时间戳
func test06() {
	t := time.Now()
	fmt.Println("t=", t)
	strTime := "2021-04-08T00:00:00Z"
	loc, _ := time.LoadLocation("UTC")
	layout := "2006-01-02T15:04:05Z"
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", strTime, loc)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("the=", theTime)
	u := theTime.Unix()
	fmt.Println("u=", u)

	t, err = time.Parse(layout, strTime)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("parse t=", t)
	t = t.Add(-8 * time.Hour)
	u = t.In(time.Local).Unix()

	fmt.Println("u=", u)
	fmt.Println("t=", time.Unix(u, 0))
}

// 获取距离凌晨的时间
func test05() {
	//
	//now := time.Now()
	//nextDay := now.AddDate(0, 0, 1)
	//fmt.Println("now=", now, "\tnextDay=", nextDay)
	//now := time.Now()
	//endTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location()).Unix()
	//fmt.Println("endTime-now=", endTime-now.Unix())
	now := time.Now()
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location()).Unix()
	var dTime time.Duration
	dTime = time.Duration(endTime-now.Unix()) * time.Second
	fmt.Println("dTime=", dTime)
}

func test4() {
	// 验证 1626940541412 是否是时间戳
	// 1648608103 现在的时间戳
	// 1648607913432  id   1648607913 432
	timeStamp := 1648607913
	t := time.Unix(int64(timeStamp), 0)
	fmt.Println("t=", t)
}
