package main

import (
	"fmt"

	"github.com/XuananLe/ConcurrentSafeDS/lists/arraylist"
)

func main() {
	x := arraylist.ArrayList[int]{};
	x.Initialize([]int{});
	x.Append([]int{1,2,3});
	fmt.Printf("x.String(): %v\n", x.String())
	x.Delete(3);
	fmt.Printf("x.String(): %v\n", x.String())
}