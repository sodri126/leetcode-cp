package main

import "fmt"

type LRUCache struct {
	capacity        int
	data            map[int]data // map[key]data
	orderingDataKey []int
}

type data struct {
	value        int
	indexDataKey int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:        capacity,
		data:            make(map[int]data),
		orderingDataKey: make([]int, 0),
	}
}

func (this *LRUCache) check(key int) (data, int) {
	d, ok := this.data[key]
	if !ok {
		return d, -1
	}

	return d, d.value
}

func (this *LRUCache) totalOrderingDataKey() int {
	return len(this.data)
}

func (this *LRUCache) addOrderingDataKey(key int) {
	this.orderingDataKey = append(this.orderingDataKey, key)
	if len(this.orderingDataKey) < this.capacity {
		return
	}

	this.deleteLowerOrderingKey()
}

func (this *LRUCache) deleteLowerOrderingKey() {
	this.orderingDataKey = this.orderingDataKey[1:]
}

func (this *LRUCache) deleteIdxOrderingKey(index int) {
	this.orderingDataKey = append(this.orderingDataKey[:index], this.orderingDataKey[index+1:]...)
}

func (this *LRUCache) upsertData(key int, d data) {
	this.data[key] = d
}

func (this *LRUCache) Get(key int) int {
	d, val := this.check(key)
	if val == -1 {
		return val
	}

	//lastRank := this.getLastRank()
	//this.deleteRank(d.lastRank)
	//this.rank[lastRank] = key
	//d.lastRank = lastRank
	//this.data[key] = d
	this.deleteIdxOrderingKey(d.indexDataKey)
	d.indexDataKey = this.totalOrderingDataKey()
	this.upsertData(key, d)
	this.addOrderingDataKey(key)

	return d.value
}

func (this *LRUCache) Put(key int, value int) {
	d, val := this.check(key)
	if val > -1 {
		this.deleteIdxOrderingKey(d.indexDataKey)
	}

	d.value = value
	d.indexDataKey = this.totalOrderingDataKey()
	this.upsertData(key, d)
	this.addOrderingDataKey(key)
	//var (
	//	lastRank                  = this.getLastRank()
	//	isAlreadyUpdatedLowerRank = false
	//)
	//
	//if this.isCapacityFull() {
	//	this.deleteData(this.lowerRank.key)
	//}
	//
	//d, val := this.check(key)
	//if val > -1 {
	//	this.updateLowerRank()
	//	isAlreadyUpdatedLowerRank = true
	//}
	//
	//d.value = value
	//d.lastRank = lastRank
	//this.data[key] = d
	//
	//l := lowerRank{
	//	key:      key,
	//	lastRank: lastRank,
	//}
	//
	//this.rank[lastRank] = key
	//
	//if this.isRankFull() && !isAlreadyUpdatedLowerRank {
	//	this.updateLowerRank()
	//}
	//
	//this.setLowerRank(l)
}

func (this *LRUCache) updateLowerRank() {
	//previousLowerRank := this.lowerRank
	//this.lowerRank.lastRank = previousLowerRank.lastRank + 1
	//this.lowerRank.key = this.rank[this.lowerRank.lastRank]
	//this.deleteRank(previousLowerRank.lastRank)
	//previousLowerRank := this.getLowerRank()
	//this.deleteLowerRankKey()
	//this.deleteRank(previousLowerRank.lastRank)
}

func (this *LRUCache) isCapacityFull() bool {
	return len(this.data) > this.capacity
}

//
//func (this *LRUCache) isRankFull() bool {
//	return len(this.rank) > this.capacity
//}
//
//func (this *LRUCache) getLastRank() int {
//	this.lastRank++
//	return this.lastRank
//}
//
//func (this *LRUCache) setLowerRank(l lowerRank) {
//	if l.lastRank <= this.lowerRank.lastRank {
//		this.lowerRank = l
//	}
//}
//
//func (this *LRUCache) deleteData(key int) {
//	delete(this.data, key)
//}
//
//func (this *LRUCache) deleteRank(lastRank int) {
//	delete(this.rank, lastRank)
//}
//
//func (this *LRUCache) addDataKeyRank(rank int) {
//	this.dataKeyRank = append(this.dataKeyRank, rank)
//	if len(this.dataKeyRank) < this.capacity {
//		return
//	}
//
//	this.deleteLowerRankKey()
//}
//
//func (this *LRUCache) deleteLowerRankKey() {
//	this.dataKeyRank = this.dataKeyRank[1:]
//}
//
//func (this *LRUCache) getLowerRank() lowerRank {
//	return lowerRank{
//		key:      this.rank[this.dataKeyRank[0]],
//		lastRank: this.dataKeyRank[0],
//	}
//}

func main() {
	//case1()
	//case2()
	commands := []string{
		"LRUCache", "put", "put", "put", "put", "put", "get", "put", "get", "get",
		"put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get",
		"put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put",
		"get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put",
		"put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put",
		"put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get",
		"put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put",
		"get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put",
		"get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get",
		"put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put",
		"put",
	}

	input := [][]int{
		{10}, {10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3},
		{5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9},
		{10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18},
		{13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21},
		{5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12},
		{4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9},
		{3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1},
		{6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28},
		{11, 26},
	}

	readInput(commands, input)
}

func case1() {
	lru := Constructor(2)
	lru.Put(1, 1)           // cache is {1=1}
	lru.Put(2, 2)           // cache is {1=1, 2=2}
	fmt.Println(lru.Get(1)) // return 1
	lru.Put(3, 3)           // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(lru.Get(2)) // returns -1 (not found)
	lru.Put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lru.Get(1)) // return -1 (not found)
	fmt.Println(lru.Get(3)) // return 3
	fmt.Println(lru.Get(4)) // return 4
}

func case2() {
	lru := Constructor(2)
	fmt.Print(lru.Get(2), " ") // get -1
	lru.Put(2, 6)
	fmt.Print(lru.Get(1), " ") // get -1
	lru.Put(1, 5)
	lru.Put(1, 2)
	fmt.Print(lru.Get(1), " ") // get 2
	fmt.Print(lru.Get(2), " ") // get 6
}

func readInput(commands []string, input [][]int) {
	var (
		lru LRUCache
		res = make(map[int]int)
		sc  []int
	)
	for index, command := range commands {
		if command == "LRUCache" {
			lru = Constructor(input[index][0])
		} else if command == "put" {
			lru.Put(input[index][0], input[index][1])
			fmt.Printf("Put -> Key: %d | Value: %d  | Index: %d | capacity: %d  | data: %+v | ordering keys: %v \n\n",
				input[index][0], input[index][1], index, len(lru.data), lru.data, lru.orderingDataKey)
		} else {
			val := lru.Get(input[index][0])
			fmt.Printf("Get -> Key: %d | Value: %d | Index: %d | capacity: %d  | data: %+v | ordering keys: %v\n\n",
				input[index][0], val, index, len(lru.data), lru.data, lru.orderingDataKey)
			res[index] = val
			sc = append(sc, val)
		}
	}

	fmt.Printf("Slice: %+v\n", sc)
	fmt.Printf("Map: %+v", res)
	return
}
