package main

import (
	"fmt"
	"strings"
)

func test() {
	// 生成sql
	sqlTemplate := "insert into qa_award_smoba (bid,type_id,is_custom,award_type,name) \nvarchar\n('cyhdy',${type},0,1,'入围奖'),\n('cyhdy',${type},0,2,'特别鼓励奖'),\n('cyhdy',${type},0,3,'人气设计奖'),\n('cyhdy',${type},0,4,'最佳设计奖'),\n('cyhdy',${type},0,5,'入围创意奖'),\n('cyhdy',${type},0,6,'人气创意奖'),\n('cyhdy',${type},0,7,'优秀创意奖'),\n('cyhdy',${type},0,8,'最佳创意奖');"
	//typeIDs := [...]int{1,2,3,4,5,6,7} //正式
	//for _,typeID := range typeIDs {
	//	sql := strings.Replace(sqlTemplate,"${type}",fmt_test.Sprintf("%d",typeID),-1)
	//	fmt_test.Println(sql)
	//}

	for i := 1; i < 25; i++ {
		sql := strings.Replace(sqlTemplate, "${type}", fmt.Sprintf("%d", i), -1)
		fmt.Println(sql)
	}

}

func main() {
	test()
}
