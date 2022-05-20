package logic

import (
	"test/wuji/v2/st"
	"testing"
)

func TestFlattenTree(t *testing.T) {
	// 创建一棵树
	tree := NewOneTree()
	//tree := NewSimpleTree()
	dest := make(map[string]*st.TreeItem)
	FlattenTree(dest, tree, "", nil,nil)
	//t.Logf("dest=%+v", dest)
	i := 0
	for k, v := range dest {
		i++
		t.Logf("%d----k=%s,v=%+v\n", i, k, v)
	}
	t.Logf("tree.Ns=%+v\n",tree.Ns)


}
