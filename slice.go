package main

import "fmt"
import "slices"

func main(){
//	var stringArr [] string // this is a slice 
	var s [] string // this is an array

	fmt.Println("uninit:", s, s==nil, len(s))
	
	var r = make([]string, 3);
	fmt.Println("emp", r, len(s), "cap:",cap(r))

	r[0] ="s"
	r[1] ="p"
    r[2] ="K"
    r = append(r,"see if this will work or not")	
	fmt.Println(r[0])
	fmt.Println(r)
	// here length is zero but capacity is 3
    fmt.Println(len(r))

	// copying a slice in go

    cop := make([]string, len(r))	
	copy(cop,r)
	fmt.Println(cop)
	// now lets do some actual slicing.

	slice1 := r[:3]
	fmt.Println("slice1", slice1)

	t1 := []string {"g", "h"}
	t2 := []string {"g", "h"}
    
	fmt.Println(slices.Equal(t1,t2))

    // two dimensional slices 

	t3 := make([][]int, 3,3)

    fmt.Println(t3)

}
