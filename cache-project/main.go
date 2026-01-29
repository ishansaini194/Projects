package main

import "fmt"

const Size = 5

type Node struct {
	val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(word string) {
	if node, found := c.Hash[word]; found {
		c.Remove(node)
		c.Add(node)
		return
	}

	node := &Node{val: word}
	c.Add(node)
	c.Hash[word] = node
}

func (c *Cache) Remove(node *Node) *Node {
	// never remove dummy nodes
	if node == c.Queue.Head || node == c.Queue.Tail {
		return nil
	}

	left := node.Left
	right := node.Right

	left.Right = right
	right.Left = left

	node.Left = nil
	node.Right = nil

	c.Queue.Length--
	delete(c.Hash, node.val)

	return node
}

func (c *Cache) Add(node *Node) {
	// insert right after head (MRU)
	first := c.Queue.Head.Right

	c.Queue.Head.Right = node
	node.Left = c.Queue.Head
	node.Right = first
	first.Left = node

	c.Queue.Length++
	c.Hash[node.val] = node

	if c.Queue.Length > Size {
		// remove LRU (node before tail)
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [ ", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.val)
		if i < q.Length-1 {
			fmt.Print(", ")
		}
		node = node.Right
	}
	fmt.Println(" ]")
}

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()

	for _, word := range []string{"parrot", "avacado", "dragonfruit", "tree", "potato", "tomato", "tree", "dog"} {
		cache.Check(word)
		cache.Display()
	}
}
