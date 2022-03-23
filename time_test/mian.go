package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func test() {
	//当前时间戳
	t := time.Now().Unix()
	fmt.Println("t=",t)
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
	message := fmt.Sprintf("%d|%s|%s|%d|%s|%d|%s",appid,appsecret,QauthSeq,QauthTimestamp,QauthUserDomain,QauthUserId,QauthUserOpenid)

	h.Write([]byte(message))
	bytes := h.Sum(nil)
	hs := hex.EncodeToString(bytes)
	fmt.Println("hs=",hs)

}

func test2() {
	t := time.Now()
	fmt.Println(t)
}


func main() {
	test2()

}