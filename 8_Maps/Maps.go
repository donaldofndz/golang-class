package main

import "fmt"

func main() {
	fmt.Println("This is the maps explanation")
	//mapNil()
	//creatingAMap()
	//validatingExistingOnMap()
	//lengthAndDelete()
	iterateAMap()

}

func mapNil() {
	var m map[string]int

	fmt.Println(m)
	fmt.Println("m == nil", m == nil)
}

func creatingAMap() {
	mapAge := make(map[string]int)
	mapAge["donaldo"] = 25
	mapAge["ariana"] = 19

	anotherMap := map[int]string{
		1: "Cruz azul",
		2: "Atlas",
		3: "America",
	}

	fmt.Println(mapAge)
	fmt.Printf("The age of Donaldo is: %v ", mapAge["donaldo"])
	fmt.Println(anotherMap)
}

func validatingExistingOnMap() {
	mapAnimals := map[string]int{
		"perro": 1,
		"gato":  2,
		"rana":  3,
	}

	fmt.Println("If the element doesnt exist : ", mapAnimals["cuyo"])

	value, ok := mapAnimals["cuyo"]

	fmt.Printf("Value: %v , ok: %v\n", value, ok)

	value, ok = mapAnimals["rana"]

	fmt.Printf("Value: %v , ok: %v\n", value, ok)
}

func lengthAndDelete() {
	mapAnimals := map[string]int{
		"perro": 1,
		"gato":  2,
		"rana":  3,
	}
	fmt.Println(len(mapAnimals))
	delete(mapAnimals, "perro")
	fmt.Println(len(mapAnimals))
	fmt.Println(mapAnimals)
}

func iterateAMap() {

	mapaDeAnimales := make(map[string]int)

	mapAnimals := map[string]int{
		"perro": 1,
		"gato":  2,
		"rana":  3,
	}

	for key, value := range mapAnimals {
		mapaDeAnimales[key] = value
	}

	fmt.Println(mapaDeAnimales)
}
