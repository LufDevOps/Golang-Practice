package main

import (
	"fmt"
	 "golang.org/x/exp/slices"
)


func main(){
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "30"
	s[1] = "10"
	s[2] = "2002"

	s = append(s, "18")
	s = append(s, "11","2003")

	fmt.Println("love:", s)

	l := s[3:6]
	fmt.Println("sl1:", l)

	l = s[:3]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3", l)

	t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

	t2 := []string{"g","h","i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
		
	}
	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}

