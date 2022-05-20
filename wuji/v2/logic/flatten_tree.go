package logic

import (
	"fmt"
	"test/wuji/v2/st"
)

func FlattenTree(dest map[string]*st.TreeItem,tree *st.Tree,prefix string, parents, nparents []*st.Tree)  {
	// 将根节点入队
	if parents == nil {
		parents = append(parents,tree)
	}
	// 先处理正向关联
	for _,item := range tree.Items {
		if item.Positive != nil {
			if isContain(item.Positive.BeTree,parents) {
				continue
			}
			parents = append(parents, item.Positive.BeTree)
		}
		key := prefix + "/P/"+item.FieldName
		dest[key] = item
		fmt.Printf("P------key=%s,value=%+v\n",key,dest[key])
		if item.Positive != nil {
			FlattenTree(dest,item.Positive.BeTree,key,parents,nparents)
		}
	}
	// 后处理反向关联
	prefix += fmt.Sprintf("(%s)",tree.TreeID)
	for _,nTree := range tree.Ns {
		if isContain(nTree,parents) {
			continue
		}
		parents = append(parents,nTree)
		pre := prefix + fmt.Sprintf("[%s]",nTree.TreeID)
		for _,item := range nTree.Items {
			if item.Positive != nil {
				if isContain(item.Positive.BeTree,parents) {
					key := pre + "/N/"+item.FieldName
					dest[key] = item
					fmt.Printf("CN------key=%s,value=%+v\n",key,dest[key])
					continue
				}
				parents = append(parents,item.Positive.BeTree)
			}
			key := pre + "/N/"+item.FieldName
			dest[key] = item
			fmt.Printf("N------key=%s,value=%+v\n",key,dest[key])
			if item.Positive != nil {
				FlattenTree(dest,item.Positive.BeTree,key,parents,nparents)
			}
		}
	}
}




// 判断tree 是否已经存在在parents中
func isContain(o *st.Tree, parents []*st.Tree) bool {
	for _, p := range parents {
		if o == p {
			return true
		}
	}
	return false
}
