package aoi

type DoubleNode struct {
	int X
	int Y
	int Radius

	PrevX *DoubleNode
	NextX *DoubleNode

	PrevY *DoubleNode
	NextY *DoubleNode

	// 可以新增一些身份标识符
	uint64 Id
}

type CrossLink struct {
	HeadX *DoubleNode
	HeadY *DoubleNode
}

func NewCrossLink() *CrossLink {

	crosslink := &CrossLink{
		HeadX: NewDoubleNode()
		HeadY: NewDoubleNode()
	}

	return crosslink
}

func (crosslink *CrossLink) Insert(node *DoubleNode) {

	if crosslink.HeadX == nil {
		crosslink.HeadX = node
	} else {
	
		cur := crosslink.HeadX.NextX

		for {
		
			if cur == nil {
				break
			}

			if node.X < cur.X {
				break	
			}
		}


	}
}

func (crosslink *CrossLink) Remove(node *DoubleNode) {

}

func (crosslink *CrossLink) Update(node *DoubleNode) {

}

func (crosslink *CrossLink) Move(node *DoubleNode, x int, y int) {

}
