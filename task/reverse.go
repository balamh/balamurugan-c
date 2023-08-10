package main
import(
	"fmt"
)
func reverse(a[] byte)string{
	for i:=2;i<len(a);i+=4{
		j:=i+1
		a[i],a[j]=a[j],a[i]
	}
	return string(a)
}
func main(){
	val:="abcdefghijkl"
	result:=reverse([]byte(val))
	fmt.Println(result)
}