package main

import (
	"encoding/json"
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	test03()
}

func test04() {
	key := "SOP_ENV"
	//export SOP_ENV=prod

	//key := "GOROOT"
	env := os.Getenv(key)
	fmt.Println("env=", env)
}

const richText = " <p><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\"><img src=\"https://shp.qpic.cn/cfwebcap/0/f77a88689b3dbf61af412a80a176ee3d/0/?width=554&amp;height=255\" alt=\"\" /></span></p><p><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">与王者恰恰相反大乔的定位更像是一个法师</span></p><h1><br /><span style=\"font-size: 14px; font-family: 'Microsoft YaHei', 微软雅黑;\"><strong>一技能：沧溟之珠</strong></span></h1><p><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\"><img src=\"https://shp.qpic.cn/cfwebcap/0/170d13b03a7a660c957e6b6f4fdca4c1/0/?width=554&amp;height=255\" alt=\"\" /></span></p><p><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">基础击飞值16%最低7%递减值1%-2%</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">（注：连续使用一个技能伤害值会递减）</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">技能距离：三个身位</span><br /><br /><strong><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">解析：</span></strong><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">1.一技能总共有8个方向，用方向键控制</span></p><p><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">2.一技能在地面释放和空中释放距离不一样，空中释放比地面释放多半个球距离</span></p><p><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">3.地面释放会向上浮动，空中释放不会浮动</span></p><p><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">4.一技能可以配合三技能使用，配合三技能后球不会浮动并且存在时间刷新</span></p><p><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">5.一技能在场上只能存在一个，第一下释放后球还没有消失可以再按一次一技能将水球引爆，</span><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">水球消失后可以再按一次召唤一个新的水球</span></p><p><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">6.水球只有敌人触碰引爆和手动引爆（参考4），引爆过程中伤害从内向外扩散递减，是属于一个小范围技能会对周围敌人造成伤害</span></p><p><strong><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">水球使用技巧：</span></strong><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">1.卡位置</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">一技能在场可以威慑对手，比如一些空对地的技能就可以用来威慑对手，也可以分割战场阻挡对手，或者场外拦截卡位置</span></p><p><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">2.做防御</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">破晓有拼刀机制就是技能互相抵消，而大乔一技能是脱离本体以为的个体，在释放其他技能没有影响，而且技能对大乔没有影响，可以直接躲在一技能下面后者中间</span></p><h1><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">3.当道具</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">一技能释放完之后可以配合三技能再次移动他，这样就相当与场上的一个限时道具，释放在地面上是拾取道具，释放在地面下是地雷</span><br /><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">当然一技能的释放要在确保自身安全的情况下，一技能的前摇是很大的并且有滞空效果，如同活靶子，在空中释放的时候可以预判位置球不会下落。</span><br /><br /><span style=\"font-size: 14px; font-family: 'Microsoft YaHei', 微软雅黑;\"><strong>二技能：鲤跃之潮</strong></span></h1><p><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\"><img src=\"https://shp.qpic.cn/cfwebcap/0/94aec6a03a876b29d7228afbd2173551/0/?width=554&amp;height=255\" alt=\"\" /></span></p><p><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">基础击飞值</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">单点基础击飞值15%最低击飞值6%</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">蓄力基础击飞值25%最低击飞值11%</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">递减指1%-2%</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">技能距离：左右各一个身位</span></p><p><br /><strong><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">解析：</span></strong><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">1.大乔二技能类似其他角色的划a蓄力技能，触发方式为两种轻点触发和长按蓄力触发</span></p><p><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">2.身下没有伤害判断范围</span></p><p><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">3.前摇和后摇都很大</span></p><p>&nbsp;</p><p><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">4.技能是从内向外散发有一个前摇时间，特效没有到伤害不会到</span></p><p><br /><strong><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">使用技巧</span></strong><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">1.地对空</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">左右跳跃快速切换位置使用二技能应对空中对手，二技能的范围从内向外扩散，在二技能范围顶点去对撞对手有几率可以优先对手向下的攻击，二技能和一技能不一样没有拼刀机制不会被抵消但是会被打断。</span></p><h1><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">2.拦截</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">二技能的击飞方向是向上的众所周知地图是一个长方形水平线上向上击飞离暗心边境都是一样距离，侧边击飞还要看在地图哪一个位置了。二技能还有滞空效果这个滞空大于一技能，就很适合场外拦截，先预判对手落点在空中蓄力二技能接住对手再向上击飞出暗心边界</span><br /><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">二技能算有点鸡肋，基本比较厉害玩家都可以躲掉这个技能，而且后摇太大了，而我目前也并没有完全开放其用法，如果说是地对空的技能但是他优先也不高必须要卡的非常准用外环打，</span><br /><br /><span style=\"font-size: 14px; font-family: 'Microsoft YaHei', 微软雅黑;\"><strong>三技能：天涯之跃</strong></span></h1><p><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\"><img src=\"https://shp.qpic.cn/cfwebcap/0/6223ad192ad4a4c33b0d265320696ecc/0/?width=554&amp;height=277\" alt=\"\" /></span></p><p>&nbsp;</p><p><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">击飞值无</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">技能距离：8个身位</span></p><p><br /><strong><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">解析：</span></strong></p><p><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">1.三技能和一技能一样也是8个方向，方向盘回中与向上方向一样</span></p><p><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">2.三技能可以被打断不存在无敌效果，起点特效完全消失才可以</span></p><p><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">3.三技能在结束前可以微调位置</span></p><p><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">4.三技能可以配合一技能使用达到搬运一技能并刷新一技能的效果</span></p><p><br /><strong><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">使用技巧</span></strong><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">1.回场</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">三技能的距离是很远的而且消失期间无敌，是一个很不错的会场手段，但是要注意落脚点，三技能后摇大</span></p><p><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">2.搬运</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">三技能本身没有伤害，在配合一技能的时候就有了伤害，一个球也可以多次搬运直到打到对手为止，搬运最好是出其不意的时候不然很容易被对手预判盾住</span></p><h1><br /><span style=\"font-size: 14px; font-family: 'Microsoft YaHei', 微软雅黑;\"><strong>破晓技：惊潮之唤</strong></span></h1><p><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\"><img src=\"https://shp.qpic.cn/cfwebcap/0/8d8084448e050ff3d28bdfa1b5d4e1fd/0/?width=554&amp;height=255\" alt=\"\" /></span></p><p><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">基础击飞值75%</span></p><p><br /><strong><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">使用技巧：</span></strong><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">破晓技只有上面和前面有伤害判定，起点为大乔身后一格距离</span></p><p><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">1.板外击飞接大招</span></p><p><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">把对手击飞出板外之后卡边放一个大招，可以直接封对手走位，记得卡时间挂边就打不到了！</span></p><p><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">2.贴脸接大招</span><br /><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">大乔破晓技有上方判定，而且有击退效果，正面配合击退效果基本跑不掉</span></p><p>&nbsp;</p><p><span style=\"font-size: 12px; font-family: 'Microsoft YaHei', 微软雅黑;\">大乔主要还是用技能和平a配合拉扯来打伤害，也没有真正意义上的连招，希望这期可以帮助新手玩家打开对大乔的理解</span></p>"
const richText2 = "<p>\n  <br/>\n</p>\n<h1>哈哈</h1>\n<h2>笑什么</h2>\n<p>没有笑</p>\n<h1>2022年</h1>\n<h2>12月</h2>\n<p>20号</p>\n<p>\n  <br/>\n</p>"

const richText3 = "<p>\n  <br/>\n</p>\n<h1>哈哈</h1>\n<h2>笑什么</h2>\n<p>没有笑</p>\n<h1>2022年</h1>\n<h2>12月</h2>\n<p>20号</p>\n<h2>11月</h2>\n<p>\n  <br/>\n</p>"

const patternH1 = `<h1>([\S\s]*?)</h1>`
const patternH = `<h[1-6]>(.*?)</h[1-6]>`
const patterStrong = `<strong>([\S\s]*?)</strong>`

func test07() {
	tree := HTree{
		Child: nil,
		HData: HData{},
	}
	tree.Child = append(tree.Child, &HTree{
		Child: nil,
		HData: HData{
			Tag:     "h1",
			No:      1,
			Content: "hhhhh",
		},
	})

	fmt.Printf("tree child=%+v\n", *tree.Child[0])
	fmt.Printf("c=%+v\n", cap(tree.Child))
}

// 用html包匹配的情况
func test05() {
	doc, err := html.Parse(strings.NewReader(richText2))
	if err != nil {
		fmt.Println("html parse err=", err)
		return
	}
	// ok 问题是先全部匹配第一个
	expr := "//h1|//h2|//h3|//h4|//h5|//h6"
	//不行
	//expr := "//h1or//h2or//h3or//h4or//h5or//h6"
	//不行，括号的 //h1  h1 h1/h2
	//expr := "(//h1,//h2,//h3,//h4,//h5,//h6)"
	//expr := "(/h1,/h2,/h3,/h4,/h5,/h6)"
	//expr := "(h1,h2,h3,h4,h5,h6)"
	//expr := "(h1/h2/h3/h4/h5/h6)"
	//expr := "h1/h2/h3/h4/h5/h6"
	//expr := "//h1"
	//expr1 := "//h1"
	h1 := make([]string, 0)
	h2 := make([]string, 0)
	h3 := make([]string, 0)
	h4 := make([]string, 0)
	h5 := make([]string, 0)
	h6 := make([]string, 0)
	htmlNode := htmlquery.Find(doc, expr)
	for _, h := range htmlNode {

		fmt.Println(htmlquery.InnerText(h))
		fmt.Printf("h=%+v\n", h)
		fmt.Printf("h.Data=%+v\n", h.Data) // 打印出标签类型
		switch h.Data {
		case "h1":
			h1 = append(h1, htmlquery.InnerText(h))
		case "h2":
			h2 = append(h2, htmlquery.InnerText(h))
		case "h3":
			h3 = append(h3, htmlquery.InnerText(h))
		case "h4":
			h4 = append(h4, htmlquery.InnerText(h))
		case "h5":
			h5 = append(h5, htmlquery.InnerText(h))
		case "h6":
			h6 = append(h6, htmlquery.InnerText(h))
		}
		//fmt.Printf("h.FirstChild=%+v\n", h.FirstChild)
		//fmt.Printf("h.LastChild=%+v\n", h.LastChild)
		//fmt.Printf("h.NextSibling=%+v\n", h.NextSibling)
		//fmt.Printf("h.Parent=%+v\n", h.Parent)
		//fmt.Printf("h.PrevSibling=%+v\n", h.PrevSibling)
		fmt.Println("--------------------")
		//htmlNode2 := htmlquery.Find(h, "//h2")
		//for _, h2 := range htmlNode2 {
		//	fmt.Println(htmlquery.InnerText(h2))
		//	fmt.Printf("h2=%+v\n", h2)
		//	fmt.Printf("h2.Data=%+v\n", h2.Data) // 打印出标签类型
		//	fmt.Println("+++++++")
		//}
	}

}

// 测试 是否可以拆出h标签
func test06() {
	doc, err := html.Parse(strings.NewReader(richText2))
	if err != nil {
		fmt.Println("html parse err=", err)
		return
	}
	// 失败
	expr := "//h"
	htmlNode := htmlquery.Find(doc, expr)
	for _, h := range htmlNode {
		fmt.Println(htmlquery.InnerText(h))
	}
}

type HData struct {
	Tag     string
	No      int
	Content string
}

type HTree struct {
	Child []*HTree
	HData
}

// 正则表达式匹配的目录的情况
func test03() {
	reg := regexp.MustCompile(patternH)
	result := reg.FindAllString(richText3, -1)

	//list := make([]string, 0, len(result))
	fmt.Printf("result len=%+v\n", len(result))
	fmt.Printf("result= %+v\n", result)
	hDatas := make([]HData, 0, len(result))
	for _, r := range result {
		//regStrong := regexp.MustCompile(patterStrong)
		//re := regStrong.FindAllString(r, -1)
		//if len(re) > 1 {
		//	fmt.Println("err len=", len(re))
		//	return
		//}
		//re[0] = strings.TrimLeft(re[0], `<strong>`)
		//re[0] = strings.TrimRight(re[0], `</strong>`)
		//list = append(list, re[0])
		//fmt.Println("r=", r)
		//fmt.Println("content=", r[4:len(r)-5])
		//fmt.Println("标签=", r[1:3])
		no, _ := strconv.Atoi(r[2:3])
		hData := HData{
			Tag:     r[1:3],
			Content: r[4 : len(r)-5],
			No:      no,
		}
		hDatas = append(hDatas, hData)
	}
	//fmt.Printf("lits=%#v\n", list)
	////fmt.Printf("index str1 =%s\n", richText[352:392])
	//index := reg.FindAllStringIndex(richText2, -1)
	//fmt.Printf("index=%+v\n", index)
	fmt.Printf("hDatas=%+v\n", hDatas)

	hTree := HTree{
		Child: make([]*HTree, 0),
		HData: HData{},
	}
	addChild(&hTree, hDatas, 1)
	fmt.Printf("hTree=%v+\n", hTree)
	treeData, err := json.Marshal(hTree)
	if err != nil {
		return
	}
	fmt.Printf("treeString=%s\n", string(treeData))
}

// hTree ，hData 是输入也是输出
func addChild(hTree *HTree, hData []HData, no int) ([]HData, int) {
	index := 0
	if len(hData) < 2 {
		return nil, 0
	}
	fmt.Println("addChild run")

	for i := 0; i < len(hData); i++ {
		fmt.Println("111111")
		cHTree := &HTree{
			Child: nil,
			HData: HData{},
		}
		if hData[i].No == no {
			fmt.Printf("等于,hData[i]=%+v\n", hData[i])
			//cHTree = HTree{
			//	Child: make([]HTree, 0),
			//	HData: hData[i],
			//}
			cHTree.HData = hData[i]
			//// text2 注释了正常
			//hTree.Child = append(hTree.Child, cHTree)
			hTree.Child = append(hTree.Child, cHTree)
			fmt.Printf("++++hTree=%+v\n", hTree)
		}
		if hData[i].No < no {
			index = i
			hTree.Child = append(hTree.Child, cHTree)
			fmt.Printf("异常出现小于的情况，no=%d,i=%d,hData=%+v\n", no, i, hData[i])
			break
		}
		if hData[i].No > no {
			// 递归
			fmt.Printf("大于,hData[i]=%+v\n", hData[i])
			icHTree := &HTree{
				Child: nil,
				HData: hData[i],
			}
			_, index = addChild(icHTree, hData[i-1:], hData[i].No)
			index = index + i
			i = index
			fmt.Println("-------index=", index)
			fmt.Printf("-------icHTree=%+v\n", icHTree)

			cHTree.Child = append(cHTree.Child, icHTree)
			//text2 没注释正常
			//hTree.Child = append(hTree.Child, cHTree)
			fmt.Printf("-------cHTree =%+v\n", cHTree)
			continue
		}

	}
	return hData[index+1:], index
}

func test1() {
	reg, err := regexp.Compile(`\${[\w_]+\}`)
	if err != nil {
		fmt.Println(" regexp err")
	}
	txt := "smoba_weapon_#{stage}_${version}"

	strSlice := reg.FindAllString(txt, -1)
	for _, v := range strSlice {
		txt = strings.Replace(txt, v, "V73", -1)

	}
	fmt.Printf("strSlice=%v\n", strSlice)
	fmt.Printf("txt=%v\n", txt)
	//fmt_test.Printf("strSlice[0]=%v\n",templateGetValue(strSlice[0]))
	fmt.Println("len=", len(strSlice))
}

// #{stage} -> stage  ${version -> version
func templateGetValue(str string) string {
	b := []byte(str)
	return string(b[2 : len(b)-1])
}

func test2() {

	selector := "task"
	match, _ := regexp.MatchString("^([a-z]+,)*[a-z]+$", selector)

	fmt.Println(match)
}
