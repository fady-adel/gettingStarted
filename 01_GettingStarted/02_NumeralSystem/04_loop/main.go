package main

import "fmt"

func main()  {
	for i := 100; i < 100000; i++ {
		fmt.Printf("%d\t%b\t%#x\n", i, i, i)
	}
}
