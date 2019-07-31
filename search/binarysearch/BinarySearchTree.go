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
func (bts *BTS) GetRoot() *node {
	return bts.root
}

// 向二分搜索树中插入一个新的(key, value)数据对
func (bts *BTS) Insert(key int, value int) {
	bts.root = insert(bts.root, key, value)
	bts.count += 1
}

//在二分搜索树中搜索键key所对应的值。如果这个值不存在, 则返回null
func (bts *BTS) Search(key int) interface{} {
 return search(bts.root,key)
}

//将对应的key删除，并返回根节点
func (bts *BTS) Remove(key int) *node  {
	bts.root=remove(bts.root,key)
	return bts.root
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func removeMin(node *node) *node  {
	if node.left==nil{
		right:=node.right
		node=nil
		return right
	}
	node.left=removeMin(node.left)
	return node
}

// 返回以node为根的二分搜索树的最小键值所在的节点
func minimum(node *node) *node  {
	if node.left==nil{
		return node
	}
	return minimum(node.left)
}
func remove(node *node,key int) *node  {
	if node==nil{
		return nil
	}
	if node.key>key{
		node.right=remove(node.right,key)
		return node
	}else if node.key<key{
		node.left=remove(node.left,key)
		return node
	}else{
		if node.left==nil{
			right:=node.right
			node=nil
			return right
		}
		if node.right==nil{
			left:=node.left
			node=nil
			return left
		}

		// 待删除节点左右子树均不为空的情况

		// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor:=minimum(node.right)
		successor.right=removeMin(node.right)
		successor.left=node.left
		node=nil
		return successor
	}
}

func search(node *node,key int) interface{}{
	if node.key==key{
		return node.value
	}else if node.key>key{
		return search(node.left,key)
	}else{
		return search(node.right,key)
	}

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
