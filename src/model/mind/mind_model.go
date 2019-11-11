package mind

type Mind struct {
	Node []MindNode `json:"roots"`
}

/**
 * 定义脑图数据结构: Mind
 */
type MindNode struct {
	Label string	`json:"label"`
	Side string		`json:"side,omitempty"`
	Children []MindNode	`json:"children,omitempty"`
}
