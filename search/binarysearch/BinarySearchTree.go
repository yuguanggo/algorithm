package binarysearch

type node struct {
	key int
	value int
	left *node
	right *node
}

type BTS struct {
	root node
	count int
}

func Init() *BTS   {
	return &BTS{
		root:nil,
		count:0,
	}
}


