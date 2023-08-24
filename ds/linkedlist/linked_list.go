package linkedlist

// Node represents a node in a doubly linked list.
type Node struct {
	Val  any
	Next *Node
	Prev *Node
}

// LinkedList represents a doubly linked list.
type LinkedList struct {
	Head *Node
	Tail *Node
	len  int
}

// New returns an initialized list.
func New() *LinkedList {
	return &LinkedList{}
}

// setFirstNode sets the first node of the list
func (l *LinkedList) setFirstNode(node *Node) {
	l.Head = node
	l.Tail = node
}

// NewNode returns a new node with given value
func NewNode(val any) *Node {
	return &Node{Val: val}
}

// Len returns the number of elements of list l.
func (l *LinkedList) Len() int {
	return l.len
}

// PushFront inserts a new node with given val at the front of the list
func (l *LinkedList) PushFront(val any) {
	node := NewNode(val)
	if l.len == 0 {
		l.setFirstNode(node)
	} else {
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	}
	l.len++
}

// PushBack inserts a new node with given val at the back of the list
func (l *LinkedList) PushBack(val any) {
	node := NewNode(val)
	if l.len == 0 {
		l.setFirstNode(node)
	} else {
		l.Tail.Next = node
		node.Prev = l.Tail
		l.Tail = node
	}
	l.len++
}

// MoveToFront moves node to the front of the list
func (l *LinkedList) MoveToFront(node *Node) {
	if l.len == 0 {
		return
	} else {
		if node == l.Head {
			return
		} else if node == l.Tail {
			node.Next = l.Head
			l.Head.Prev = node
			node.Prev.Next = nil
			l.Tail = node.Prev
			node.Prev = nil
			l.Head = node
		} else {
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
			node.Next = l.Head
			l.Head.Prev = node
			l.Head = node
			node.Prev = nil
		}
	}
}

// MoveToBack moves node to the back of the list
func (l *LinkedList) MoveToBack(node *Node) {
	if l.len == 0 {
		return
	} else {
		if node == l.Tail {
			return
		} else if node == l.Head {
			node.Prev = l.Tail
			l.Tail.Next = node
			node.Next.Prev = nil
			l.Head = node.Next
			node.Next = nil
			l.Tail = node
		} else {
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
			l.Tail.Next = node
			node.Prev = l.Tail
			l.Tail = node
			node.Next = nil
		}
	}
}

// Remove removes node from list l.
func (l *LinkedList) Remove(node *Node) {
	if l.len == 0 {
		return
	} else {
		node.Val = nil
		if l.len == 1 {
			l.Head = nil
			l.Tail = nil
		} else {
			if l.Tail == node {
				l.Tail = node.Prev
				l.Tail.Next = nil
			} else if l.Head == node {
				l.Head = node.Next
				l.Head.Prev = nil
			} else {
				node.Prev.Next = node.Next
				node.Next.Prev = node.Prev
			}
		}
		l.len--
	}
}

// InsertAfter inserts a new node with given val after the given node
func (l *LinkedList) InsertAfter(after *Node, val any) {
	node := NewNode(val)
	if l.len == 0 {
		l.setFirstNode(node)
	} else {
		if after == l.Tail {
			after.Next = node
			node.Prev = after
			l.Tail = node
		} else {
			node.Prev = after
			node.Next = after.Next
			after.Next = node
			node.Next.Prev = node
		}
	}
	l.len++
}

// InsertBefore inserts a new element with val node immediately before given node.
func (l *LinkedList) InsertBefore(before *Node, val any) {
	node := NewNode(val)
	if l.len == 0 {
		l.setFirstNode(node)
	} else {
		if before == l.Head {
			before.Prev = node
			node.Next = before
			l.Head = node
		} else {
			node.Next = before
			node.Prev = before.Prev
			node.Prev.Next = node
			before.Prev = node
		}
	}
	l.len++
}

// Clear removes all nodes from the list.
func (l *LinkedList) Clear() {
	l.Head = nil
	l.Tail = nil
	l.len = 0
}
