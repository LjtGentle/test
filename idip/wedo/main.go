package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	test01()
}

type HeroInfo struct {
	ID      int
	Name    string
	Level   int
	GetTime time.Time
	Power   int
}
type PageInfo struct {
	PageNo      int
	TotalPageNo int
}

func test01() {
	response := "result=0&error_info=&list=22|10 李白 1 1645089427 0|11 李信 1 1645089427 0|12 大乔 1 1645089427 0|13 安琪拉 1 1645089427 0|14 马可波罗 1 1645089427 0|18 铠 1 1645089394 0|20 程咬金 1 1645089427 0|21 程咬金 1 1645089427 0|25 小乔 1 1645089427 0|31 鲁班七号 1 1645089427 0|32 猩红守护者 1 1645089427 0|36 孙悟空 1 1645089427 0|37 澜 1 1645089427 0|39 铠 1 1645089427 0|40 李白 1 1645089427 0|41 安琪拉 1 1645089427 0|42 鲁班七号 1 1645089427 0|47 黄金分割率 1 1645089427 0|55 李白 1 1645089427 0|56 大乔 1 1645089427 0|59 小乔 1 1645089427 0|79 孙悟空 1 1645089427 0|&page_no=1&totalpageno=1"
	pageInfo := new(PageInfo)
	heroesInfo := make([]HeroInfo, 0)
	rspData := strings.Split(response, "&")
	for _, item := range rspData {
		itemArr := strings.Split(item, "=")
		if len(itemArr) >= 2 {
			key, val := itemArr[0], itemArr[1]
			switch key {
			case "result":
				if val != "0" {
					log.Println("wedo GetUserHaveHeroes idip failed,rspData=", rspData)
					fmt.Println("wedo GetUserHaveHeroes idip failed")
					return
				}
			case "list":
				infos := strings.Split(val, "|")
				if len(infos) < 2 {
					fmt.Println("<2")
					return
				}
				length, _ := strconv.Atoi(infos[0])
				heroesInfo = make([]HeroInfo, 0, length)
				for i := 1; i < length; i++ {
					hero := strings.Split(infos[i], " ")
					heroID, _ := strconv.Atoi(hero[0])
					heroLevel, _ := strconv.Atoi(hero[2])
					getTimeStamp, _ := strconv.Atoi(hero[3])
					getTime := time.Unix(int64(getTimeStamp), 0)
					//fmt.Println("getTimeStamp=", getTimeStamp)
					//
					//fmt.Println("getTime1=", getTime)
					//getTime = time.UnixMicro(int64(getTimeStamp))
					//fmt.Println("getTime2=", getTime)
					heroPower, _ := strconv.Atoi(hero[4])

					heroInfo := HeroInfo{
						ID:      heroID,
						Name:    hero[1],
						Level:   heroLevel,
						GetTime: getTime,
						Power:   heroPower,
					}
					heroesInfo = append(heroesInfo, heroInfo)
				}
			case "page_no":
				ival, _ := strconv.Atoi(val)
				pageInfo.PageNo = ival
			case "totalpageno":
				ival, _ := strconv.Atoi(val)
				pageInfo.TotalPageNo = ival
			case "last_season_best_duan_id":

			}
		}
	}

	fmt.Printf("pageInfo=%+v\n", pageInfo)
	fmt.Printf("heroesInfo=%+v\n", heroesInfo)
}
