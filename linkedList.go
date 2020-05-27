package main


type list struct {
	head *Node
	tail *Node
}

type Node struct {
	value int
	next *Node
}

func (l *list)first() *Node {
   return l.head
} 

func (l *list) Push(value int) {
	node := &Node{value: value}

	if(l.head == nil) {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

func (n *Node)Next() *Node {
	return n.next
}

func main(){
	l := &list{}
	l.Push(11)
	l.Push(32)
	l.Push(3)
		n := l.first()

	for {
		println(n.value)
		n = n.Next()
		if n == nil {
			break
		}
	}

}