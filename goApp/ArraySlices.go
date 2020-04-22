package main

import "fmt"

func arrSl() {

	// giving size of array
	grade2 := [3]int{96, 94, 95}
	fmt.Printf("%v\n", grade2)

	// [...]  it says get the size exactly as required
	grade3 := [...]int{96, 94, 95}
	fmt.Printf("%v\n", grade3)

	var stud [3]string
	stud[0] = "Sahil"
	stud[1] = "Ravish"
	stud[2] = "Akshay"
	fmt.Printf("%v\n", stud)
	fmt.Printf("%v\n", len(stud))

	//*V IMp : Here in GO array works different when we copy an array then in here refrence
	// is not copied to that variable instead another copy of whole value is created in the memory
	ax := [...]int{1, 2, 3}
	bx := ax
	bx[1] = 5
	fmt.Println(ax)
	fmt.Println(bx)

	// TO only refer the pointer of an array to another variable we use
	cx := &ax // so now same array is present in memory but pointer is copied refrencing
	// the same array to cx
	cx[1] = 8 //this will change value in both cx and ax
	fmt.Println(ax)
	fmt.Println(bx)
	fmt.Println(cx)

	// Here no size is given , this is SLICE . Array without any size is slice
	grade := []int{96, 94, 95}
	fmt.Printf("%v\n", grade)
	fmt.Printf("Length %v\n", len(grade))
	fmt.Printf(" Capacity %v\n", cap(grade))

	//But in cases of slice values are not copied the refrence is copied seperately
	gradeTest := grade
	grade[1] = 89
	fmt.Println(grade)
	fmt.Println(gradeTest)

	asl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //Slice
	bsl := asl[:]                               //slice of all elements
	csl := asl[3:]                              //slice from 4th to last element
	dsl := asl[:6]                              //slice of first 6 elements
	esl := asl[3:6]                             //slice of 4,5,6th element
	asl[5] = 87
	fmt.Println(asl)
	fmt.Println(bsl)
	fmt.Println(csl)
	fmt.Println(dsl)
	fmt.Println(esl)

	//Now this slicing operations work with arrays as well so,
	asl2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //Array
	bsl2 := asl2[:]                                 //slice of all elements
	csl2 := asl2[3:]                                //slice from 4th to last element
	dsl2 := asl2[:6]                                //slice of first 6 elements
	esl2 := asl2[3:6]                               //slice of 4,5,6th element
	asl2[5] = 87                                    // this value is changed and then slicing function is performed at runtime
	fmt.Println(asl2)
	fmt.Println(bsl2)
	fmt.Println(csl2)
	fmt.Println(dsl2)
	fmt.Println(esl2)
	//The result remains the same.  Slicing operation interprets arrays and slices both

	//Make function to create a slice
	//Unlink arrays slices don't have to have a fixed size their entire life
	makeSlice := make([]int, 3, 100) //(datatype, size, capacity) of the slice
	fmt.Println(makeSlice)
	fmt.Println(len(makeSlice))
	fmt.Println(cap(makeSlice))
	//We can actually add and remove elements from slice
	makeSlice2 := []int{}
	printSlice(makeSlice2)
	//makeSlice2 = append(makeSlice2, 1, 2, 3, 4, 5)
	// printSlice(makeSlice2)
	//Spread operator to add the slices together
	//Concatination of slices
	makeSlice2 = append(makeSlice2, []int{1, 2, 3, 4, 5}...)
	printSlice(makeSlice2)

	//POP elements from slice
	makeSlice3 := []int{11, 12, 13, 14, 15}
	printSlice(makeSlice3)
	makeSlice32 := append(makeSlice3[:2], makeSlice3[4:]...) //removing middle element
	printSlice(makeSlice32)
	printSlice(makeSlice3) // very wierd behaviour
	// this is because we are playing with refrences here and changing makeSlice3 to get
	// makeSlice32 changes makeSlice3 itself

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
