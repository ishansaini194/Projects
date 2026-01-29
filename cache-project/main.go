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
	Hash  map[string]*Node
}

func NewCache() Cache {
	return Cache{
		Queue: NewQueue(),
		Hash:  make(map[string]*Node),
	}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

/* ---------------- Core Logic ---------------- */

func (c *Cache) Check(word string) {
	if node, found := c.Hash[word]; found {
		c.moveToFront(node)
	} else {
		node := &Node{val: word}
		c.addToFront(node)
		c.Hash[word] = node

		if c.Queue.Length > Size {
			lru := c.Queue.Tail.Left
			c.removeNode(lru)
			delete(c.Hash, lru.val)
		}
	}
}

func (c *Cache) moveToFront(node *Node) {
	c.removeNode(node)
	c.addToFront(node)
}

func (c *Cache) removeNode(node *Node) {
	left := node.Left
	right := node.Right

	left.Right = right
	right.Left = left

	c.Queue.Length--
}

func (c *Cache) addToFront(node *Node) {
	first := c.Queue.Head.Right

	c.Queue.Head.Right = node
	node.Left = c.Queue.Head
	node.Right = first
	first.Left = node

	c.Queue.Length++
}

/* ---------------- Display ---------------- */

func (c *Cache) Display() {
	node := c.Queue.Head.Right
	fmt.Printf("%d - [ ", c.Queue.Length)
	for i := 0; i < c.Queue.Length; i++ {
		fmt.Printf("{%s}", node.val)
		if i < c.Queue.Length-1 {
			fmt.Print(", ")
		}
		node = node.Right
	}
	fmt.Println(" ]")
}

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()
	for _, word := range []string{
		"parrot", "avacado", "dragonfruit", "tree",
		"potato", "tomato", "tree", "dog",
	} {
		cache.Check(word)
		cache.Display()
	}
}
