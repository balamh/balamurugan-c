package main
import(
	"fmt"
)
type queue struct{
	value[] int
}
func(q *queue)enqueue(val int){
	q.value=append(q.value,val)
}
func (q *queue)dequeue(){
	if len(q.value)==0{
		fmt.Println("No ELement Present")
	}else{
		q.value=q.value[1:]
	}
}
func(q *queue)display(){
	if len(q.value)==0{
		fmt.Println("No Element Present")
	}else{
		fmt.Println("The Queue Elements are:")
	for i:=0;i<len(q.value);i++{
		fmt.Println(q.value[i])
	}
}
}
func main(){
	q:=queue{}
	q.enqueue(5)
	q.enqueue(7)
	q.enqueue(2)
	q.enqueue(4)
	q.dequeue()
	q.display()


}