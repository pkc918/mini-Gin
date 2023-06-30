package coco

import "strings"

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// matchChild 匹配 part 和 当前子结点，能匹配到就返回，匹配不到，就nil
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// insert 将没有匹配到的子结点插入，构造前缀树
/*
这里就是有两种情况，比如 api/login/a 和 api/login/b
当遇到第一个 api/login/a 的时候，api 和 login 都会创建节点，
当第二次遇到的时候，就会直接进入对应的 children 匹配，如果匹配不到，继续创建节点
*/
func (n *node) insert(pattern string, parts []string, height int) {
	// height 初始值为0，当为0时，pattern可能为 /
	// 当前就是 root 节点
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	// 比如 [api, login]，这样依次取值
	part := parts[height]
	// 初始化的时候返回 nil
	child := n.matchChild(part)
	// 初始化阶段，没有匹配到，重新创建一个分支
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// matchChildren 搜集所有匹配成功的子节点，比如在  lang 节点下有a，b，c，nodes就会取出[a]，或者[a,b]等符合要求的
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// search 在 Trie 树上查找满足 parts 条件的节点
func (n *node) search(parts []string, height int) *node {
	// HasPrefix 判断某个字符串是否从某个字符开始
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		// 找到最后匹配成功的 node
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
