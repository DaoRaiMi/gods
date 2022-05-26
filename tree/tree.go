package tree

type Node struct {
	left *Node
	right *Node
	data interface{}
}



func createBinTree() *Node {
	var (
		a = &Node{data: "A"}
		b = &Node{data: "B"}
		c = &Node{data: "C"}
		d = &Node{data: "D"}
		e = &Node{data: "E"}
		f = &Node{data: "F"}
		g = &Node{data: "G"}
		h = &Node{data: "H"}
		i = &Node{data: "I"}
		j = &Node{data: "J"}
		k = &Node{data: "K"}
	)
	a.left,a.right = b,c
	b.left,b.right = d,e
	d.left = h
	h.right = k
	c.left,c.right = f,g
	f.left = i
	g.right = j

	return a
}

