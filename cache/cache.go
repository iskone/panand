package cache

import "time"

var cacheTtl = 60

var cache = &Node{
	Parent:    nil,
	Id:        "root",
	Name:      "",
	Ttl:       -1,
	ChildNode: map[string]*Node{},
}

type Node struct {
	Parent    *Node
	Id        string
	Ttl       int
	Name      string
	File      bool
	ChildNode map[string]*Node
	Files     map[string]*Node
}

func (n *Node) Add(node *Node) {
	node.Parent = n
	if node.File {
		if n.Files == nil {
			n.Files = make(map[string]*Node)
		}
		n.Files[node.Name] = node
	} else {
		if n.ChildNode == nil {
			n.ChildNode = make(map[string]*Node)
		}
		n.ChildNode[node.Name] = node
	}

}
func (n *Node) Get(path ...string) (*Node, []string) {
	if len(path) == 0 {
		return cache, nil
	}
	var f = n
	for k, v := range path {
		if node, ok := f.ChildNode[v]; !ok {
			return f, path[k:]
		} else {
			f = node
		}
	}
	return f, nil
}

func (n *Node) Clean() {
	if n.Ttl != -1 && n.Parent != nil && int(time.Now().Unix()) > n.Ttl {
		delete(n.Parent.ChildNode, n.Name)
	} else {
		for _, v := range n.ChildNode {
			v.Clean()
		}
	}

}
func Get(path ...string) (*Node, []string) {

	return cache.Get(path...)
}
func Add(node *Node) () {
	cache.Add(node)
}
func SetHome(id string) {
	cache.Id = id
}
