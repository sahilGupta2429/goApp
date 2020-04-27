package main

import "fmt"

/**
Creaing pointers
derefrencing pointers
new function
working with nil
types of internal pointers
*/
func pointerStudy() {
	v1 := 42          //integer a is declared
	var p1 *int = &v1 // this astirix before type is declaring pointer to data of that type
	var p2 int = v1   // this astirix is derefrencing, drill throught the runtime and find the value at the address

	fmt.Println(v1, p1, p2, *p1, &v1, &p2) //&p2 and &v1(p1) are different

	//pointer arithmatic
	pAr := [3]int{1, 2, 3}
	p4 := &pAr[0]
	p5 := &pAr[1]
	//p6 := &pAr[1] - 4 //this operation for pointer arithmatic is not allowed in go
	var p6 *[3]int = &pAr                            //address of array [1,2,3] = &[1,2,3]
	fmt.Println(pAr, p4, *p4, p5, *p5, p6, *p6, &p6) //&p6 is address of address of pAr

	//Structure and pointers
	var stP *myPointerStruct = &myPointerStruct{foo: 42}
	fmt.Println(stP, *stP)

	//new function
	//with new function you can not initialize the value within pointer object
	var stP1 *myPointerStruct
	fmt.Println(stP1) //poiner object without initializing : it gives special value <nil>
	stP1 = new(myPointerStruct)
	//to initialize the value , we derefrence the pointer object and assign value to foo
	(*stP1).foo = 14               // * has lower precedence than . ,thus () are necessary
	fmt.Println(stP1, (*stP1).foo) //it gives address to struct = &{0}
	stP1.foo = 21                  // go actually interprets this exactly as (*stP1).foo ;
	fmt.Println(stP1.foo)          // this is just syntactical sugar

	// the data is copied in case of array
	apoi := [3]int{1, 2, 3}
	bpoi := apoi
	fmt.Println(apoi, bpoi)
	apoi[1] = 21
	fmt.Println(apoi, bpoi)

	//but in slice
	// slice apoi2 is copied in bpoi2 , but part of the data that is copied is pointer , not the
	// underlined data itself
	apoi2 := []int{1, 3, 3}
	bpoi2 := apoi2
	fmt.Println(apoi2, bpoi2)
	apoi2[1] = 43
	fmt.Println(apoi2, bpoi2)

	//another built in type that has this behaviour is map
	amap := map[string]string{"name": "ajay", "age": "64"}
	bmap := amap
	fmt.Println(amap, bmap)
	amap["name"] = "neelam"
	fmt.Println(amap, bmap)

}

type myPointerStruct struct {
	foo int
}