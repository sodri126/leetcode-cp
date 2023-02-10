package main

import (
	"fmt"
	"strings"
)

type LRUCache struct {
	capacity   int
	data       map[int]*Node // map[key]data
	linkedList *LinkedList
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

func (l *LinkedList) Insert(node *Node) {
	if l.head != nil {
		node.next = l.head
		l.head.prev = node
	}

	if l.tail == nil {
		l.tail = node
	}

	l.head = node
	l.length++
}

func (l *LinkedList) Delete(node *Node) {
	if node == nil {
		return
	}

	var (
		prevNode = node.prev
		nextNode = node.next
	)

	if prevNode != nil {
		prevNode.next = nextNode
	}

	if nextNode != nil {
		nextNode.prev = prevNode
	}

	if l.head == node {
		l.head = nextNode
	}

	if l.tail == node {
		l.tail = prevNode
	}

	node.next = nil
	node.prev = nil
	l.length--
}

func (l *LinkedList) Length() int {
	return l.length
}

func (l *LinkedList) GetHead() *Node {
	return l.head
}

func (l *LinkedList) GetTail() *Node {
	return l.tail
}

func (l *LinkedList) Print() string {
	buf := strings.Builder{}
	buf.WriteByte('[')
	for node := l.head; node != nil; node = node.next {
		buf.WriteString(fmt.Sprintf("%d ", node.key))
	}
	buf.WriteByte(']')
	return buf.String()
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:   capacity,
		data:       make(map[int]*Node),
		linkedList: &LinkedList{},
	}
}

func (this *LRUCache) check(key int) (*Node, int) {
	d, ok := this.data[key]
	if !ok {
		return d, -1
	}

	return d, d.value
}

func (this *LRUCache) checkFullCapacity() bool {
	return this.linkedList.Length() >= this.capacity
}

func (this *LRUCache) Get(key int) int {
	node, val := this.check(key)
	if val == -1 {
		return val
	}

	if this.linkedList.GetHead() != node {
		this.linkedList.Delete(node)
		this.linkedList.Insert(node)
	}

	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, val := this.check(key)
	if val > -1 {
		node.value = value
		this.linkedList.Delete(node)
		delete(this.data, key)
	} else {
		node = &Node{
			key:   key,
			value: value,
		}
	}

	if this.checkFullCapacity() {
		tail := this.linkedList.GetTail()
		if tail == nil {
			return
		}

		this.linkedList.Delete(tail)
		delete(this.data, tail.key)
	}

	this.data[key] = node
	this.linkedList.Insert(node)
}

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
			fmt.Printf("Put -> Key: %d | Value: %d  | Index: %d | capacity data: %d | capacity linkedlist: %d | data: %s | linkedlist: %s  \n\n",
				input[index][0], input[index][1], index, len(lru.data), lru.linkedList.Length(), printMap(lru.data), lru.linkedList.Print())
		} else {
			val := lru.Get(input[index][0])
			fmt.Printf("Get -> Key: %d | Value: %d | Index: %d | capacity data: %d | capacity linkedlist: %d | data: %s | linkedlist: %s  \n\n",
				input[index][0], val, index, len(lru.data), lru.linkedList.Length(), printMap(lru.data), lru.linkedList.Print())
			res[index] = val
			sc = append(sc, val)
		}
	}

	fmt.Printf("Slice: %+v\n", sc)
	fmt.Printf("Map: %+v", res)
	return
}

func printMap(data map[int]*Node) string {
	buf := strings.Builder{}
	buf.Write([]byte("["))
	for key, value := range data {
		val := "<nil>"

		if value != nil {
			val = fmt.Sprintf("%d", value.value)
		}

		format := fmt.Sprintf("Key: %d ~ Value: %s, ", key, val)
		buf.WriteString(format)
	}

	buf.Write([]byte("]"))
	return buf.String()
}
