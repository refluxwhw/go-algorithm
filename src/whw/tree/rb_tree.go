package tree

type Color int

const (
	Red   Color = 0
	Black Color = 1
)

type Node struct {
	D int
	R *Node
	L *Node
	C Color
}

type RBTree struct {
	root Node
}

func (*RBTree)Insert(d int) {

}
