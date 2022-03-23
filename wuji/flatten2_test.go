package wuji
//
//import "testing"
//
//func TestFlattenTree(t *testing.T) {
//	// 创建一个树
//	aTree := &FullTree{}
//	//->B
//	a1 := &TreeItem{
//		Label:     "a1(A1)",
//		FieldName: "a1",
//		Type:      "7",
//	}
//	// ->E
//	a2 := &TreeItem{
//		Label:     "a2(A2)",
//		FieldName: "a2",
//		Type:      "7",
//	}
//	// not
//	a3 := &TreeItem{
//		Label:     "a3(A3)",
//		FieldName: "a3",
//		Type:      "3",
//	}
//	aTree.Tree = append(aTree.Tree,a1)
//	aTree.Tree = append(aTree.Tree,a2)
//	aTree.Tree = append(aTree.Tree,a3)
//
//	bTree := &FullTree{}
//	// ->D
//	b1 := &TreeItem{
//		Label:     "b1(B1)",
//		FieldName: "b1",
//		Type:      "7",
//	}
//	//A->
//	b2 := &TreeItem{
//		Label:     "b2(B2)",
//		FieldName: "b2",
//		Type:      "2",
//	}
//	// C->
//	b3 := &TreeItem{
//		Label:     "b3(B3)",
//		FieldName: "b3",
//		Type:      "3",
//	}
//	// not
//	b4 := &TreeItem{
//		Label:     "b4(B4)",
//		FieldName: "b4",
//		Type:      "4",
//	}
//	bTree.Tree = append(bTree.Tree,b1)
//	bTree.Tree = append(bTree.Tree,b2)
//	bTree.Tree = append(bTree.Tree,b3)
//	bTree.Tree = append(bTree.Tree,b4)
//
//
//	cTree := &FullTree{}
//	// ->B
//	c1 := &TreeItem{
//		Label:     "c1(C1)",
//		FieldName: "c1",
//		Type:      "7",
//	}
//	// not
//	c2 := &TreeItem{
//		Label:     "c2(C2)",
//		FieldName: "c2",
//		Type:      "2",
//	}
//	// not
//	c3 := &TreeItem{
//		Label:     "c3(C3)",
//		FieldName: "c3",
//		Type:      "3",
//	}
//	cTree.Tree = append(cTree.Tree,c1)
//	cTree.Tree = append(cTree.Tree,c2)
//	cTree.Tree = append(cTree.Tree,c3)
//
//	dTree := &FullTree{}
//	//B->
//	d1 := &TreeItem{
//		Label:     "d1(D1)",
//		FieldName: "d1",
//		Type:      "1",
//	}
//	// F->
//	d2 := &TreeItem{
//		Label:     "d2(D2)",
//		FieldName: "d2",
//		Type:      "2",
//	}
//	// not
//	d3 := &TreeItem{
//		Label:     "d3(D3)",
//		FieldName: "d3",
//		Type:      "3",
//	}
//
//	dTree.Tree = append(dTree.Tree,d1)
//	dTree.Tree = append(dTree.Tree,d2)
//	dTree.Tree = append(dTree.Tree,d3)
//
//	eTree := &FullTree{}
//	//A->
//	e1 := &TreeItem{
//		Label:     "e1(E1)",
//		FieldName: "e1",
//		Type:      "1",
//	}
//	// not
//	e2 := &TreeItem{
//		Label:     "e2(E2)",
//		FieldName: "e2",
//		Type:      "2",
//	}
//
//	eTree.Tree = append(eTree.Tree,e1)
//	eTree.Tree = append(eTree.Tree,e2)
//	fTree := &FullTree{}
//	//->D
//	f1 := &TreeItem{
//		Label:     "f1(F1)",
//		FieldName: "f1",
//		Type:      "7",
//	}
//	// not
//	f2 := &TreeItem{
//		Label:     "f2(F2)",
//		FieldName: "f2",
//		Type:      "2",
//	}
//
//	fTree.Tree = append(fTree.Tree,f1)
//	fTree.Tree = append(fTree.Tree,f2)
//
//	// A 的正向关联
//	a1.Positive = bTree
//	a2.Positive = eTree
//	//B 的正向关联
//	b1.Positive = dTree
//	//C 的正向关联
//	c1.Positive = bTree
//	// F的正向关联
//	f1.Positive = dTree
//
//	// B的反向关联
//	bNegativeA := Tree{}
//	bNegativeA = aTree.Tree
//	bNegativeD := dTree.Tree
//	bTree.Negative = append(bTree.Negative,&bNegativeA)
//	bTree.Negative = append(bTree.Negative,&bNegativeD)
//
//	// D 的反向关联
//
//	dNegativeB := bTree.Tree
//	dNegativeF := fTree.Tree
//
//	dTree.Negative = append(dTree.Negative,&dNegativeB)
//	dTree.Negative = append(dTree.Negative,&dNegativeF)
//
//
//	// E的反向关联
//	eNegativeA := aTree.Tree
//	eTree.Negative = append(eTree.Negative,&eNegativeA)
//
//	//aTree.SchemaID = "aTree"
//	//bTree.SchemaID = "bTree"
//	//cTree.SchemaID = "cTree"
//	//dTree.SchemaID = "dTree"
//	//eTree.SchemaID = "eTree"
//	//fTree.SchemaID = "fTree"
//
//	dest := make(map[string]*TreeItem)
//	flattenTree(dest,aTree,"")
//	var i int
//	for k,v:= range dest {
//		i++
//		t.Logf("i=%d  k=%s, v=%+v\n",i,k,v)
//	}
//
//}
//
