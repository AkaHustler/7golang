package gee

import "strings"

//前缀树节点信息
type node struct {
	pattern 	string  //待匹配路由，例如 /p/:lang
	part 		string  //路由中的一部分，例如 :lang
	children	[]*node //子节点，例如[doc,intro]
	isWild		bool	//是否精确匹配，part含有 : 或 * 时为true
}

//第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children{
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

//所有匹配成功的节点，用于查询
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

//节点插入
func (n *node) insert(pattern string, parts []string, heignt int)  {
	if len(parts) == heignt {
		n.pattern = pattern
		return
	}
	part := parts[heignt]
	child := n.matchChild(pattern)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, heignt + 1)
}

//节点查询
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
