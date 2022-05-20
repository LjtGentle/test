package wuji

import (
	"fmt"
	"strings"
)

type TreeItem struct {
	Label     string // 标签 name(desc）
	FieldName string //字段名
	Relative  string
	Absolute  string
	Positive  *Tree
	Relation  string //args 修改这个值
	Type      string
}

type Tree struct {
	Items    []*TreeItem
	SchemaID string
	Ns       []*Tree
}

//type Negative struct {
//	Field string // 自己的字段
//	RelationTreeID string //关联的树id
//	RelationField string // 关联树的字段
//	Nt *Tree
//}
//管理端勾选的应该时一个 []string

//flattenTree将树扁平化 ，value时确定的存放treeItem key如何定义呢？
// P_TreeID_FieldName
func flattenTree(dest map[string]*TreeItem, tree *Tree, prefix string, pParent []*Tree) {
	//对于正向的操作
	for _, item := range tree.Items {
		if isContain(item.Positive, pParent) {
			continue
		}

		key := prefix + "/" + "P_" + tree.SchemaID + "_" + item.FieldName
		key = strings.TrimLeft(key, "/")
		fmt.Printf("key=%s,v=%+v\n",key,item)
		dest[key] = item
		if item.Positive != nil {
			pParent = append(pParent, item.Positive)
			for _, p:= range pParent {
				fmt.Printf("parent=%+v\n",*p)
			}
			flattenTree(dest, item.Positive, key, pParent)
			for _, p:= range pParent {
				fmt.Printf("2parent=%+v\n",*p)
			}
		}

	}
	//对于反向的操作
	for _, nTree := range tree.Ns {
		// bTree eTree
		if isContain(nTree, pParent) {
			continue
		}
		pParent = append(pParent,nTree)
		for _, item := range nTree.Items {
			if isContain(item.Positive,pParent) {
				continue
			}
			key := prefix + "/" + "N_" + tree.SchemaID +"("+ nTree.SchemaID + ")_" + item.FieldName// 反向关联的树id
			fmt.Printf("key=%s,v=%+v\n",key,item)
			dest[key] = item
			if item.Positive != nil {
				flattenTree(dest, item.Positive, key, pParent)
			}

		}

	}
}

// 判断o 是否在祖先当中
func isContain(o *Tree, pParent []*Tree) bool {

	for _, v := range pParent {
		if v == o {
			return true
		}
	}

	return false
}
