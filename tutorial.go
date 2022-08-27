package main

import (
	"fmt"
	// "os"
)



func main(){
	// var x ="c";
	// fmt.Printf("%T\t%v\n",x,x);
	// fmt.Printf("%b\n%o\n%x\n",143,143,143);
	// fmt.Printf("%4.2f\n",231.456);
	// scanner:=bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// fmt.Printf("you typed %q",scanner.Text())
var x int16=6
var y int16=3
ans:=float32(x)/float32(y)
fmt.Printf("answer is %.2f\n",ans)
if(x==7){
	fmt.Println("same as 7")
}else if(x==6){
	fmt.Println("same as 6")
}else{
	fmt.Println("not same!!!")
}
}
