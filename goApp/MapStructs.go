package main

import (
	"fmt"
	"reflect"
)

func testMapStruct() {

	//Map Key value pair key:string , value:int
	//main condition for key is , it should be tested for equality
	// valid key types are strings,boolean,all numeric types,arrays,structures,pointers,channels
	// invalid key types slices,maps and other functions
	statePopulation := map[string]int{
		"California": 35476394,
		"Texas":      125457566,
		"Florida":    515435455,
	}
	fmt.Println(statePopulation)

	//Using make function for map
	stPopulation := make(map[string]int)
	stPopulation = map[string]int{
		"California": 35476394,
		"Texas":      125457566,
		"Florida":    515435455,
	}
	fmt.Println(stPopulation)
	fmt.Println(stPopulation["California"])

	//Adding value in map
	//Order of map is not fixed
	stPopulation["Georgia"] = 102554684
	fmt.Println(stPopulation)

	//Deleting from map
	delete(stPopulation, "Georgia")
	fmt.Println(stPopulation)

	// check if key is present
	// Ok tells if key is present and value 0 can be if actually 0 is present in map also
	populationCheck, ok := stPopulation["Texa"]
	fmt.Println(populationCheck, ok)

	populationCheck2, ok2 := stPopulation["Texas"]
	fmt.Println(populationCheck2, ok2)

	//If we dont care about population value
	_, ok3 := stPopulation["Texas"]
	fmt.Println(ok3)

	//length of map
	fmt.Println(len(stPopulation))

	//copying map to another map is done by reference
	// so manupulation of one map will have effect on underlaying map
	fmt.Println(statePopulation)
	childMap := statePopulation
	delete(childMap, "Texas")
	fmt.Println(childMap)
	fmt.Println(statePopulation)
	strucFunction()
}

type Doctor struct {
	number     int
	actor      string
	companions []string
}

func strucFunction() {
	aDoc := Doctor{
		number: 3,
		actor:  "Sahil",
		companions: []string{
			"Akshay",
			"Ravish",
		},
	}

	fmt.Println(aDoc)
	fmt.Println(aDoc.actor)
	fmt.Println(aDoc.companions)
	fmt.Println(aDoc.companions[1])

	// anonymous struct
	aDoctor := struct{ name string }{name: "Sahil"}
	fmt.Println(aDoctor)

	//passing a struct to another
	//Here the  structure is passed by value thus the value is not affecting underlaying structure
	// like in slice or map
	fmt.Println(aDoctor)
	anotherDoc := aDoctor
	anotherDoc.name = "Ravish"
	fmt.Println(anotherDoc)
	fmt.Println(aDoctor)

	// If you want to refer to same structure it can be done by address operator &
	fmt.Println(aDoctor)
	anDoc := &aDoctor
	anDoc.name = "Akshay"
	fmt.Println(anDoc)
	fmt.Println(aDoctor)

	//
	compReln()
}

// tag
type Animal struct {
	Name   string `max:"100"`
	Origin string
}

//Composition type relaitionship : this is called EMBEDDING
type Bird struct {
	Animal
	Speed  float32
	CanFly bool
}

func compReln() {
	//Here bird has a animal type charactorstics which
	// is similar to inheritance but not inheritance
	// This is called composition type relationship i.e EMBEDDING, which 'has a'
	// whereas inheritance is more like is a relationship

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.Speed = 48
	b.CanFly = false
	fmt.Println(b)

	////////////////////////////////////
	strucReflect := reflect.TypeOf(Animal{})
	field, _ := strucReflect.FieldByName("Name")
	fmt.Println(strucReflect)
	fmt.Println(field.Tag)
}
