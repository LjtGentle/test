package wuji

import "strings"

type TreeItem struct {
	Label string // 标签 name(desc）
	FieldName string //字段名
	Relative string
	Absolute string
	Positive *Tree
	Relation string //args 修改这个值
	Type string
}

type Tree struct {
	Items []*TreeItem
	SchemaID string
	Ns []*Tree
}

//type Negative struct {
//	Field string // 自己的字段
//	RelationTreeID string //关联的树id
//	RelationField string // 关联树的字段
//	Nt *Tree
//}
//管理端勾选的应该时一个 []string


//flattenTree将树扁平化 ，value时确定的存放treeItem key如何定义呢？
func flattenTree(dest map[string]*TreeItem,tree *Tree,prefix string, parent...*Tree) {
	//对于正向的操作
	for _,item := range tree.Items {
		if isContain(item.Positive,parent...) {
			continue
		}

		key := prefix + "/" + tree.SchemaID+"_"+item.FieldName
		key = strings.TrimLeft(key,"/")
		dest[key] = item
		if item.Positive != nil {
			parent = append(parent,item.Positive)
			flattenTree(dest,item.Positive,key,parent...)
		}

	}
	//对于反向的操作
	for _,negative := range tree.Ns {
		if isContain(negative,parent...){
			continue
		}
		key := prefix +"/" + tree.SchemaID
		flattenTree(dest,negative,key,parent...)
	}
}


// 判断o 是否在祖先当中
func isContain(o *Tree, parent... *Tree) bool {
	for _, v := range parent {
		if v == o {
			return true
		}
	}
	return false
}

