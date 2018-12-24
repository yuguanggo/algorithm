package binarysearch

type node struct {
	key   int
	value int
	left  *node
	right *node
}

type BTS struct {
	root  *node
	count int
}

func Init() *BTS {
	return &BTS{
		root:  nil,
		count: 0,
	}
}

func (bts *BTS) Size() int {
	return bts.count
}

func (bts *BTS) IsEmpty() bool {
	return bts.count == 0
}

// 向二分搜索树中插入一个新的(key, value)数据对
func (bts *BTS) Insert(key int, value int) {
	bts.root = insert(bts.root, key, value)
}

//在二分搜索树中搜索键key所对应的值。如果这个值不存在, 则返回null
func (bts *BTS) Search(key int,value int) int {
 return 0
}


// 向以node为根的二分搜索树中, 插入节点(key, value), 使用递归算法
// 返回插入新节点后的二分搜索树的根
func insert(root *node, key int, value int) *node {
	if root == nil {
		return &node{
			key:   key,
			value: value,
			left:  nil,
			right: nil,
		}
	}
	if root.key == key {
		root.value = value
	} else if root.key > key {
		root.left = insert(root.left, key, value)
	} else {
		root.right = insert(root.right, key, value)
	}
	return root
}
