package main
import(
	"fmt"
)
type node struct{
	value int
	next *node
}
type linkedlist struct{
	head *node
	tail *node
	length int
}
func (l *linkedlist) append(val int){
	newlist :=&node{val,nil}
	if l.length==0{
		l.head=newlist
		l.tail=newlist
		l.length++
	}else{
		l.tail.next=newlist
		l.tail=newlist
	}
}
func(l *linkedlist) print(){
	temp:=l.head
	if l.length==0{
		fmt.Println("No Element Present")
	}else{
		fmt.Println("Elements in the LinkedList")
		for temp!=nil{
			fmt.Println(temp.value)
			temp=temp.next
		}
		}	
}
func main(){
	list:=linkedlist{}
	list.append(1)
	list.append(2)
	list.print()
	}