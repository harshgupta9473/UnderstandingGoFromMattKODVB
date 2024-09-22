package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func main() {
// 	s := "elite"
// 	fmt.Printf("%8T %[1]v %d\n",s,len(s))
// 	fmt.Printf("%8T %[1]v\n",[]rune(s))
// 	fmt.Printf("%8T %[1]v %d\n",[]byte(s),len([]byte(s)))

// 	str:="hello, world"
// 	// x:=len(str)
// 	fmt.Println(strings.Contains(str,"w"))
// 	fmt.Println(strings.HasPrefix(str,"h"))
// 	fmt.Println(strings.ToUpper(str))
// 	fmt.Println(str)

// }

func main(){
	if len(os.Args)<3{
		fmt.Fprintln(os.Stderr,"not enough  args")
		os.Exit(-1)
	}

	old,nw:=os.Args[1],os.Args[2]
	scan:=bufio.NewScanner(os.Stdin)

	for scan.Scan(){
		s:=strings.Split(scan.Text(),old)
		t:=strings.Join(s,nw)
		fmt.Println(t)
	}
}

//go run main.go harsh satyam < test.txt