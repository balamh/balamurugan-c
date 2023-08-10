package main
import(
	"fmt"
)
var(
	c1[] int
	sum int
	add int
)
func valid(val1 int){
	value:=val1
	for i:=0;value>0;i++{
		k:=value%10
		c1=append(c1,k)
		value/=10
	}
	if len(c1)<9 || len(c1)>16{
		fmt.Println(len(c1))
		fmt.Println("Invalid Card")
	}else{
		find1()
	}
}
func find1(){
	for i:=1;i<len(c1);i+=2{
		c1[i]=c1[i]*2
		if c1[i]>9{
			c1[i]=c1[i]-9
		}
	}
	for i:=0;i<len(c1);i++{
		
		add+=c1[i]
	}
	if add%10==0{
		fmt.Println("This is a Valid Card")
	}else{
		fmt.Println("Invalid Card,else")
	}
}
func main(){
	var cardno int
	fmt.Println("Enter Your Card Number:")
	fmt.Scan(&cardno)
	valid(cardno)
}