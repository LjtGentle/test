package wuji

import (
	"testing"
)

func TestFlattenTree(t *testing.T) {
	// 创建一个树
	aTree := &Tree{}
	//->B
	a1 := &TreeItem{
		Label:     "a1(A1)",
		FieldName: "a1",
		Type:      "7",
	}
	// ->E
	a2 := &TreeItem{
		Label:     "a2(A2)",
		FieldName: "a2",
		Type:      "7",
	}
	// not
	a3 := &TreeItem{
		Label:     "a3(A3)",
		FieldName: "a3",
		Type:      "3",
	}
	aTree.Items = append(aTree.Items, a1)
	aTree.Items = append(aTree.Items, a2)
	aTree.Items = append(aTree.Items, a3)
	bTree := &Tree{}
	// ->D
	b1 := &TreeItem{
		Label:     "b1(B1)",
		FieldName: "b1",
		Type:      "7",
	}
	//A->
	b2 := &TreeItem{
		Label:     "b2(B2)",
		FieldName: "b2",
		Type:      "2",
	}
	// C->
	b3 := &TreeItem{
		Label:     "b3(B3)",
		FieldName: "b3",
		Type:      "3",
	}
	// not
	b4 := &TreeItem{
		Label:     "b4(B4)",
		FieldName: "b4",
		Type:      "4",
	}
	bTree.Items = append(bTree.Items, b1)
	bTree.Items = append(bTree.Items, b2)
	bTree.Items = append(bTree.Items, b3)
	bTree.Items = append(bTree.Items, b4)

	cTree := &Tree{}
	// ->B
	c1 := &TreeItem{
		Label:     "c1(C1)",
		FieldName: "c1",
		Type:      "7",
	}
	// not
	c2 := &TreeItem{
		Label:     "c2(C2)",
		FieldName: "c2",
		Type:      "2",
	}
	// not
	c3 := &TreeItem{
		Label:     "c3(C3)",
		FieldName: "c3",
		Type:      "3",
	}
	cTree.Items = append(cTree.Items, c1)
	cTree.Items = append(cTree.Items, c2)
	cTree.Items = append(cTree.Items, c3)
	dTree := &Tree{}
	//B->
	d1 := &TreeItem{
		Label:     "d1(D1)",
		FieldName: "d1",
		Type:      "1",
	}
	// F->
	d2 := &TreeItem{
		Label:     "d2(D2)",
		FieldName: "d2",
		Type:      "2",
	}
	// not
	d3 := &TreeItem{
		Label:     "d3(D3)",
		FieldName: "d3",
		Type:      "3",
	}
	dTree.Items = append(dTree.Items, d1)
	dTree.Items = append(dTree.Items, d2)
	dTree.Items = append(dTree.Items, d3)
	eTree := &Tree{}
	//A->
	e1 := &TreeItem{
		Label:     "e1(E1)",
		FieldName: "e1",
		Type:      "1",
	}
	// not
	e2 := &TreeItem{
		Label:     "e2(E2)",
		FieldName: "e2",
		Type:      "2",
	}
	eTree.Items = append(eTree.Items, e1)
	eTree.Items = append(eTree.Items, e2)
	fTree := &Tree{}
	//->D
	f1 := &TreeItem{
		Label:     "f1(F1)",
		FieldName: "f1",
		Type:      "7",
	}
	// not
	f2 := &TreeItem{
		Label:     "f2(F2)",
		FieldName: "f2",
		Type:      "2",
	}
	fTree.Items = append(fTree.Items, f1)
	fTree.Items = append(fTree.Items, f2)

	// A 的正向关联
	a1.Positive = bTree
	a2.Positive = eTree
	//B 的正向关联
	b1.Positive = dTree
	//C 的正向关联
	c1.Positive = bTree
	// F的正向关联
	f1.Positive = dTree

	// B的反向关联
	bNegativeA := &Tree{}
	bNegativeA = aTree
	bNegativeD := &Tree{}
	bNegativeD = dTree

	//bNegativeA := &Negative{
	//	Field:          "b2",
	//	RelationTreeID: "aTree",
	//	RelationField:  "a1",
	//	Nt:             aTree,
	//}
	//bNegativeD := &Negative{
	//	Field:          "b3",
	//	RelationTreeID: "cTree",
	//	RelationField:  "c1",
	//	Nt:             cTree,
	//}
	bTree.Ns = append(bTree.Ns, bNegativeA)
	bTree.Ns = append(bTree.Ns, bNegativeD)
	// D 的反向关联
	dNegativeB := &Tree{}
	dNegativeB = bTree
	dNegativeF := &Tree{}
	dNegativeF = fTree
	//dNegativeB := &Negative{
	//	Field:          "d1",
	//	RelationTreeID: "bTree",
	//	RelationField:  "b1",
	//	Nt:             bTree,
	//}
	//
	//dNegativeF := &Negative{
	//	Field:          "d2",
	//	RelationTreeID: "fTree",
	//	RelationField:  "f1",
	//	Nt:             fTree,
	//}
	dTree.Ns = append(dTree.Ns, dNegativeB)
	dTree.Ns = append(dTree.Ns, dNegativeF)

	// E的反向关联
	//eNegativeA := &Negative{
	//	Field:          "e1",
	//	RelationTreeID: "aTree",
	//	RelationField:  "a2",
	//	Nt:             aTree,
	//}
	eNegativeA := aTree
	eTree.Ns = append(eTree.Ns, eNegativeA)
	aTree.SchemaID = "aTree"
	bTree.SchemaID = "bTree"
	cTree.SchemaID = "cTree"
	dTree.SchemaID = "dTree"
	eTree.SchemaID = "eTree"
	fTree.SchemaID = "fTree"

	dest := make(map[string]*TreeItem)
	flattenTree(dest, aTree, "",nil)
	var i int
	for k, v := range dest {
		i++
		t.Logf("i=%d  k=%s, v=%+v\n", i, k, v)
	}

}
