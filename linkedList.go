package main

type Node struct {
	prev *Node
	next *Node
	key interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Display() {
    list := l.head
    for list != nil {
        fmt.Printf("%+v ->", list.key)
        list = list.next
    }
    fmt.Println()
}

func main(){
	link := List{5, 9, 10, 11}
	link.Display()
}