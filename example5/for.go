package main

import "fmt"

func main() {
	/* i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := 0; i < 3; i++ {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	} */
	for {
		fmt.Println("with continue")
		break
	}

	 for n := 0; n < 6; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	for {
		fmt.Println("print inside if")
		break
	}

	 for n := 0; n < 6; n++ {
		if n%2 == 0 {
			fmt.Println(n)
			// continue
		}
		
	}

	for {
		fmt.Println("without continue")
		break
	}

	for i := 0; i < 6; i++{
		if i%2 == 0 {
			// continue
		}
		fmt.Println(i)
	}
}