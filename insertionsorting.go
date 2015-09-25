package main
import "fmt"

func insert(a[] int){
	for i:=1;i<len(a) ;i++  {
		temp := a[i]
		j:=i-1
		for (j>=0 && a[j]> temp)  {
			a[j+1] = a[j]
			j=j-1

		}
		a[j+1]=temp
	}
	fmt.Println(" the sorted values is ",a)
}

func main(){
 list := []int{31, 41, 59, 26, 53, 58, 97, 93, 23, 84}
 fmt.Println(" the unsorted values in array are",list)
 insert(list)


}
