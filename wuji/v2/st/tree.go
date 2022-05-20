package st

type TreeItem struct {
	Label     string // 标签 name(desc）
	FieldName string //字段名
	Positive  *TreeBranch
	Args      string //一些其他信息
	Type      string
}

type Tree struct {
	Items  []*TreeItem
	TreeID string //自己的树id
	Ns     []*Tree
}

// TreeBranch 如何关联
// 哪个字段关联 哪个棵树 的哪个字段
type TreeBranch struct {
	Field    string //字段
	BeField  string // 被关联的字段
	BeTreeID string // 被关联的树id
	BeTree	*Tree	// 被关联的树
}
