package main

import (
	"fmt"
	"agents/world"
	"agents/worker"
	"agents/state"
)

func main () {
/*
	wood_1 := world.CreateRessource(12.58,16.54,"Wood")
	coal_1 := world.CreateRessource(12.2,14.30,"Coal")
	// fmt.Printf("%#s\n",n)
	wood_1.Display()
	coal_1.Display()
	// m.Display()
	fmt.Println("=====")

	// for _, element := range world.AllRessources {
	// 	// index is the index where we are
	// 	// element is the element from someSlice for where we are
	// 	element.Display()
	// }
	fmt.Println(world.RessourcesCreated)

	fmt.Println(r1)
	fmt.Println(r1.GetCity().GetNode())
*/
	mV2 := world.CreateMapV2(50,50,4,20)
	
	fmt.Printf("%v\n",world.AllCities)
	mV2.Display()
	t := worker.GenerateRuler(world.AllCities)
	fmt.Printf("%v\n",t)
	fmt.Printf("%v\n",t[0].GetCity())

	s := state.State{}
	fmt.Printf("%v\n",s)
}