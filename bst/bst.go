package bst

type node struct {
	value     int
	leftNode  *node
	rightNode *node
}

func insertNode(val int, head *node) *node {
	tmpNode := &node{value: val, leftNode: nil, rightNode: nil}
	if head == nil {
		head = tmpNode
	}
	searchNode := head
	for {
		if val < searchNode.value {
			if searchNode.leftNode == nil {
				searchNode.leftNode = tmpNode
				return head
			}
			searchNode = searchNode.leftNode
		} else if val >= searchNode.value {
			if searchNode.rightNode == nil {
				searchNode.rightNode = tmpNode
				return head
			}
			searchNode = searchNode.rightNode
		}

	}

}
