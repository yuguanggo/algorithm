package main

/**
LFU缓存
设计并实现最不经常使用（LFU）缓存的数据结构。它应该支持以下操作：get 和 put。

get(key) - 如果键存在于缓存中，则获取键的值（总是正数），否则返回 -1。
put(key, value) - 如果键不存在，请设置或插入值。当缓存达到其容量时，它应该在插入新项目之前，
使最不经常使用的项目无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，最近最少使用的键将被去除。
 */

type lfuVc struct {
	value int
	count int
}

type LFUCache struct {
	len     int           //长度
	cap     int           //容量
	minFreq int           //最少使用的次数
	mapKvs  map[int]lfuVc //存储key value find o(1)
	mapFreq map[int][]int //存储有相同次数的key
	mapIter map[int]int   //存储key在mapFreq queue中的位置
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		len:     0,
		cap:     capacity,
		minFreq: 1,
		mapKvs:  make(map[int]lfuVc),
		mapFreq: make(map[int][]int),
		mapIter: make(map[int]int),
	}
}

func (this *LFUCache) Get(key int) int {
	if vc, ok := this.mapKvs[key]; ok {
		//更新key的使用次数
		index := this.mapIter[key]
		queue := this.mapFreq[vc.count]
		if len(queue) > 1 {
			queue = append(queue[:index], queue[index+1:]...) //删除key
			this.mapFreq[vc.count] = queue
			for k, v := range queue {
				this.mapIter[v] = k
			}
		} else {
			delete(this.mapFreq, vc.count)
			if vc.count==this.minFreq{
				this.minFreq++
			}
		}

		vc.count += 1
		this.mapKvs[key] = vc
		this.mapFreq[vc.count] = append(this.mapFreq[vc.count], key)
		this.mapIter[key] = len(this.mapFreq[vc.count]) - 1
		return vc.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap <= 0 {
		return
	}
	if this.Get(key) != -1 {
		vc := this.mapKvs[key]
		vc.value = value
		this.mapKvs[key] = vc
		return
	}
	if this.len >= this.cap {
		//删除最近不经常使用的值
		delkey := this.mapFreq[this.minFreq][0]
		if len(this.mapFreq[this.minFreq]) > 1 {
			this.mapFreq[this.minFreq] = this.mapFreq[this.minFreq][1:]
			for k, v := range this.mapFreq[this.minFreq] {
				this.mapIter[v] = k
			}
		} else {
			delete(this.mapFreq, this.minFreq)
		}

		delete(this.mapKvs, delkey)
		delete(this.mapIter, delkey)
		this.len--
	}
	vc := lfuVc{
		value: value,
		count: 1,
	}
	this.minFreq = 1
	this.mapKvs[key] = vc
	this.mapFreq[this.minFreq] = append(this.mapFreq[this.minFreq], key)
	this.mapIter[key] = len(this.mapFreq[vc.count]) - 1
	this.len++
}

func main() {
	cache := Constructor(10)
	s := [][]int{
		{10, 13},
		{3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26}}
	for _,v:=range s{
		if len(v)>1{
			cache.Put(v[0],v[1])
		}else{
			cache.Get(v[0])
		}
	}
}
