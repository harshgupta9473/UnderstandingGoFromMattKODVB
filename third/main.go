package main

import (
	"fmt"
	"os"
)

// import "fmt"

// // Go initializes all variables to "zero" by default if you don't:
// //  All numerical types get 0 (float 0.0, complex 0i)
// // bool gets false
// // string gets ""
// // Everything ele get nil e.g: pointers , slices maps,channels,function,interfaces
// // for aggregate types, all members get their "zero" values

// var a int

// var (
// 	b = 2
// 	f = 2.01
// )

// func main() {

// 	// only inside function
// 	c := 2 // := short declaration operator
// 	fmt.Printf("a: %8T %[1]v\n",a)
// 	fmt.Printf("b: %8T %[1]v\n",b)
// 	fmt.Printf("f: %8T %[1]v\n",f)
// 	fmt.Printf("c: %8T %[1]v\n",c)

// 	a=int(f)
// 	fmt.Printf("a: %8T %[1]v\n",a)

// }

// // bool has two values false , true
// // these values are not convertible to/from integers!

// //pointers are physically addresses, logically opaque
// // a pointer may be nil or non nil
// // no pointer manipulation except through "package unsafe"

// // Only numbers, strings and booleans can be constants (immutable)
// // Constant can be a literal or a compile-time function of a constant

// const(
// 	x=1    //int
// 	y=2*1024 // 2048
// 	c=y<<3   // 16384

// 	g uint8 = 0x07      //7
// 	h uint8 = g & 0x03   //3

// 	s="a string"
// 	t=len(s)
// 	// u=s[2:]     //SYNTAX ERROR
// )

func main(){
	var sum  float64
	var n int

	for{
		var val float64
		_,err:=fmt.Fscanln(os.Stdin,&val)
		if err!=nil{
			break;
		}
		sum+=val
		n++
	}
	if n==0{
		fmt.Fprintln(os.Stderr,"no values")
		os.Exit(-1)
	}

	fmt.Println("The average is ",sum/float64(n))
}

//go run main.go < nums.txt  (=>prints the average by taking input from file)


//cat num.txt | go run main.go
//Error:'cat' is not recognized as an internal or external command,
//operable program or batch file.