package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

const richText2 = "<p>\n  <br/>\n</p>\n<h1>哈哈</h1>\n<h2>笑什么</h2>\n<p>没有笑</p>\n<h1>2022年</h1>\n<h2>12月</h2>\n<p>20号</p>\n<p>\n  <br/>\n</p>"

func main() {
	test10()
}
func test10() {
	doc := getDoc()
	//获取h标签
	headings := extractHeadings10(doc, 0)
	printHeadings(headings, 0)
}

type Heading struct {
	Title    string
	Level    int
	Children []*Heading
}

func extractHeadings10(n *html.Node, level int) []*Heading {
	fmt.Println("level=", level)
	if n.Type == html.ElementNode && (n.Data == "h1" || n.Data == "h2" || n.Data == "h3" || n.Data == "h4" || n.Data == "h5" || n.Data == "h6") {
		heading := &Heading{Title: n.FirstChild.Data, Level: level, Children: []*Heading{}}
		return []*Heading{heading}
	}

	var result []*Heading
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//fmt.Println("1111")
		result = append(result, extractHeadings10(c, level+1)...)
	}

	return result
}

func printHeadings(headings []*Heading, level int) {
	for _, h := range headings {
		fmt.Printf("%s%d. %s\n", strings.Repeat("  ", level), level, h.Title)
		printHeadings(h.Children, level+1)
	}
}

func getDoc() *html.Node {
	doc, err := html.Parse(strings.NewReader(richText2))
	if err != nil {
		fmt.Println("html parse err=", err)
		return nil
	}
	return doc
}

func test09() {
	doc, err := html.Parse(strings.NewReader(richText2))
	if err != nil {
		fmt.Println("html parse err=", err)
		return
	}
	//获取h标签
	hNodes := extractHeadings09(doc, 0)

	printHeading(hNodes, 0)

}

type heading09 struct {
	level    int
	text     string
	children []*heading09
}

func (h *heading09) addChild(child *heading09) {
	h.children = append(h.children, child)
}

func printHeading(h *heading09, level int) {
	fmt.Printf("%*s%d. %s\n", level*4, "", h.level, h.text)
	for _, child := range h.children {
		printHeading(child, level+1)
	}
}

func extractHeadings09(n *html.Node, level int) *heading09 {
	if n.Type == html.ElementNode && (n.Data == "h1" || n.Data == "h2" || n.Data == "h3" || n.Data == "h4" || n.Data == "h5" || n.Data == "h6") {
		h := &heading09{
			level: level,
			text:  n.FirstChild.Data,
		}

		level++
		for c := n.NextSibling; c != nil; c = c.NextSibling {
			child := extractHeadings09(c, level)
			if child != nil {
				h.addChild(child)
			}
		}

		return h
	}

	var result *heading09
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		h := extractHeadings09(c, level)
		if h != nil {
			if result == nil {
				result = h
			} else {
				result.addChild(h)
			}
		}
	}
	return result
}

func test08() {
	/*
		package main

		import (
		    "fmt"
		    "io/ioutil"
		    "net/http"
		    "golang.org/x/net/html"
		)

		type Heading struct {
		    Level int
		    Text  string
		}

		type Node struct {
		    Heading
		    Children []*Node
		}

		func main() {
		    resp, err := http.Get("https://golang.org")
		    if err != nil {
		        panic(err)
		    }
		    defer resp.Body.Close()

		    b, err := ioutil.ReadAll(resp.Body)
		    if err != nil {
		        panic(err)
		    }

		    node, err := html.Parse(bytes.NewReader(b))
		    if err != nil {
		        panic(err)
		    }

		    root := &Node{}
		    current := root
		    extractHeadings(node, current)

		    printTree(root, "")
		}

		func extractHeadings(n *html.Node, current *Node) {
		    if n.Type == html.ElementNode {
		        if n.Data == "h1" || n.Data == "h2" || n.Data == "h3" || n.Data == "h4" || n.Data == "h5" || n.Data == "h6" {
		            level := 0
		            if n.Data == "h1" {
		                level = 1
		            } else if n.Data == "h2" {
		                level = 2
		            } else if n.Data == "h3" {
		                level = 3
		            } else if n.Data == "h4" {
		                level = 4
		            } else if n.Data == "h5" {
		                level = 5
		            } else if n.Data == "h6" {
		                level = 6
		            }
		            heading := &Node{Heading{level, n.FirstChild.Data}, nil}
		            if level <= current.Level {
		                for current.Level >= level {
		                    current = current.Parent
		                }
		                current.Children = append(current.Children, heading)
		                current = heading
		                heading.Parent = current
		            } else if level > current.Level {
		                current.Children = append(current.Children, heading)
		                current = heading
		                heading.Parent = current
		            }
		        }
		    }

		    for c := n.FirstChild; c != nil; c = c.NextSibling {
		        extractHeadings(c, current)
		    }
		}

		func printTree(node *Node, prefix string) {
		    fmt.Printf("%s%s\n", prefix, node.Text)
		    for _, child := range node.Children {
		        printTree(child, prefix+"  ")

	*/
}

func test07() {
	doc, err := html.Parse(strings.NewReader(richText2))
	if err != nil {
		fmt.Println("html parse err=", err)
		return
	}
	menu := extractHeadings03(doc, 0)
	printMenu(menu, 0)
}

func printMenu(menu []*MenuItem, level int) {
	for _, item := range menu {
		fmt.Printf("%s%s\n", strings.Repeat("  ", level), item.Name)
		printMenu(item.Subitems, level+1)
	}
}

type MenuItem struct {
	Name     string
	Subitems []*MenuItem
}

// stack overflow 栈溢出
func extractHeadings03(n *html.Node, level int) []*MenuItem {
	var result []*MenuItem
	if n.Type == html.ElementNode && (n.Data == "h1" || n.Data == "h2" || n.Data == "h3" || n.Data == "h4" || n.Data == "h5" || n.Data == "h6") {
		result = append(result, &MenuItem{
			Name:     n.FirstChild.Data,
			Subitems: extractHeadings03(n, level+1),
		})
		return result
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, extractHeadings03(c, level)...)
	}
	return result
}

func test06() {
	doc, err := html.Parse(strings.NewReader(richText2))
	if err != nil {
		fmt.Println("html parse err=", err)
		return
	}
	//获取h标签
	hNodes := extractHeadings02(doc, 0)
	for _, h := range hNodes {
		fmt.Printf("%s%s\n", h.indent, h.text)
	}
	fmt.Printf("hNodes=%+v\n", hNodes)
}

type heading struct {
	level  int
	text   string
	indent string
}

func extractHeadings02(n *html.Node, level int) []heading {
	if n.Type == html.ElementNode && (n.Data == "h1" || n.Data == "h2" || n.Data == "h3" || n.Data == "h4" || n.Data == "h5" || n.Data == "h6") {
		h := heading{
			level:  level + 1,
			text:   n.FirstChild.Data,
			indent: strings.Repeat("\t", level),
		}
		return []heading{h}
	}

	var result []heading
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		fmt.Println("----level=", level)
		result = append(result, extractHeadings02(c, level)...)
	}

	return result
}
func test05() {
	doc, err := html.Parse(strings.NewReader(richText2))
	if err != nil {
		fmt.Println("html parse err=", err)
		return
	}
	//获取h标签
	hNodes := extractHeadings(doc)
	for i, h := range hNodes {
		fmt.Printf("%d. %s\n", i+1, h)
	}
}

func extractHeadings(n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "h1" || n.Data == "h2" || n.Data == "h3" || n.Data == "h4" || n.Data == "h5" || n.Data == "h6" {
		return []string{n.FirstChild.Data}
	}

	var result []string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, extractHeadings(c)...)
	}

	return result
}
func test04() {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(richText2))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 查询所有的h1~h6标签
	doc.Find("h1, h2, h3, h4, h5, h6").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}

// faild
func test03() {
	doc, err := html.Parse(strings.NewReader(richText2))
	if err != nil {
		fmt.Println("html parse err=", err)
		return
	}
	var h1 []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h1" {
			// []
			fmt.Printf("n.Attr=%+v\n", n.Attr)
			for _, a := range n.Attr {
				fmt.Println("a.Key=", a.Key)
				if a.Key == "class" {
					h1 = append(h1, a.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	for i, h := range h1 {
		fmt.Printf("%d. %s\n", i+1, h)
	}
}

func test02() {

	doc, err := html.Parse(strings.NewReader(richText2))
	if err != nil {
		fmt.Println("html parse err=", err)
		return
	}
	// 遍历元素树
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h1" {
			fmt.Println(n.FirstChild.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			//fmt.Printf("c=%+v\n", c)
			f(c)
		}
	}
	f(doc)
}

func test01() {
	// Open the HTML file
	//file, err := os.Open("example.html")
	//if err != nil {
	//	fmt.Println("Error opening file:", err)
	//	os.Exit(1)
	//}
	//defer file.Close()
	//
	//// Create a new HTML parser
	//doc, err := html.Parse(file)
	//if err != nil {
	//	fmt.Println("Error parsing HTML:", err)
	//	os.Exit(1)
	//}

	//me
	doc, err := html.Parse(strings.NewReader(richText2))
	if err != nil {
		fmt.Println("html parse err=", err)
		return
	}

	// Search for h1, h2, h3, ... tags
	var hTags []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h1" {
			hTags = append(hTags, "h1")
		} else if n.Type == html.ElementNode && n.Data == "h2" {
			hTags = append(hTags, "h2")
		} else if n.Type == html.ElementNode && n.Data == "h3" {
			hTags = append(hTags, "h3")
		} else if n.Type == html.ElementNode && n.Data == "h4" {
			hTags = append(hTags, "h4")
		} else if n.Type == html.ElementNode && n.Data == "h5" {
			hTags = append(hTags, "h5")
		} else if n.Type == html.ElementNode && n.Data == "h6" {
			hTags = append(hTags, "h6")
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	// Print the h tags
	for _, h := range hTags {
		fmt.Println(h)
	}
}
