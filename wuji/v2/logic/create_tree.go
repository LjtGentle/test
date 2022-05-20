package logic

import "test/wuji/v2/st"

// NewOneTree 创建一棵树
func NewOneTree() *st.Tree {
	// 创建一些节点
	// A -> B
	aItem1 := &st.TreeItem{
		FieldName: "a1",
		Type:      "7",
	}
	// A -> E
	aItem2 := &st.TreeItem{
		FieldName: "a2",
		Type:      "7",
	}
	// 无关联关系的节点
	aItem3 := &st.TreeItem{
		FieldName: "a3",
		Type:      "3",
	}
	aTree := &st.Tree{
		Items:  []*st.TreeItem{aItem1,aItem2,aItem3},
		TreeID: "aTreeID",
	}

	// B -> D
	bItem1 := &st.TreeItem{
		FieldName: "b1",
		Type:      "7",
	}
	// C -> B
	bItem2 := &st.TreeItem{
		FieldName: "b2",
		Type:      "22",
	}
	// A -> B
	bItem3 := &st.TreeItem{
		FieldName: "b3",
		Type:      "33",
	}
	// not
	bItem4 := &st.TreeItem{
		FieldName: "b4",
		Type:      "4",
	}

	bTree := &st.Tree{
		Items:  []*st.TreeItem{bItem1,bItem2,bItem3,bItem4},
		TreeID: "bTreeID",
	}


	// C -> B
	cItem1 := &st.TreeItem{
		FieldName: "c1",
		Type:      "7",
	}
	// not
	cItem2 := &st.TreeItem{
		FieldName: "c2",
		Type:      "2",
	}
	cTree := &st.Tree{
		Items:  []*st.TreeItem{cItem1,cItem2},
		TreeID: "cTreeID",
	}

	// B -> D
	dItem1 := &st.TreeItem{
		FieldName: "d1",
		Type:      "11",
	}
	// F -> D
	dItem2 := &st.TreeItem{
		FieldName: "d2",
		Type:      "22",
	}
	// not
	dItem3 := &st.TreeItem{
		FieldName: "d3",
		Type:      "3",
	}
	dTree := &st.Tree{
		Items:  []*st.TreeItem{dItem1,dItem2,dItem3},
		TreeID: "cTreeID",
	}
	// A -> E
	eItem1 := &st.TreeItem{
		FieldName: "e1",
		Type:      "11",
	}
	//not
	eItem2 := &st.TreeItem{
		FieldName: "e2",
		Type:      "2",
	}
	eTree := &st.Tree{
		Items:  []*st.TreeItem{eItem1,eItem2},
		TreeID: "cTreeID",
	}
	// F -> D
	fItem1 := &st.TreeItem{
		FieldName: "f1",
		Type:      "7",
	}
	// not
	fItem2 := &st.TreeItem{
		FieldName: "f2",
		Type:      "2",
	}
	fTree := &st.Tree{
		Items:  []*st.TreeItem{fItem1,fItem2},
		TreeID: "cTreeID",
	}

	// 先处理正向关系
	aTreeBranchB := &st.TreeBranch{
		Field:    "a1",
		BeField:  "b3",
		BeTreeID: bTree.TreeID,
		BeTree:     bTree,
	}
	aItem1.Positive = aTreeBranchB

	aTreeBranchE := &st.TreeBranch{
		Field:    "a2",
		BeField:  "e1",
		BeTreeID: eTree.TreeID,
		BeTree:   eTree,
	}
	aItem2.Positive = aTreeBranchE

	bTreeBranchD := &st.TreeBranch{
		Field:    "b1",
		BeField:  "d1",
		BeTreeID: dTree.TreeID,
		BeTree:   dTree,
	}
	bItem1.Positive = bTreeBranchD

	cTreeBranchB := &st.TreeBranch{
		Field:    "c1",
		BeField:  "b2",
		BeTreeID: bTree.TreeID,
		BeTree:   bTree,
	}
	cItem1.Positive = cTreeBranchB

	fTreeBranchD := &st.TreeBranch{
		Field:    "f1",
		BeField:  "d2",
		BeTreeID: dTree.TreeID,
		BeTree:   dTree,
	}
	fItem1.Positive =fTreeBranchD
	// 处理反向关联
	bTree.Ns = []*st.Tree{aTree,cTree}
	dTree.Ns = []*st.Tree{bTree,fTree}
	eTree.Ns = []*st.Tree{aTree}
	return aTree
}

func NewSimpleTree() *st.Tree{
	// A -> B
	aItem1 := &st.TreeItem{
		FieldName: "a1",
		Type:      "7",
	}
	// A -> E
	aItem2 := &st.TreeItem{
		FieldName: "a2",
		Type:      "7",
	}
	// 无关联关系的节点
	aItem3 := &st.TreeItem{
		FieldName: "a3",
		Type:      "3",
	}
	aTree := &st.Tree{
		Items:  []*st.TreeItem{aItem1,aItem2,aItem3},
		TreeID: "aTreeID",
	}
	// B -> D
	bItem1 := &st.TreeItem{
		FieldName: "b1",
		Type:      "7",
	}
	// C -> B
	bItem2 := &st.TreeItem{
		FieldName: "b2",
		Type:      "22",
	}
	// A -> B
	bItem3 := &st.TreeItem{
		FieldName: "b3",
		Type:      "33",
	}
	// not
	bItem4 := &st.TreeItem{
		FieldName: "b4",
		Type:      "4",
	}
	// 正向
	bTree := &st.Tree{
		Items:  []*st.TreeItem{bItem1,bItem2,bItem3,bItem4},
		TreeID: "bTreeID",
	}
	aTreeBranchB := &st.TreeBranch{
		Field:    "a1",
		BeField:  "b3",
		BeTreeID: bTree.TreeID,
		BeTree:     bTree,
	}
	aItem1.Positive = aTreeBranchB
	// 反向
	bTree.Ns = []*st.Tree{aTree}

	// --- a b tree 测试可以，只有正向的部分没有问题

	// 现在加入 c tree
	// C -> B
	cItem1 := &st.TreeItem{
		FieldName: "c1",
		Type:      "7",
	}
	// not
	cItem2 := &st.TreeItem{
		FieldName: "c2",
		Type:      "2",
	}
	cTree := &st.Tree{
		Items:  []*st.TreeItem{cItem1,cItem2},
		TreeID: "cTreeID",
	}

	// 正向
	cTreeBranchB := &st.TreeBranch{
		Field:    "c1",
		BeField:  "b2",
		BeTreeID: bTree.TreeID,
		BeTree:   bTree,
	}
	cItem1.Positive = cTreeBranchB
	//反向
	bTree.Ns = []*st.Tree{aTree,cTree}
	return aTree
}