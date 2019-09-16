package util

type BTree struct {
	Value int
	Right *BTree
	Left  *BTree
}

func NewBTreeFromArray(tree []int) *BTree {
	len := len(tree)
	if len <= 0 {
		return nil
	}
	var bTreeArr []*BTree
	for i := 0; i <= len-1; i++ {
		bTreeArr = append(bTreeArr, &BTree{Value: tree[i]})
	}
	for i := 0; i <= (len-1)/2; i++ {
		if 2*i+1 <= len-1 {
			bTreeArr[i].Left = bTreeArr[2*i+1]
		}
		if 2*i+2 <= len-1 {
			bTreeArr[i].Right = bTreeArr[2*i+2]
		}
	}
	return bTreeArr[0]
}

func (bTree *BTree) ToArray() []int {
	if bTree == nil {
		return nil
	}
	var bTreeArr []*BTree
	var result []int
	i := 0
	bTreeArr = append(bTreeArr, bTree)
	result = append(result, bTree.Value)
	for i < len(bTreeArr) {
		cur := bTreeArr[i]
		if cur.Left != nil {
			bTreeArr = append(bTreeArr, cur.Left)
			result = append(result, cur.Left.Value)
		}
		if cur.Right != nil {
			bTreeArr = append(bTreeArr, cur.Right)
			result = append(result, cur.Right.Value)
		}
		i++
	}
	return result
}
