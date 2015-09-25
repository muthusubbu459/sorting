ackage main

import "fmt"

func sorting(a[20]int, n int){
	unsorted:=true
	for unsorted{
		unsorted=false
		for i:=n-1;i>0 ;i--{
			if(a[i]<a[i-1]){
				a[i],a[i-1]= a[i-1],a[i]
				unsorted=true
			}
		}
	}
	fmt.Println("the sorted array")
	for t:=0;t<n;t++{
		fmt.Println(a[t])
	}

}

func main(){
	var a[20] int
	var n int
	fmt.Println("enter the number of values you want to sort")
	fmt.Scanf("%d\t",&n)
	//fmt.Println("The total numbers going to be sorted is",n)
	for i:=0;i<n; i++ {
		fmt.Println("enter the number ",i)
	    fmt.Scanf("%d\t",&a[i])
	}
	sorting(a,n)

}
