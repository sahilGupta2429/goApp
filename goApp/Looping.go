package main

import "fmt"

func loopPrac() {
	//For Loop => i is scoped to for loop
	for i := 0; i < 5; i++ {
		fmt.Println("for loop :", i)
	}

	//multiple initialize values
	//i,j = i++ ,j++ => This will not work, i++ is not an increment expression
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Printf("i : %v and j : %v \n", i, j)
	}

	//index is scoped to main function
	index := 0
	for ; index < 5; index++ {
		fmt.Println("index :", index)
	}

	// HOW go does while loop
	iWhile := 0
	for iWhile < 5 {
		fmt.Println("While vala for :", iWhile)
		iWhile++
	}

	//ANother syntax for the same
	iWhile2 := 0
	for iWhile2 < 5 {
		//No need for ;
		fmt.Println("While : ", iWhile2)
		iWhile2++
	}

	// Break keyword and for as a looping keyword
	iFor := 0
	for {
		fmt.Println("break keyword :", iFor)
		iFor++
		if iFor == 5 {
			break
		}
	}

	// Adding loop label , break Loop puts the control outside both for loops
Loop:
	// Label and Break in for loops
	for i1 := 1; i1 <= 3; i1++ {
		for i2 := 1; i2 <= 3; i2++ {
			fmt.Println(i1 * i2)
			if i1*i2 >= 3 {
				break Loop // this just breaks out of closest loop
			}
		}
	}

	// FOR - RANGE LOOP
	fRange := []int{1, 2, 3, 4, 5}
	for k, v := range fRange {
		fmt.Printf("key : %v , Val : %v \n", k, v)
	}

	// for k := range x   => when u only want keys
	// for _,v := range x => when you only want values

}
