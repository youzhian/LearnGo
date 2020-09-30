package main

type Node struct {
	key      string
	value    string
	hashcode int
	next     *Node
}

func newNode(key string, value string, hashCode int, next *Node) *Node {
	return &Node{key: key, value: value, hashcode: hashCode, next: next}
}

func (n *Node) setValue(value string) {
	n.value = value
}
func (n *Node) getValue() string {
	return n.value
}
func (n *Node) getKey() string {
	return n.key
}

func (n *Node) hashCode() int {
	return getHashCode(n.key) ^ getHashCode(n.value)
}

func getHashCode(k string) int {
	if len(k) == 0 {
		return 0
	}
	var hashCode int = 0
	for i := range k {
		hashCode = hashCode*31 + int(k[i])
	}
	return hashCode
}
